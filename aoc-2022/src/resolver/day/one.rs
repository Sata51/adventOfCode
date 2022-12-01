use crate::resolver::challenge::ChallengeResolver;

pub struct D1P1;
pub struct D1P2;

impl ChallengeResolver for D1P1 {
    fn handle(&self, input: String) {
        // Split input by blank lines
        let lines: Vec<&str> = input.split("\n\n").collect();
        println!("Lines: {:?}", lines);
        // for every group, split by lines, convert the line to number an sum it
        let sums: Vec<i32> = lines
            .iter()
            .map(|line| {
                line.split("\n")
                    .map(|line| line.parse::<i32>().unwrap())
                    .sum::<i32>()
            })
            .collect();
        println!("Sum: {:?}", sums);
        // Get the max
        let max = sums.iter().max().unwrap();
        println!("Max: {:?}", max);
    }
}

impl ChallengeResolver for D1P2 {
    fn handle(&self, input: String) {
        // Split input by blank lines
        let lines: Vec<&str> = input.split("\n\n").collect();
        println!("Lines: {:?}", lines);
        // for every group, split by lines, convert the line to number an sum it
        let sums: Vec<i32> = lines
            .iter()
            .map(|line| {
                line.split("\n")
                    .map(|line| line.parse::<i32>().unwrap())
                    .sum::<i32>()
            })
            .collect();

        println!("Sums: {:?}", sums);
        // Sort the sums
        let mut sorted_sums = sums.clone();
        sorted_sums.sort();
        // Reverse the sorted sums
        sorted_sums.reverse();
        println!("Sorted sums: {:?}", sorted_sums);
        // Sum the 3 first elements
        let top3sum = sorted_sums[0] + sorted_sums[1] + sorted_sums[2];
        println!("top3sum: {:?}", top3sum);
    }
}
