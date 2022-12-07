use std::{cell::RefCell, collections::HashMap, rc::Rc};

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let root = parse_file_system(&input);

        println!("result {}", root.size_at_most(100000).to_string());
    }
    fn solve2(&self, input: String) {
        let root = parse_file_system(&input);

        let total = 70000000;
        let required = 30000000;
        let current = total - root.size();
        let to_free = required - current;

        println!(
            "result {}",
            root.free_up_space(to_free).unwrap().to_string()
        );
    }
}

#[derive(Debug, Clone)]
struct Folder<'a> {
    is_file: bool,
    size: Option<u64>,
    parent: Option<Rc<RefCell<Folder<'a>>>>,
    children: HashMap<&'a str, Rc<RefCell<Folder<'a>>>>,
}

impl<'a> Folder<'a> {
    fn size(&self) -> u64 {
        if self.is_file {
            return self.size.unwrap();
        }
        return self.children.values().map(|c| c.borrow().size()).sum();
    }

    fn size_at_most(&self, limit: u64) -> u64 {
        if self.is_file {
            return 0;
        }
        let size = self.size();
        // Get the size of the ancestors
        let ancestors = if size <= limit { size } else { 0 }
            + self
                .children
                .values()
                .map(|c| c.borrow().size_at_most(limit))
                .sum::<u64>();

        ancestors
    }

    fn free_up_space(&self, to_free: u64) -> Option<u64> {
        if self.is_file {
            return None;
        }

        // Try to free up space in the children
        let children_size = self
            .children
            .values()
            .filter_map(|c| c.borrow().free_up_space(to_free))
            .min();

        if children_size.is_some() {
            return children_size;
        }

        let size = self.size();
        if size >= to_free {
            return Some(size);
        }

        None
    }
}

fn parse_file_system(input: &str) -> Folder {
    let root = Folder {
        size: None,
        is_file: false,
        parent: None,
        children: Default::default(),
    };
    // Create the root node
    let root = Rc::new(RefCell::new(root));

    // The current directory is the root
    let mut current = Rc::clone(&root);
    for line in input.trim().lines() {
        let l = line.split_ascii_whitespace().collect::<Vec<&str>>();
        let l = l.as_slice();

        match l {
            // If the command is cd, we need to change the current directory
            ["$", "cd", dest] => match *dest {
                "/" => current = Rc::clone(&root), // We move to the root
                ".." => {
                    // We move to the parent
                    let parent = Rc::clone(current.borrow().parent.as_ref().unwrap());
                    current = parent;
                }

                dest => {
                    // We move to a sub directory
                    let sub_dir = Rc::clone(current.borrow().children.get(dest).unwrap());
                    current = sub_dir;
                }
            },
            // If the command is ls, Just skip it
            ["$", "ls"] => (),
            // If the command is a dir, we need to create a new directory
            ["dir", sub_dir] => {
                let node = Folder {
                    size: None,
                    is_file: false,
                    parent: Some(Rc::clone(&current)),
                    children: Default::default(),
                };

                current
                    .borrow_mut()
                    .children
                    .insert(sub_dir, Rc::new(RefCell::new(node)));
            }
            // If the line is something else, it's a file with the size
            [size, file] => {
                let file_size: u64 = size.parse().unwrap();

                // Register the file in the current directory
                let file_entry = Folder {
                    size: Some(file_size),
                    is_file: true,
                    parent: Some(Rc::clone(&current)),
                    children: Default::default(),
                };

                current
                    .borrow_mut()
                    .children
                    .insert(file, Rc::new(RefCell::new(file_entry)));
            }
            _ => unreachable!(),
        }
    }

    return root.borrow().clone();
}
