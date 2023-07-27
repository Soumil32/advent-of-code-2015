use std::{fs, str::Chars, path::Path};

fn main() {
    let input: String = fs::read_to_string(Path::new("input.txt")).expect("Couldn't read input");

    let input: Chars<'_> = input.chars();

    part_one(&input);
    part_two(&input);
}

fn part_one(input: &Chars<'_>) {
    let mut current_floor: i32 = 0;
    for bracket in input.clone() {
        match bracket {
            '(' => current_floor+=1,
            ')' => current_floor-=1,
            _ => continue
        }
    }
    println!("Part 1: Santa ended up at floor {}", current_floor)
}

fn part_two(input: &Chars<'_>) {
    let mut current_floor: i32 = 0;
    let mut basement_index: i32 = 0;
    for (i, bracket) in input.clone().enumerate() {
        match bracket {
            '(' => current_floor+=1,
            ')' => current_floor-=1,
            _ => continue
        }
        if current_floor == -1 {
            basement_index = i as i32 + 1;
            break;
        }
    }
    println!("Part 2: Santa entered the basement as index {basement_index}")
}