use crate::resolver::challenge::ChallengeResolver;

#[derive(PartialEq, Clone, Copy)]
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
    my_choice: RPS,
    opponent_choice: RPS,
}

impl Game {
    fn new_from_str(opponent_choice: &str, my_choice: &str) -> Game {
        let my_choice = match my_choice {
            "X" => RPS::Rock,
            "Y" => RPS::Paper,
            "Z" => RPS::Scissors,
            _ => panic!("Invalid choice m:{:?}", my_choice),
        };

        let opponent_choice = match opponent_choice {
            "A" => RPS::Rock,
            "B" => RPS::Paper,
            "C" => RPS::Scissors,
            _ => panic!("Invalid choice o:{:?}", opponent_choice),
        };

        Game {
            my_choice,
            opponent_choice,
        }
    }

    fn get_result(&self) -> RPSResult {
        let my_choice = self.my_choice.to_int();
        let opponent_choice = self.opponent_choice.to_int();

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
        return self.get_result().to_int() + self.my_choice.to_int();
    }

    fn get_rounds_points_for_expected_result(&mut self) -> i32 {
        // Mychoice will be updated based on the expected result
        // "X" => RPS::Rock, but means "I expect to lose"
        // "Y" => RPS::Paper, but means "I expect to draw"
        // "Z" => RPS::Scissors, but means "I expect to win"

        if self.my_choice == RPS::Paper {
            print!("Draw expected, ");
            self.my_choice = self.opponent_choice.clone(); // Draw
            return self.get_round_points(); // Get the new result
        }

        if self.my_choice == RPS::Rock {
            print!("Loose expected, ");
            // I need to loose so i get the choice based on the opponent choice
            self.my_choice = match self.opponent_choice {
                RPS::Rock => RPS::Scissors,
                RPS::Paper => RPS::Rock,
                RPS::Scissors => RPS::Paper,
            }; // Lose
            return self.get_round_points(); // Get the new result
        }

        if self.my_choice == RPS::Scissors {
            print!("Win expected, ");
            // I need to win so i get the choice based on the opponent choice
            self.my_choice = match self.opponent_choice {
                RPS::Rock => RPS::Paper,
                RPS::Paper => RPS::Scissors,
                RPS::Scissors => RPS::Rock,
            }; // Win
            return self.get_round_points(); // Get the new result
        }

        return self.get_round_points();
    }
}

// Opponent value
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
        let mut total = 0;
        for c in get_rounds(input) {
            total += c.get_round_points();
        }
        println!("{:?} ", total);
    }
}

impl ChallengeResolver for D2P2 {
    fn handle(&self, input: String) {
        let mut total = 0;
        for mut c in get_rounds(input) {
            total += c.get_rounds_points_for_expected_result();

            println!("{:?} ", total);
        }
        println!("{:?} ", total);
    }
}

fn get_rounds(input: String) -> Vec<Game> {
    return input
        .split("\n")
        .filter(|line| line.len() > 0)
        .map(|game| {
            // For every line, split it into two parts
            let mut parts = game.split(" ");
            let my_choice = parts.next().unwrap();
            let opponent_choice = parts.next().unwrap();
            return Game::new_from_str(my_choice, opponent_choice);
        })
        .collect();
}
