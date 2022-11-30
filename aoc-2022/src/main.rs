use std::path::Path;

use aoc_2022::Resolver;
use clap::Parser;

#[derive(Parser, Debug)]
#[clap(
    version = "1.0",
    author = "Virgile MATHIEU<vm.mathieu@gmail.com>",
    about = "Advent of code 2022"
)]
struct Args {
    #[clap(short = 'd', long = "day", default_value = "1")]
    day: u8,

    #[clap(short = 'p', long = "part", default_value = "1")]
    part: u8,

    #[clap(short = 'i', long = "input", default_value = "input.txt")]
    input: String,
}

fn main() {
    let args = Args::parse();
    println!(
        "Day: {}, Part: {}, Input: {}",
        args.day, args.part, args.input
    );

    let input_file =
        std::fs::read_to_string(Path::new(&args.input)).expect("Unable to open input file");

    let resolver = Resolver::new(args.day, args.part, input_file);
    resolver.resolve();
}
