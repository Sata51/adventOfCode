use array_tool::vec::Intersect;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let sum = get_ruck_sacs(input)
            .into_iter()
            .map(|r| get_priority(r.common_letter))
            .sum::<i32>();
        println!("sum {:?}", sum);
    }
    fn solve2(&self, input: String) {
        let sum = get_ruck_sacs(input)
            .chunks(3)
            .into_iter()
            .map(|r| get_common_letter_in_group(r))
            .map(|r| get_priority(r))
            .sum::<i32>();
        println!("sum {:?}", sum);
    }
}

fn get_priority(c: char) -> i32 {
    return c as i32 - if c.is_lowercase() { 96 } else { 38 };
}

#[derive(Debug)]
struct RuckSack {
    str: String, // Raw input
    pub common_letter: char,
}

impl RuckSack {
    fn new(input: String) -> RuckSack {
        //Chunk input in half
        let chars: Vec<char> = input.chars().collect::<Vec<_>>();
        let char_len = chars.clone().len() / 2;

        let mut rucksack = RuckSack {
            str: input,
            common_letter: ' ',
        };

        let mut base = chars.clone().into_iter().take(char_len).collect::<Vec<_>>();
        base = base.intersect(chars.into_iter().rev().take(char_len).collect::<Vec<_>>());

        // Base should only have one char
        if base.len() != 1 {
            panic!("Base should only have one char");
        }

        rucksack.common_letter = base[0];

        rucksack
    }
}

fn get_ruck_sacs(input: String) -> Vec<RuckSack> {
    return input
        .split("\n")
        .map(|line| RuckSack::new(line.to_string()))
        .collect();
}

fn get_common_letter_in_group(r_group: &[RuckSack]) -> char {
    let elf_letters: Vec<Vec<char>> = r_group
        .iter()
        .map(|r| r.str.chars().collect::<Vec<_>>())
        .collect();

    let mut base = elf_letters[0].clone();
    base = base.intersect(elf_letters[1].clone());
    base = base.intersect(elf_letters[2].clone());

    if base.len() > 1 {
        panic!("More than one common letter");
    }
    return base[0];
}
