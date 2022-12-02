use std::str::Split;

use itertools::{IntoChunks, Itertools};

use crate::resolver::challenge::ChallengeResolver;

#[derive(PartialEq)]
enum RPS {
    Rock,
    Paper,
    Scissors,
}

impl RPS {
    fn to_int(&self) -> i32 {
        match self {
            RPS::Rock => 1,
            RPS::Paper => 2,
            RPS::Scissors => 3,
        }
    }
}

#[derive(PartialEq)]
enum RPSResult {
    Win,
    Lose,
    Draw,
}

impl RPSResult {
    fn to_int(&self) -> i32 {
        match self {
            RPSResult::Win => 6,
            RPSResult::Lose => 0,
            RPSResult::Draw => 3,
        }
    }
}

struct Game {
    MyChoice: RPS,
    OpponentChoice: RPS,
}

impl Game {
    fn new(my_choice: RPS, opponent_choice: RPS) -> Game {
        Game {
            MyChoice: my_choice,
            OpponentChoice: opponent_choice,
        }
    }

    fn new_from_str(my_choice: &str, opponent_choice: &str) -> Game {
        let my_choice = match my_choice {
            "X" => RPS::Rock,
            "Y" => RPS::Paper,
            "Z" => RPS::Scissors,
            _ => panic!("Invalid choice"),
        };

        let opponent_choice = match opponent_choice {
            "A" => RPS::Rock,
            "B" => RPS::Paper,
            "C" => RPS::Scissors,
            _ => panic!("Invalid choice"),
        };

        Game {
            MyChoice: my_choice,
            OpponentChoice: opponent_choice,
        }
    }

    fn get_result(&self) -> RPSResult {
        let my_choice = self.MyChoice.to_int();
        let opponent_choice = self.OpponentChoice.to_int();

        if my_choice == opponent_choice {
            return RPSResult::Draw;
        }

        if my_choice == 1 && opponent_choice == 3 {
            return RPSResult::Win;
        }

        if my_choice == 2 && opponent_choice == 1 {
            return RPSResult::Win;
        }

        if my_choice == 3 && opponent_choice == 2 {
            return RPSResult::Win;
        }

        return RPSResult::Lose;
    }

    fn get_round_points(&self) -> i32 {
        return self.get_result().to_int() + self.MyChoice.to_int();
    }
}

// Oponent value
// A for Rock
// B for Paper
// C for Scissors

// My value
// X Rock
// Y Paper
// Z Scissors

// If lost add 0
// If draw add 3
// If win add 6

pub struct D2P1;
pub struct D2P2;

impl ChallengeResolver for D2P1 {
    fn handle(&self, input: String) {
        let mut my_value = RPS::Rock;
        let mut oponent_value = RPS::Rock;
        let mut total = 0;
        for c in input.chars() {
            let mut roundValue = 0;
            let mut roundResult = RPSResult::Draw;
            match c {
                'A' => {
                    oponent_value = RPS::Rock;
                }
                'B' => {
                    oponent_value = RPS::Paper;
                }
                'C' => {
                    oponent_value = RPS::Scissors;
                }
                'X' => {
                    my_value = RPS::Rock;
                }
                'Y' => {
                    my_value = RPS::Paper;
                }
                'Z' => {
                    my_value = RPS::Scissors;
                }
                _ => {}
            }
            if my_value == oponent_value {
                roundResult = RPSResult::Draw;
            } else if my_value == RPS::Rock && oponent_value == RPS::Scissors {
                roundResult = RPSResult::Win;
            } else if my_value == RPS::Paper && oponent_value == RPS::Rock {
                roundResult = RPSResult::Win;
            } else if my_value == RPS::Scissors && oponent_value == RPS::Paper {
                roundResult = RPSResult::Win;
            } else {
                roundResult = RPSResult::Lose;
            }

            total += roundResult.to_int();
            total += my_value.to_int();

            println!("{:?} ", total);
        }
        println!("{:?} ", total);
    }
}

impl ChallengeResolver for D2P2 {
    fn handle(&self, input: String) {}
}

fn get_rounds(input: String) -> Vec<Game> {
    return input
        .split("\n\n")
        .map(|game| {
            game.split(" ")
                .chunks(2)
                .into_iter()
                .map(|p| Game::new_from_str(p[0], p[1]))
                .collect()
        })
        .flatten()
        .collect();
}
