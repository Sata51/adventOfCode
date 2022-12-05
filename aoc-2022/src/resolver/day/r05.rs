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
        stacks = execute_move_by_one(stacks, moves);

        print_last_of_stack(stacks);
    }

    fn solve2(&self, input: String) {
        let (crates, operations) = split_input(input);
        let mut stacks = parse_stacks(crates);
        let moves = parse_moves(operations);

        // Execute moves
        stacks = execute_all_move_at_once(stacks, moves);

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
    // move instruct have form
    // move 1 from 2 to 1
    let re = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
    for line in input.lines() {
        let caps = re.captures(line).unwrap();
        let quantity = caps.get(1).unwrap().as_str().parse::<u32>().unwrap();
        let from = caps.get(2).unwrap().as_str().parse::<u32>().unwrap();
        let to = caps.get(3).unwrap().as_str().parse::<u32>().unwrap();

        moves.push(Move { quantity, from, to });
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
        // Remove every %4 char
        let mut new_char_line: Vec<char> = Vec::new();
        for (i, c) in line.chars().enumerate() {
            if (i + 1) % 4 == 0 {
                continue;
            }
            new_char_line.push(c);
        }

        for w in new_char_line.chunks(3) {
            // If the 3 chars whitespace, we are on an empty stack, skip it
            if w[0] == ' ' && w[1] == ' ' && w[2] == ' ' {
                stack_index += 1;
                continue;
            }
            // If the first char is a [ we are on a stack
            if w[0] == '[' {
                stacks.get_mut(&stack_keys[stack_index]).unwrap().push(w[1]);
                // Increment the stack index
                stack_index += 1;
            }
        }
    }

    return stacks;
}

fn execute_move_by_one(
    stack: HashMap<u32, Vec<char>>,
    moves: Vec<Move>,
) -> HashMap<u32, Vec<char>> {
    // Copy the stack
    let mut stack = stack;
    for m in moves {
        // get the quantity of crates from the from stack
        let crates = stack.get(&m.from).unwrap().clone();
        let crates_to_move = crates[crates.len() - m.quantity as usize..].to_vec();
        // Reverse the crates to move
        let crates_to_move = crates_to_move.into_iter().rev().collect::<Vec<char>>();
        // remove the crates from the from stack
        let from_stack = stack.get_mut(&m.from).unwrap();
        from_stack.truncate(from_stack.len() - m.quantity as usize);
        // add the crates to the to stack
        let to_stack = stack.get_mut(&m.to).unwrap();
        to_stack.append(&mut crates_to_move.clone());
    }
    return stack;
}

fn execute_all_move_at_once(
    stack: HashMap<u32, Vec<char>>,
    moves: Vec<Move>,
) -> HashMap<u32, Vec<char>> {
    // Copy the stack
    let mut stack = stack;
    for m in moves {
        // get the quantity of crates from the from stack
        let crates = stack.get(&m.from).unwrap().clone();
        let crates_to_move = crates[crates.len() - m.quantity as usize..].to_vec();
        // remove the crates from the from stack
        let from_stack = stack.get_mut(&m.from).unwrap();
        from_stack.truncate(from_stack.len() - m.quantity as usize);
        // add the crates to the to stack
        let to_stack = stack.get_mut(&m.to).unwrap();
        to_stack.append(&mut crates_to_move.clone());
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
