use std::collections::HashMap;

use super::challenge::{Challenge, ChallengeResolver};

use super::day::{D1P1, D1P2, D2P1, D2P2};

pub struct Resolver {
    day: u8,
    part: u8,
    input: String,
    resolver: HashMap<Challenge, Box<dyn ChallengeResolver>>,
}

impl Resolver {
    pub fn new(day: u8, part: u8, input: String) -> Resolver {
        let mut resolvers: HashMap<Challenge, Box<dyn ChallengeResolver>> = HashMap::new();
        // Day 1
        resolvers.insert(Challenge::new(1, 1), Box::new(D1P1));
        resolvers.insert(Challenge::new(1, 2), Box::new(D1P2));
        // Day 2
        resolvers.insert(Challenge::new(2, 1), Box::new(D2P1));
        resolvers.insert(Challenge::new(2, 2), Box::new(D2P2));

        Resolver {
            day,
            part,
            input,
            resolver: resolvers,
        }
    }

    pub fn resolve(&self) {
        println!(
            "Working on Day: {}, Part: {}, Input: {}",
            self.day,
            self.part,
            self.input.len()
        );

        let challenge = Challenge::new(self.day, self.part);
        let resolver = self.resolver.get(&challenge).unwrap();

        resolver.handle(self.input.clone());
    }
}
