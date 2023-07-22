package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type instruction struct {
	target string
	command []string
	dependencies []string
}

func main() {
	input := strings.Split(readFile("input.txt"), "\n")
	/*input := strings.Split(`123 -> x
	456 -> y
	x AND y -> d
	x OR y -> e
	x LSHIFT 2 -> f
	y RSHIFT 2 -> g
	NOT x -> h
	NOT y -> i`, "\n")*/
	partOne(input)
}

func partOne(input []string) {
	
}

func threeLenCommand() {
	
}

func singleInputCommand() {
	
}

func twoInputCommand() {
	
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}