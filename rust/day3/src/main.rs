use std::fs;

#[derive(PartialEq, Clone)]
struct Coordinate {
    x: i32,
    y: i32
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("could not read input file");

    part_one(&input);
    part_two(&input);
}

fn part_one(input: &String) {
    let mut houses_visited: Vec<Coordinate> = vec![Coordinate{x: 0, y: 0}];
    let mut current_position = Coordinate {x: 0, y: 0};

    for direction in input.chars() {
        match direction {
            '^' => current_position.y += 1,
            '>' => current_position.x += 1,
            'v' => current_position.y -= 1,
            '<' => current_position.x -= 1,
            _ => continue
        }
        if !houses_visited.contains(&current_position) {
            houses_visited.insert(houses_visited.len(), current_position.clone())
        }
    }
    println!("Part 1: {} houses got atleast one present", houses_visited.len())
}

fn part_two(input: &String) {
    let mut houses_visited: Vec<Coordinate> = vec![Coordinate{x: 0, y: 0}];
    let mut santa_current_position = Coordinate{x:0, y: 0};
    let mut robo_current_position = Coordinate{x:0, y: 0};

    for (i, direction) in input.chars().enumerate() {
        let current_person = if i % 2 == 0 { &mut santa_current_position } else { &mut robo_current_position };
        match direction {
            '^' => current_person.y += 1,
            '>' => current_person.x += 1,
            'v' => current_person.y -= 1,
            '<' => current_person.x -= 1,
            _ => continue
        }
        if !houses_visited.contains(current_person) {
            houses_visited.insert(houses_visited.len(), current_person.clone())
        }
    }
    println!("Part 2: {} houses got atleast one present", houses_visited.len())
}
