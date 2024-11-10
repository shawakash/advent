use std::{collections::HashMap, fs};

fn main() -> std::io::Result<()> {
    let contents = fs::read_to_string("input.txt").unwrap();

    let mut string_to_num = HashMap::new();
    string_to_num.insert("one", 1);
    string_to_num.insert("two", 2);
    string_to_num.insert("three", 3);
    string_to_num.insert("four", 4);
    string_to_num.insert("five", 5);
    string_to_num.insert("six", 6);
    string_to_num.insert("seven", 7);
    string_to_num.insert("eight", 8);
    string_to_num.insert("nine", 9);

    let mut sum = 0;
    for line in contents.lines() {
        sum += get_digit(line, &string_to_num);
    }
    println!("{}", sum);

    Ok(())
}

fn get_digit(line: &str, string_to_num: &HashMap<&str, u32>) -> u32 {
    let mut first_digit = None;
    let mut last_digit = None;

    for i in 0..line.len() {
        if line[i..i + 1].chars().next().unwrap().is_digit(10) {
            let digit = line[i..i + 1].parse::<u32>().unwrap();
            if first_digit.is_none() {
                first_digit = Some(digit);
            }
            last_digit = Some(digit);
        }

        for (&word, &num) in string_to_num.iter() {
            if line[i..].starts_with(word) {
                if first_digit.is_none() {
                    first_digit = Some(num);
                }
                last_digit = Some(num);
            }
        }
    }

    let first = first_digit.unwrap();
    let last = last_digit.unwrap();
    first * 10 + last
}
