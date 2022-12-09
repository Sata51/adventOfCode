use std::collections::HashSet;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

#[derive(Debug)]
enum Direction {
    Right,
    Left,
    Up,
    Down,
}

#[derive(Debug)]
struct Move {
    direction: Direction,
    distance: i32,
}

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let moves = parse_moves(input);
        println!("{:?}", moves);

        let mut tail: (i32, i32) = (0, 0);
        let mut head: (i32, i32) = (0, 0);
        let mut tail_visited: HashSet<(i32, i32)> = HashSet::from([(0, 0)]);

        // At the start of each move, the tail is the head of the previous move
        // The tail does not move, if the head is not far enough away from the tail (distance < 2)
        // Also diagonally

        for move_ in moves {
            for _ in 0..move_.distance {
                let (tail_x, tail_y) = tail;
                let (head_x, head_y) = head;

                match move_.direction {
                    Direction::Right => {
                        head = (head_x + 1, head_y);
                    }
                    Direction::Left => {
                        head = (head_x - 1, head_y);
                    }
                    Direction::Up => {
                        head = (head_x, head_y + 1);
                    }
                    Direction::Down => {
                        head = (head_x, head_y - 1);
                    }
                }

                if !is_touching(tail, head) {
                    tail = (head_x, head_y); // Move the tail to the head before the move
                    tail_visited.insert(tail);
                }
            }

            println!("tail: {:?}, head: {:?}", tail, head);
        }
        tail_visited.insert(tail);
        println!("tail_visited: {:?}", tail_visited);
        println!("tail_visited: {:?}", tail_visited.len());
    }
    fn solve2(&self, input: String) {}
}

impl Direction {
    fn from_char(c: char) -> Direction {
        match c {
            'R' => Direction::Right,
            'L' => Direction::Left,
            'U' => Direction::Up,
            'D' => Direction::Down,
            _ => panic!("Unknown direction"),
        }
    }
}

fn parse_moves(input: String) -> Vec<Move> {
    let moves = input
        .lines()
        .map(|line| {
            let instructions = line.split_ascii_whitespace().collect::<Vec<&str>>();
            let direction = instructions[0].chars().next().unwrap();
            let distance = instructions[1].parse::<i32>().unwrap();
            Move {
                direction: Direction::from_char(direction),
                distance,
            }
        })
        .collect::<Vec<Move>>();
    return moves;
}

fn is_touching(tail: (i32, i32), head: (i32, i32)) -> bool {
    let (head_x, head_y) = head;
    let around: HashSet<(i32, i32)> = HashSet::from([
        (head_x, head_y),         // No move
        (head_x, head_y + 1),     //up
        (head_x, head_y - 1),     //down
        (head_x + 1, head_y),     //right
        (head_x - 1, head_y),     //left
        (head_x + 1, head_y + 1), //up right
        (head_x - 1, head_y - 1), //down left
        (head_x - 1, head_y + 1), //up left
        (head_x + 1, head_y - 1), //down right
    ]);

    around.contains(&tail)
}
