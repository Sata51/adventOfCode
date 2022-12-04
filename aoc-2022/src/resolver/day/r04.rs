use std::ops::Range;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let ranges = get_ranges(input);
        let mut total_overlap = 0;
        for i in 0..ranges.len() {
            let (elf1range, elf2range) = ranges[i].clone();
            if full_overlap(elf1range.clone(), elf2range.clone()) {
                total_overlap += 1;
            }
        }
        println!("overlap : {:?}", total_overlap);
    }

    fn solve2(&self, input: String) {
        let ranges = get_ranges(input);
        let mut total_overlap = 0;
        for i in 0..ranges.len() {
            let (elf1range, elf2range) = ranges[i].clone();
            if overlap(elf1range.clone(), elf2range.clone()) {
                total_overlap += 1;
            }
        }
        println!("overlap : {:?}", total_overlap);
    }
}

fn get_ranges(input: String) -> Vec<(Vec<u32>, Vec<u32>)> {
    input
        .lines()
        .map(|line| {
            let mut parts = line.split(',');
            let range_1 = parts.next().unwrap();
            let range_2 = parts.next().unwrap();
            (get_one_range(range_1), get_one_range(range_2))
        })
        .collect()
}

fn get_one_range(input: &str) -> Vec<u32> {
    let mut parts = input.split('-');
    let start = parts.next().unwrap().parse::<u32>().unwrap();
    let end = parts.next().unwrap().parse::<u32>().unwrap();
    Range {
        start,
        end: end + 1,
    }
    .collect()
}

// Check if two vectors fully overlap in any way
fn full_overlap(a: Vec<u32>, b: Vec<u32>) -> bool {
    if a.iter().all(|x| b.contains(x)) {
        return true; // Overlap in one way
    }
    if b.iter().all(|x| a.contains(x)) {
        return true; // Overlap in the other way
    }
    false
}

fn overlap(a: Vec<u32>, b: Vec<u32>) -> bool {
    if a.iter().any(|x| b.contains(x)) {
        return true; // Overlap in one way
    }
    if b.iter().any(|x| a.contains(x)) {
        return true; // Overlap in the other way
    }
    false
}
