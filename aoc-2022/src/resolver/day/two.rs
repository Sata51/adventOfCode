use crate::resolver::challenge::ChallengeResolver;

pub struct D2P1;
pub struct D2P2;

impl ChallengeResolver for D2P1 {
    fn handle(&self, input: String) {}
}

impl ChallengeResolver for D2P2 {
    fn handle(&self, input: String) {}
}
