use itertools::Itertools;

use crate::resolver::challenge::ChallengeResolver;

pub struct D1P1;
pub struct D1P2;

impl ChallengeResolver for D1P1 {
    fn handle(&self, input: String) {
        // Split input by blank lines
        let max = get_elves_total_calories(input).into_iter().max().unwrap();
        println!("Max: {:?}", max);
    }
}

impl ChallengeResolver for D1P2 {
    fn handle(&self, input: String) {
        // Split input by blank lines
        let top3sum: i32 = get_elves_total_calories(input)
            .into_iter()
            .sorted()
            .rev()
            .take(3)
            .sum::<i32>();
        println!("top3sum: {:?}", top3sum);
    }
}

fn get_elves_total_calories(input: String) -> Vec<i32> {
    return input
        .split("\n\n")
        .map(|line| {
            line.split("\n")
                .map(|line| line.parse::<i32>().unwrap())
                .sum::<i32>()
        })
        .collect();
}
