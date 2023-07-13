package main

import (
	"fmt";
	"os";
	"strings";
	"strconv"
	"sort"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := strings.Split(readTextFile("./input.txt"), "\n");
	var totalWrappingPaper int;
	for _, box := range(input) {
		dimensions := convertStringsToInts(strings.Split(box, "x"))
		// Equation to find wrapping paper needed - 2*l*w + 2*w*h + 2*h*l + SMALLEST_SIDE
		// data is in format l,w,h
		surfaceAreas := []int{dimensions[0]*dimensions[1], dimensions[1]*dimensions[2], dimensions[2]*dimensions[0]}
		smallestArea := surfaceAreas[0]
		for _, dimension := range(surfaceAreas) {
			if dimension < smallestArea {
				smallestArea = dimension
			}
		}
		totalWrappingPaper += 2*surfaceAreas[0] + 2*surfaceAreas[1] + 2*surfaceAreas[2] + smallestArea
	}
	defer fmt.Println(totalWrappingPaper)
}

func partTwo() {
	input := strings.Split(readTextFile("./input.txt"), "\n");
	var totalRibbon int;
	for _, box := range(input) {
		dimensions := convertStringsToInts(strings.Split(box, "x"))
		sort.Ints(dimensions)
		// 
		// data is in format l,w,h
		totalRibbon += dimensions[0]*2 + dimensions[1]*2 + (dimensions[0] * dimensions[1] * dimensions[2])

	}
	fmt.Println(totalRibbon)
}

func convertStringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, str := range strings {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func readTextFile(path string) string {
	file_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return string(file_bytes)
}