use std::fs;
use md5::Digest;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Error reading input.txt");

    part_one(&input);
    part_two(&input);
}

fn part_one(input: &str){
    for i in 0.. {
        let hash = md5::compute(format!("{}{}", input, i));
        // the last 5 characters of the hexidecimal hash must be 0
        if format!("{:x}", hash).starts_with("00000") {
            println!("Part 1: {} produces a hash starting with 5 zeros", i);
            break;
        }
    }
}

fn part_two(input: &str){
    for i in 0.. {
        let hash: Digest = md5::compute(format!("{}{}", input, i));
        // the last 5 characters of the hexidecimal hash must be 0
        if format!("{:x}", hash).starts_with("000000")  {
            println!("Part 2: {} produces a hash starting with 6 zeros", i);
            break;
        }
    }
}