use std::collections::HashSet;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        println!("Visible: {}", count_visible(&parse_trees(&input)));
    }
    fn solve2(&self, input: String) {
        println!("Scenic score: {}", get_scenic_score(&parse_trees(&input)));
    }
}

fn parse_trees(input: &str) -> Vec<Vec<usize>> {
    input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap() as usize)
                .collect()
        })
        .collect()
}

fn count_visible(tree_map: &Vec<Vec<usize>>) -> usize {
    let mut h: HashSet<(usize, usize)> = HashSet::new();

    for y in 0..tree_map.len() {
        for x in 0..tree_map[0].len() {
            // Exteriors
            if x == 0 || x == tree_map[0].len() - 1 || y == 0 || y == tree_map.len() - 1 {
                h.insert((x, y));
            } else {
                // We are in the interior
                let left = tree_map[y][0..x].iter().copied().rev().collect::<Vec<_>>();
                let right = tree_map[y][x + 1..tree_map[0].len()].to_vec();

                let top = tree_map[0..y].iter().map(|r| r[x]).collect::<Vec<_>>();
                let bottom = tree_map[y + 1..tree_map.len()]
                    .iter()
                    .map(|r| r[x])
                    .collect::<Vec<_>>();

                // visible from one direction
                if left.iter().all(|h| *h < tree_map[y][x])
                    || right.iter().all(|h| *h < tree_map[y][x])
                    || top.iter().all(|h| *h < tree_map[y][x])
                    || bottom.iter().all(|h| *h < tree_map[y][x])
                {
                    h.insert((x, y));
                }
            }
        }
    }

    h.len()
}

fn get_tree_scenic_score(value: usize, direction: &Vec<usize>) -> usize {
    let mut score = 0;
    for v in direction {
        if value > *v {
            score += 1;
        } else {
            return score + 1;
        }
    }
    score
}

fn get_scenic_score(tree_map: &Vec<Vec<usize>>) -> usize {
    let mut score: Vec<Vec<usize>> = Vec::new();

    for y in 0..tree_map.len() {
        let mut tmp_score: Vec<usize> = Vec::new();
        for x in 0..tree_map[0].len() {
            let left = tree_map[y][0..x].iter().copied().rev().collect::<Vec<_>>();
            let right = tree_map[y][x + 1..tree_map[0].len()].to_vec();

            let top = tree_map[0..y]
                .iter()
                .map(|r| r[x])
                .rev()
                .collect::<Vec<_>>();
            let bottom = tree_map[y + 1..tree_map.len()]
                .iter()
                .map(|r| r[x])
                .collect::<Vec<_>>();

            tmp_score.push(
                get_tree_scenic_score(tree_map[y][x], &right)
                    * get_tree_scenic_score(tree_map[y][x], &left)
                    * get_tree_scenic_score(tree_map[y][x], &top)
                    * get_tree_scenic_score(tree_map[y][x], &bottom),
            );
        }

        score.push(tmp_score);
    }

    *score.iter().map(|r| r.iter().max().unwrap()).max().unwrap()
}
