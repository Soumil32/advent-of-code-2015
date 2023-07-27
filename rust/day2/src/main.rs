use std::{fs};

fn main() {
    let input = fs::read_to_string("input.txt").expect("Could not read input file");
    let input = input.split("\n").collect();

    part_one(&input);
    part_two(&input)
}

fn part_one(input: &Vec<&str>) {
    let mut total_wrapping_paper: u32 = 0;
    for present in input {
        let dimensions: Vec<u32> = present.split("x").map(|dimension: &str| {
            let dimension: u32 = dimension.parse().expect("Could not read one of the numbers");
            dimension
        }).collect();

        let mut sides = [dimensions[0]*dimensions[1], dimensions[1]*dimensions[2], dimensions[0]*dimensions[2]];
        sides.sort();

        total_wrapping_paper += sides[0] * 3 + sides[1] * 2 + sides[2] * 2
    }
    println!("Part 1: The elves need a total of {} square feet of wrapping paper", total_wrapping_paper)
}

fn part_two(input: &Vec<&str>) {
    let mut total_ribbon: u32 = 0;
    for present in input {
        let mut dimensions: Vec<u32> = present.split("x").map(|dimension: &str| {
            let dimension: u32 = dimension.parse().expect("Could not read one of the numbers");
            dimension
        }).collect();

        dimensions.sort();

        total_ribbon += dimensions[0] * 2 + dimensions[1] * 2 + dimensions.iter().fold(1, |total, dimension| {total * dimension});
    }
    println!("Part 2: Santa's Elves need {} feet of ribbon", total_ribbon);
}