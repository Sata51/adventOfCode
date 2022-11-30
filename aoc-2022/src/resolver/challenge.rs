use std::hash::{Hash, Hasher};

#[derive(Debug, PartialEq, Eq)]
pub struct Challenge {
    day: u8,
    part: u8,
}

impl Challenge {
    pub fn new(day: u8, part: u8) -> Challenge {
        Challenge { day, part }
    }
}

impl Hash for Challenge {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.day.hash(state);
        self.part.hash(state);
    }
}

pub trait ChallengeResolver {
    fn handle(&self, input: String); // This should print the result of whatever the challenge is
}
