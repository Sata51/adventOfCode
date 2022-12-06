use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        for line in input.lines() {
            let rslt = start_of_any(line, 4);
            println!("{}", rslt);
        }
    }
    fn solve2(&self, input: String) {
        for line in input.lines() {
            let rslt = start_of_any(line, 14);
            println!("{}", rslt);
        }
    }
}

fn start_of_any(line: &str, window_size: usize) -> i32 {
    // Should detect the start of a packet
    for window in line.chars().collect::<Vec<_>>().windows(window_size) {
        // Is there any duplicate
        if window
            .iter()
            .collect::<std::collections::HashSet<_>>()
            .len()
            == window_size
        {
            // println!("Last: {}", last);
            // Get the first index of the last char in the window
            return (line.find(&window.iter().collect::<String>()).unwrap() + window_size) as i32;
        }
    }

    0
}
