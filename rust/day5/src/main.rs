use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();

    part_one(&input);
    part_two(&input);
}

fn part_one(input: &str) {
    let mut nice_strings: u32 = 0;

    for string_on_list in input.split("\n") {
        if contains_three_vowels(string_on_list) && contains_two_letters_in_a_row(string_on_list) && ["ab", "cd", "pq", "xy"].iter().all(|&x| !string_on_list.contains(x)) {
            nice_strings += 1;
        }
    }
    println!("Part 1: There are {} nice strings", nice_strings);
}

fn part_two(input: &str) {
    let mut nice_strings: u32 = 0;

    for string_on_list in input.split("\n") {
        if contains_pattern_twice(string_on_list) && (2..string_on_list.len()).any(|letter_index| {
            string_on_list[letter_index..letter_index+1] == string_on_list[letter_index-2..letter_index-1]
        }) {
            nice_strings += 1;
        }
    }
    println!("Part 2: There are {} nice strings", nice_strings);
}

fn contains_pattern_twice(s: &str) -> bool {
    for i in 1..s.len() {
        if s[i+1..].contains(&s[i-1..i+1]) {
            return true
        }
    }
    false
}

fn contains_two_letters_in_a_row(s: &str) -> bool {
    let mut last_char: char = ' ';
    for c in s.chars() {
        if c == last_char {
            return true;
        }
        last_char = c;
    }
    false
}

fn contains_three_vowels(s: &str) -> bool {
    let mut vowel_count: u32 = 0;

    for c in s.chars() {
        match c {
            'a' | 'e' | 'i' | 'o' | 'u' => vowel_count += 1,
            _ => (),
        }
    }

    vowel_count >= 3
}