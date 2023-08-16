use std::borrow::Cow;
use std::fs;
use std::collections::HashMap;

#[cfg(feature = "dhat-heap")]
#[global_allocator]
static ALLOC: dhat::Alloc = dhat::Alloc;

fn main() {
    #[cfg(feature = "dhat-heap")]
    let _profiler = dhat::Profiler::new_heap();
    let input = fs::read_to_string("input.txt").unwrap();

    part_one(&input);
}

fn part_one(input: &str) {
    let mut wires_to_find: Vec<Vec<&str>> = input.lines().map(|item| {item.split(" ").collect()}).collect();
    let mut wires: HashMap<&str, u16> = HashMap::with_capacity(wires_to_find.len());
    let mut wires_to_find_next = Vec::new();

    'outer: while wires_to_find.len() > 0 {
        for line in wires_to_find {
            let succesful;
            match line.len() {
                3 => {
                    succesful = assignment_command(&line, &mut wires);
                },
                4 => {
                    succesful = not_command(&line, &mut wires);
                },
                5 => {
                    succesful = binary_command(&line, &mut wires);
                },
                _ => {
                    panic!("Unknown command")
                }
            }
            if !succesful {
                wires_to_find_next.push(line);
            } else if let Some(_) = wires.get("a") {
                break 'outer;
            }
        }
        wires_to_find = wires_to_find_next.clone();
        wires_to_find_next.clear();
    }
    println!("Part 1: wire 'a' has the signal {}", wires["a"]);
}

fn binary_command<'a>(line: &Vec<&'a str>, wires: &mut HashMap<&'a str, u16>) -> bool {
    let wire = line[4];
    let value1_str = line[0];
    let value2_str = line[2];
    let value1;
    let value2;
    if !value1_str.chars().all(char::is_numeric) {
        let wire_found = wires.get(value1_str);
        match wire_found {
            Some(res) => {
                value1 = *res;
            },
            None => {
                return false
            }
        }
    } else {
        value1 = value1_str.parse::<u16>().unwrap();
    }
    if !value2_str.chars().all(char::is_numeric) {
        let wire_found = wires.get(value2_str);
        match wire_found {
            Some(res) => {
                value2 = *res;
            },
            None => {
                return false
            }
        }
    } else {
        value2 = value2_str.parse::<u16>().unwrap();
    }
    match line[1] {
        "AND" => {
            wires.insert(wire, value1 & value2);
        },
        "OR" => {
            wires.insert(wire, value1 | value2);
        },
        "LSHIFT" => {
            wires.insert(wire, value1 << value2);
        },
        "RSHIFT" => {
            wires.insert(wire, value1 >> value2);
        },
        _ => {
            panic!("Unknown command")
        }
    }
    true
}

fn not_command<'a>(line: &'a Vec<&str>, wires: &mut HashMap<&'a str, u16>) -> bool {
    let wire = line[3];
    let value = line[1];
    if !value.chars().all(char::is_numeric) {
        let wire_found = wires.get(value);
        match wire_found {
            Some(res) => {
                wires.insert(wire, !res);
                return true
            },
            None => {
                return false
            }
        }
    }
    wires.insert(wire, !value.parse::<u16>().unwrap());
    true
}

fn assignment_command<'a>(line: &Vec<&'a str>, wires: &mut HashMap<&'a str, u16>) -> bool {
    let wire = line[2];
    let value = line[0];
    if !value.chars().all(char::is_numeric) {
        let wire_found = wires.get(value);
        match wire_found {
            Some(res) => {
                wires.insert(wire, *res);
                return true
            },
            None => {
                return false
            }
        }
    }
    wires.insert(wire, value.parse().unwrap());
    true
}