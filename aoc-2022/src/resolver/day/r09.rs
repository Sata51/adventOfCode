use std::collections::HashSet;

use itertools::Itertools;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let mut rp = Rope::new(1);
        for _move in parse_moves(input) {
            rp.handle_move(_move);
        }
        println!("tail_visited: {:?}", rp.tail_visited.len());
    }
    fn solve2(&self, input: String) {
        let mut rp = Rope::new(9);
        for _move in parse_moves(input) {
            rp.handle_move(_move);
        }
        println!("tail_visited: {:?}", rp.tail_visited.len());
    }
}

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

impl Move {
    fn from_line(s: &str) -> Move {
        let instructions = s.split_ascii_whitespace().collect::<Vec<&str>>();
        Move {
            direction: Direction::from_char(instructions[0].chars().next().unwrap()),
            distance: instructions[1].parse::<i32>().unwrap(),
        }
    }
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy)]
struct RopeElement {
    x: i32,
    y: i32,
}

impl RopeElement {
    fn new() -> RopeElement {
        RopeElement { x: 0, y: 0 }
    }
}

struct Rope {
    parts: Vec<RopeElement>,
    tail_visited: HashSet<RopeElement>,
}

impl Rope {
    fn new(knots: i32) -> Rope {
        let mut parts = Vec::from([RopeElement::new()]);
        for _ in 0..knots {
            parts.push(RopeElement::new());
        }
        Rope {
            parts,
            tail_visited: HashSet::from([RopeElement::new()]), // Insert the origin
        }
    }

    fn handle_move(&mut self, move_: Move) {
        for _ in 0..move_.distance {
            // Move the head
            match move_.direction {
                Direction::Right => self.parts[0].x += 1,
                Direction::Left => self.parts[0].x -= 1,
                Direction::Up => self.parts[0].y += 1,
                Direction::Down => self.parts[0].y -= 1,
            }

            // Iterate over the tails
            for (head_id, tail_id) in (0..self.parts.len()).tuple_windows() {
                let diff = RopeElement {
                    x: self.parts[head_id].x - self.parts[tail_id].x,
                    y: self.parts[head_id].y - self.parts[tail_id].y,
                };

                let should_move = diff.x.abs() > 1 || diff.y.abs() > 1;

                if should_move {
                    self.parts[tail_id].x += diff.x.signum();
                    self.parts[tail_id].y += diff.y.signum();
                    if tail_id == self.parts.len() - 1 {
                        self.tail_visited.insert(self.parts[self.parts.len() - 1]);
                    }
                }
            }
        }
    }
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
        .map(|line| Move::from_line(line))
        .collect::<Vec<Move>>();
    return moves;
}
