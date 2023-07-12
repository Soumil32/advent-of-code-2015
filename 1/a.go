package main

import (
	"fmt"
	"os"
	"strings"
)

/* 
	PROBLEM DESCRIPTION
	-------------------
	Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character 
	at a time. An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

	For example:
		(()) and ()() both result in floor 0.
		((( and (()(()( both result in floor 3.
		))((((( also results in floor 3.
		()) and ))( both result in floor -1 (the first basement level).
		))) and )())()) both result in floor -3.
*/

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// To what floor do the instructions take santa?
	input_bytes, err:= os.ReadFile("input.txt")
	check(err) //easy
	brackets := strings.Split(string(input_bytes), "")

	var currentFloor int;
	for _, bracket := range(brackets) {
		if bracket=="(" {
			currentFloor++
		} else {
			currentFloor--
		}
	} 
	fmt.Printf("Part 1: Santa is currently at floor %d\n", currentFloor)
}

func partTwo() {
	/*
	Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

	For example:
		) causes him to enter the basement at character position 1.
		()()) causes him to enter the basement at character position 5.
		
	What is the position of the character that causes Santa to first enter the basement?
	*/
	input_bytes, err:= os.ReadFile("input.txt")
	check(err)
	brackets := strings.Split(string(input_bytes), "")

	var currentFloor int;
	for i, bracket := range(brackets) {
		switch bracket {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		default:
			panic("Unrecognised character")
		}
		if currentFloor == -1 {
			fmt.Printf("Part 2: Santa entered the basement at index %d\n", i+1) // the first bracket has position 1 but i starts from 0
			break;
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}