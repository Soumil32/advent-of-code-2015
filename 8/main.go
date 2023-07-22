package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

func main()  {
	input := strings.Split(readFile("input.txt"), "\n")
	partOne(input)
	partTwo(input)
}

func partOne(input []string) {
	charCodeVsMemory := 0
	for _, value := range input {
		charInCode := len(value)
		inMemory, err := strconv.Unquote(value)
		if err != nil {
			panic(err)
		}
		charInMemory := len(inMemory)
		charCodeVsMemory += charInCode - charInMemory
	}
	fmt.Printf("Part 1: There are %d fewer characters in memory\n", charCodeVsMemory)
}

func partTwo(input []string) {
	difference := 0
	for _, value := range input {
		difference += len(strconv.Quote(value)) - len(value)
	}
	fmt.Printf("Part 2: There are %d less characters in the original strings compared to the encoded string\n", difference)
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}