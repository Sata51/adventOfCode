use std::collections::HashMap;

use super::challenge::ChallengeResolver;

use super::day::{Solver1, Solver2, Solver3, Solver4, Solver5};

pub struct Resolver {
    day: u8,
    part: u8,
    input: String,
    resolver: HashMap<u8, Box<dyn ChallengeResolver>>,
}

impl Resolver {
    pub fn new(day: u8, part: u8, input: String) -> Resolver {
        let mut resolvers: HashMap<u8, Box<dyn ChallengeResolver>> = HashMap::new();
        resolvers.insert(1, Box::new(Solver1));
        resolvers.insert(2, Box::new(Solver2));
        resolvers.insert(3, Box::new(Solver3));
        resolvers.insert(4, Box::new(Solver4));
        resolvers.insert(5, Box::new(Solver5));

        Resolver {
            day,
            part,
            input,
            resolver: resolvers,
        }
    }

    pub fn resolve(&self) {
        println!("Working on Day: {}, Part: {}", self.day, self.part);

        let resolver = self
            .resolver
            .get(&self.day)
            .expect(format!("No resolver found for day {}", self.day).as_str());

        match self.part {
            1 => resolver.solve1(self.input.clone()),
            2 => resolver.solve2(self.input.clone()),
            _ => panic!("No resolver found for part {}", self.part),
        }
    }
}
