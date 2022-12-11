use lazy_static::lazy_static;
use regex::Regex;

use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let mut monkeys = input
            .split("\n\n")
            .map(|s| Monkey::from_string(s))
            .collect::<Vec<Monkey>>();

        for _ in 0..20 {
            for index in 0..monkeys.len() {
                monkeys = take_turn(monkeys, index, i64::MAX);
            }
        }
        monkeys.sort_by_key(|m| -m.inspected);
        println!(
            "{:?}",
            monkeys.iter().map(|m| m.inspected).take(2).product::<i64>()
        );
    }
    fn solve2(&self, input: String) {
        let mut monkeys = input
            .split("\n\n")
            .map(|s| Monkey::from_string(s))
            .collect::<Vec<Monkey>>();

        let divider: i64 = monkeys.iter().map(|m| m.test).product();
        for _ in 0..10000 {
            for index in 0..monkeys.len() {
                monkeys = take_turn(monkeys, index, divider);
            }
        }
        monkeys.sort_by_key(|m| -m.inspected);
        println!(
            "{:?}",
            monkeys.iter().map(|m| m.inspected).take(2).product::<i64>()
        );
    }
}

fn take_turn(monkeys: Vec<Monkey>, index: usize, divider: i64) -> Vec<Monkey> {
    let mut result = monkeys;
    let mut monkey = result[index].clone();
    for item in &monkey.worry_levels {
        let mut level = monkey.operation.perform(*item);
        if divider == i64::MAX {
            level = (level as f64 / 3.0).floor() as i64;
        }
        result[monkey.worry_test(level) as usize]
            .worry_levels
            .push(level % divider);
    }
    monkey.inspected += monkey.worry_levels.len() as i64;
    monkey.worry_levels.clear();
    result[index] = monkey;
    result
}

lazy_static! {
    static ref RE_MONKEY_ID: Regex = Regex::new(r"Monkey (\d+):").unwrap();
    static ref RE_MONKEY_TEST: Regex = Regex::new(r"Test: divisible by (\d+)").unwrap();
    static ref RE_MONKEY_ON_SUCCESS: Regex = Regex::new(r"If true: throw to monkey (\d+)").unwrap();
    static ref RE_MONKEY_ON_ERROR: Regex = Regex::new(r"If false: throw to monkey (\d+)").unwrap();
}

#[derive(Debug, Clone)]
struct Monkey {
    worry_levels: Vec<i64>,
    operation: Operation,
    test: i64,
    on_success: i32, // Go to the monkey with this id
    on_error: i32,   // Go to the monkey with this id
    inspected: i64,
}

#[derive(Debug, Clone)]
enum Sign {
    Plus,
    Multiply,
}

#[derive(Debug, Clone)]
struct Operation {
    sign: Sign,
    pre_sign_old: bool,
    pre_sign_value: i64,
    post_sign_old: bool,
    post_sign_value: i64,
}

impl Monkey {
    fn from_string(input: &str) -> Monkey {
        let mut lines = input.lines();
        // Ignore the first line with the id
        lines.next();

        // Parse items
        let worry_levels = lines
            .next()
            .unwrap()
            .split_ascii_whitespace()
            .skip(2)
            .map(|s| {
                let mut s = s.to_string();
                if s.ends_with(",") {
                    s = s[..s.len() - 1].to_string();
                }
                s.parse::<i64>().unwrap()
            })
            .collect::<Vec<i64>>();

        // Parse operation
        let line_operation = lines.next().unwrap(); // TODO
                                                    // split the line on the = sign
        let mut splitted = line_operation.split(" = ");
        // The first part is not interesting
        splitted.next();
        // The second part is the operation
        let operation = Operation::from_string(splitted.next().unwrap());

        // Parse test
        let caps_test = RE_MONKEY_TEST.captures(lines.next().unwrap()).unwrap();
        let test = caps_test.get(1).unwrap().as_str().parse::<i64>().unwrap();

        // Parse on_success
        let caps_on_success = RE_MONKEY_ON_SUCCESS
            .captures(lines.next().unwrap())
            .unwrap();

        let on_success = caps_on_success
            .get(1)
            .unwrap()
            .as_str()
            .parse::<i32>()
            .unwrap();

        let caps_on_error = RE_MONKEY_ON_ERROR.captures(lines.next().unwrap()).unwrap();

        let on_error = caps_on_error
            .get(1)
            .unwrap()
            .as_str()
            .parse::<i32>()
            .unwrap();

        Monkey {
            worry_levels,
            operation,
            test,
            on_success,
            on_error,
            inspected: 0,
        }
    }

    fn worry_test(&self, worry_level: i64) -> i32 {
        if worry_level % self.test == 0 {
            self.on_success
        } else {
            self.on_error
        }
    }
}

impl Operation {
    fn from_string(input: &str) -> Operation {
        let mut result_operation = Operation {
            sign: Sign::Plus,
            pre_sign_old: false,
            pre_sign_value: 0,
            post_sign_old: false,
            post_sign_value: 0,
        };

        let mut splitted = input.split_ascii_whitespace();
        // Should be in 3 parts
        let pre_sign_part = splitted.next().unwrap();
        match pre_sign_part {
            "old" => result_operation.pre_sign_old = true,
            _ => result_operation.pre_sign_value = pre_sign_part.parse::<i64>().unwrap(),
        }
        let sign = splitted.next().unwrap();
        match sign {
            "+" => result_operation.sign = Sign::Plus,
            "*" => result_operation.sign = Sign::Multiply,
            _ => panic!("Unknown sign"),
        }
        let post_sign_part = splitted.next().unwrap();
        match post_sign_part {
            "old" => result_operation.post_sign_old = true,
            _ => result_operation.post_sign_value = post_sign_part.parse::<i64>().unwrap(),
        }

        result_operation
    }

    fn perform(&self, worry_level: i64) -> i64 {
        let pre_part = if self.pre_sign_old {
            worry_level
        } else {
            self.pre_sign_value
        };
        let post_part = if self.post_sign_old {
            worry_level
        } else {
            self.post_sign_value
        };

        if let Sign::Plus = self.sign {
            pre_part + post_part
        } else {
            pre_part * post_part
        }
    }
}
