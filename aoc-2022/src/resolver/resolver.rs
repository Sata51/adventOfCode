use std::collections::HashMap;

use super::challenge::ChallengeResolver;

use super::day::{
    Solver1, Solver10, Solver11, Solver12, Solver13, Solver14, Solver15, Solver16, Solver17,
    Solver18, Solver19, Solver2, Solver20, Solver21, Solver22, Solver23, Solver24, Solver25,
    Solver3, Solver4, Solver5, Solver6, Solver7, Solver8, Solver9,
};

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
        resolvers.insert(6, Box::new(Solver6));
        resolvers.insert(7, Box::new(Solver7));
        resolvers.insert(8, Box::new(Solver8));
        resolvers.insert(9, Box::new(Solver9));
        resolvers.insert(10, Box::new(Solver10));
        resolvers.insert(11, Box::new(Solver11));
        resolvers.insert(12, Box::new(Solver12));
        resolvers.insert(13, Box::new(Solver13));
        resolvers.insert(14, Box::new(Solver14));
        resolvers.insert(15, Box::new(Solver15));
        resolvers.insert(16, Box::new(Solver16));
        resolvers.insert(17, Box::new(Solver17));
        resolvers.insert(18, Box::new(Solver18));
        resolvers.insert(19, Box::new(Solver19));
        resolvers.insert(20, Box::new(Solver20));
        resolvers.insert(21, Box::new(Solver21));
        resolvers.insert(22, Box::new(Solver22));
        resolvers.insert(23, Box::new(Solver23));
        resolvers.insert(24, Box::new(Solver24));
        resolvers.insert(25, Box::new(Solver25));

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
