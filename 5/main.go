package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	input := strings.Split(readFile("input.txt"), "\n")
	fmt.Printf("The orinal amount of strings is %d\n", len(input))
	start := time.Now()
	partOne(input)
	fmt.Printf("Time taken for part one: %s\n", time.Since(start))
	partTwo(input)
}

func partOne(input []string) {
	var niceStrings uint = 0;
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, str := range input {
		if containsVowels(str, 3) && hasLetterTwiceInARow(str) && !containsForbidden(str, forbidden){
			niceStrings++
		}
	}
	fmt.Printf("Part 1: The amount of nice strings are %d\n", niceStrings)
}

func partTwo(input []string) {
	niceStrings := 0
	for i := 0; i < len(input); i++ {
		if containsAPhraseTwice(input[i]) && hasRepeatsWithLetterInBetween(input[i]) {
			niceStrings++
		}
	}
	fmt.Printf("Part 2: The amount of nice strings are %d\n", niceStrings)
}

func containsAPhraseTwice(s string) bool {
	for i := 1; i < len(s); i++ {
		lastTwoLetters := s[i-1:i+1]
		if strings.Contains(s[i+1:], lastTwoLetters) {
			return true
		}
	}
	return false
}

func hasRepeatsWithLetterInBetween(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}
	return false
}

func hasLetterTwiceInARow(s string) bool {
	var prevLetter byte;
	for i := 0; i < len(s); i++ {
		if s[i] == prevLetter {
			return true
		}
		prevLetter = s[i]
	}
	return false
}

func containsVowels(s string, amount int) bool {
	vowels := "aeiou"
	amountFound := 0
	for _, letter := range strings.Split(s, "") {
		if strings.Contains(vowels, letter) {
			amountFound++
		}
	}
	if amountFound >= amount {
		return true
	} else {
		return false
	}
}

func containsForbidden(s string, forbidden []string) bool {
	for i := 0; i < len(forbidden); i++ {
		if strings.Contains(s, forbidden[i]) {
			return true
		}
	}
	return false
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}