use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let mut cpu = CPU::new(vec![20, 60, 100, 140, 180, 220]);

        for line in input.lines() {
            cpu.process_instruction(line);
        }

        let mut sum = 0;
        for i in 0..cpu.value_at_significant.len() {
            sum += cpu.value_at_significant[i] * cpu.significant_cycles[i];
        }
        println!("sum: {}", sum);
    }
    fn solve2(&self, input: String) {
        let mut cpu = CPU::new(vec![20, 60, 100, 140, 180, 220]);

        for line in input.lines() {
            cpu.process_instruction(line);
        }

        for (i, pixel) in cpu.screen.iter().enumerate() {
            if i % 40 == 0 {
                println!();
            }
            print!("{}", pixel);
        }
        println!();
    }
}

struct CPU {
    register: i32, // registrar value
    cycles: i32,   // current cycle
    significant_cycles: Vec<i32>,
    value_at_significant: Vec<i32>,
    screen: Vec<char>,
}

impl CPU {
    fn new(significant_cycles: Vec<i32>) -> CPU {
        CPU {
            register: 1,
            cycles: 1,
            significant_cycles,
            value_at_significant: Vec::new(),
            screen: vec!['.'; 240],
        }
    }
    fn process_instruction(&mut self, line: &str) {
        let mut op = line.split_ascii_whitespace();

        // Begin instruction
        self.check_cycle();
        self.cycles += 1; // Read op

        match op.next() {
            Some("noop") => {
                return;
            }
            Some("addx") => {
                self.check_cycle();
                self.cycles += 1; // Read arg

                // To cycle done, increase register
                self.register += op.next().unwrap().parse::<i32>().unwrap();
                return;
            }
            _ => unreachable!("Unknown instruction: {}", line),
        }
    }

    fn check_cycle(&mut self) {
        if self.significant_cycles.contains(&self.cycles) {
            self.value_at_significant.push(self.register);
        }
        let screen_pix = self.cycles % 40;

        if screen_pix == self.register
            || screen_pix == self.register + 1
            || screen_pix == self.register + 2
        {
            self.screen[(self.cycles - 1) as usize] = '#';
        }
    }
}
