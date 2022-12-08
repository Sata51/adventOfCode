use crate::resolver::challenge::ChallengeResolver;

pub struct Solver;

impl ChallengeResolver for Solver {
    fn solve1(&self, input: String) {
        let tree_map = parse_trees(&input);
        println!("{:?}", tree_map);

        let visible_from_outside = count_exteriors(&tree_map);
        println!("Visible from outside: {}", visible_from_outside);
        let visible_from_inside = count_interiors(&tree_map);
        println!("Visible from inside: {}", visible_from_inside);

        println!("Total: {}", visible_from_outside + visible_from_inside);
    }
    fn solve2(&self, input: String) {}
}

fn parse_trees(input: &str) -> Vec<Vec<usize>> {
    input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap() as usize)
                .collect()
        })
        .collect()
}

fn count_exteriors(tree_map: &Vec<Vec<usize>>) -> usize {
    let mut visible_from_outside = 0;
    // The first and last line are always visible
    // The first and last column are always visible
    visible_from_outside += tree_map[0].len();
    visible_from_outside += tree_map[tree_map.len() - 1].len();
    visible_from_outside += (tree_map.len() - 2) * 2;

    visible_from_outside
}

fn count_interiors(tree_map: &Vec<Vec<usize>>) -> usize {
    let mut visible_from_inside = 0;

    // Do not count the first and last line
    // Do not count the first and last column
    for y in 1..tree_map.len() - 1 {
        for x in 1..tree_map[0].len() - 1 {
            if is_tree_visible(tree_map, x, y) {
                println!("{}x{} is visible", x, y);
                visible_from_inside += 1;
            } else {
                println!("{}x{} is not visible", x, y);
            }
        }
    }

    visible_from_inside
}

fn is_tree_visible(tree_map: &Vec<Vec<usize>>, x: usize, y: usize) -> bool {
    let height_at_index = tree_map[y][x];

    // Collect the 4 directions in 4 vectors
    let mut from_top = tree_map.iter().map(|v| v[x]).collect::<Vec<_>>();
    let mut from_bottom = from_top.clone().into_iter().rev().collect::<Vec<_>>();

    from_top.truncate(x);
    from_bottom.truncate(from_bottom.len() - x - 1);

    let visible_from_top = from_top.iter().all(|&h| h < height_at_index);
    let visible_from_bottom = from_bottom.iter().all(|&h| h < height_at_index);

    let mut from_left = tree_map[y].clone();
    let mut from_right = from_left.clone().into_iter().rev().collect::<Vec<_>>();

    from_left.truncate(x);
    from_right.truncate(from_right.len() - x - 1);

    let visible_from_left = from_left.iter().all(|&h| h < height_at_index);
    let visible_from_right = from_right.iter().all(|&h| h < height_at_index);

    println!(
        "{}x{}:
        from_top: {:?} - visible_from_top {}
        from_bottom: {:?} - visible_from_bottom {}
        from_left: {:?} - visible_from_left {}
        from_right: {:?} - visible_from_right {}",
        x,
        y,
        from_top,
        visible_from_top,
        from_bottom,
        visible_from_bottom,
        from_left,
        visible_from_left,
        from_right,
        visible_from_right
    );

    return visible_from_top || visible_from_bottom || visible_from_left || visible_from_right;
}
