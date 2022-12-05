use std::collections::HashMap;

use crate::resolver::challenge::ChallengeResolver;
use itertools::Itertools;
use regex::Regex;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let (crates, operations) = split_input(input);
        let mut stacks = parse_stacks(crates);
        let moves = parse_moves(operations);

        // Execute moves
        stacks = execute_move(stacks, moves, false);

        print_last_of_stack(stacks);
    }

    fn solve2(&self, input: String) {
        let (crates, operations) = split_input(input);
        let mut stacks = parse_stacks(crates);
        let moves = parse_moves(operations);

        // Execute moves
        stacks = execute_move(stacks, moves, true);

        print_last_of_stack(stacks);
    }
}

#[derive(Debug)]
struct Move {
    quantity: u32,
    from: u32,
    to: u32,
}

fn split_input(input: String) -> (String, String) {
    let mut splitted = input.split("\n\n");
    return (
        splitted.next().unwrap().to_string(),
        splitted.next().unwrap().to_string(),
    );
}

fn parse_moves(input: String) -> Vec<Move> {
    let mut moves = Vec::new();
    let re = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
    for line in input.lines() {
        let caps = re.captures(line).unwrap();

        moves.push(Move {
            quantity: caps.get(1).unwrap().as_str().parse::<u32>().unwrap(),
            from: caps.get(2).unwrap().as_str().parse::<u32>().unwrap(),
            to: caps.get(3).unwrap().as_str().parse::<u32>().unwrap(),
        });
    }
    return moves;
}

fn parse_stacks(input: String) -> HashMap<u32, Vec<char>> {
    let mut stacks = HashMap::new();
    // Stacks looks like this
    //    [D]
    //[N] [C]
    //[Z] [M] [P]
    // 1   2   3

    // We need to parse the input and create a map of stacks
    // 1 => [Z, N]
    // 2 => [M, C, D]
    // 3 => [P]

    // We are starting from the bottom of the stack
    // and going up
    let lines = input.lines();
    // The last line will help us to get the key for the hashmap
    let last_line = lines.last().unwrap();
    let stack_keys = last_line
        .trim()
        .split_ascii_whitespace()
        .collect::<Vec<&str>>()
        .into_iter()
        .map(|x| x.to_string().parse::<u32>().unwrap())
        .collect::<Vec<u32>>();
    // Populate the hashmap
    for (_, stack_key) in stack_keys.iter().enumerate() {
        stacks.insert(*stack_key, Vec::new());
    }

    // Start from the bottom, skip the last line
    for line in input.lines().rev().skip(1) {
        let mut stack_index = 0;
        line.chars().skip(1).step_by(4).for_each(|c| {
            if c != ' ' {
                stacks.get_mut(&stack_keys[stack_index]).unwrap().push(c);
            }
            stack_index += 1;
        });
    }

    return stacks;
}

fn execute_move(
    stack: HashMap<u32, Vec<char>>,
    moves: Vec<Move>,
    at_once: bool,
) -> HashMap<u32, Vec<char>> {
    // Copy the stack
    let mut stack = stack;
    for m in moves {
        // get the quantity of crates from the from stack
        let from_stack = stack.get_mut(&m.from).unwrap();
        let crates_to_move = from_stack[from_stack.len() - m.quantity as usize..].to_vec();
        // remove the crates from the from stack
        from_stack.truncate(from_stack.len() - m.quantity as usize);
        // add the crates to the to stack
        let to_stack = stack.get_mut(&m.to).unwrap();
        if at_once {
            to_stack.append(&mut crates_to_move.clone());
        } else {
            to_stack.append(&mut crates_to_move.clone().into_iter().rev().collect());
        }
    }
    return stack;
}

fn print_last_of_stack(stack: HashMap<u32, Vec<char>>) {
    let mut last_of_stack: Vec<char> = Vec::new();
    for i in stack.keys().sorted() {
        let last = stack.get(i).unwrap().last().unwrap();
        last_of_stack.push(last.to_owned());
    }
    println!(
        "Last of stack: {:?}",
        last_of_stack.iter().collect::<String>()
    );
}
