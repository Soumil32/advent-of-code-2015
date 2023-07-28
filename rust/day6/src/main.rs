use std::fs;
use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq, Hash)]
struct Coordinate {
    x: i32,
    y: i32,
}

impl Coordinate {
    fn from_string(s: &str) -> Coordinate {
        let mut coordinate = s.split(",");
        Coordinate {
            x: coordinate.next().unwrap().parse().unwrap(),
            y: coordinate.next().unwrap().parse().unwrap(),
        }
    }
}

enum Action {
    Toggle,
    TurnOn,
    TurnOff
}

fn main() {
    let input = fs::read_to_string("input.txt").expect("Could not read input.txt");

    part_one(&input);
    part_two(&input);
}

fn part_one(input: &str) {
    let mut lights: HashMap<Coordinate, u8> = HashMap::with_capacity(1000 * 1000);
    // pre-fill the lights with 0
    for x in 0..1000 {
        for y in 0..1000 {
            lights.insert(Coordinate { x, y }, 0);
        }
    }
    
    for line in input.split("\n") {
        let line: Vec<&str> = line.split(" ").collect();
        let start = Coordinate::from_string(&line[line.len() - 3]);
        let end = Coordinate::from_string(&line[line.len() - 1]);

        let action = get_action_type(&line);

        for x in start.x..=end.x {
            for y in start.y..=end.y {
                let current_coordinate = Coordinate {x, y};
                let current_value = lights[&current_coordinate];
                match action {
                    Action::Toggle => {
                        lights.insert(current_coordinate, current_value ^ 1);
                    },
                    Action::TurnOff => {lights.insert(current_coordinate, 0);},
                    Action::TurnOn => {lights.insert(current_coordinate, 1);}
                }
            }
        }
    }
    let lights_on = lights.values().fold(0, |total: u32, cur| {total + *cur as u32});
    println!("Part 1: There are {} lights on", lights_on);
}


fn part_two(input: &str) {
    let mut lights: HashMap<Coordinate, u8> = HashMap::with_capacity(1000 * 1000);
    // pre-fill the lights with 0
    for x in 0..1000 {
        for y in 0..1000 {
            lights.insert(Coordinate { x, y }, 0);
        }
    }

    for line in input.split("\n") {
        let line: Vec<&str> = line.split(" ").collect();
        let start = Coordinate::from_string(&line[line.len() - 3]);
        let end = Coordinate::from_string(&line[line.len() - 1]);

        let action = get_action_type(&line);

        for x in start.x..=end.x {
            for y in start.y..=end.y {
                let current_coordinate = Coordinate {x, y};
                let current_value = lights[&current_coordinate];
                match action {
                    Action::Toggle => { lights.insert(current_coordinate, current_value + 2); },
                    Action::TurnOff => { lights.insert(current_coordinate, current_value.saturating_sub(1)); },
                    Action::TurnOn => { lights.insert(current_coordinate, current_value + 1); }
                }
            }
        }
    }
    let brightness = lights.values().fold(0, |total: u32, cur| {total + *cur as u32});
    println!("Part 2: The total brightness is {}", brightness);
}

fn get_action_type(line: &Vec<&str>) -> Action {
    match line[1] {
        "on" => Action::TurnOn,
        "off" => Action::TurnOff,
        _ => Action::Toggle
    }
}
