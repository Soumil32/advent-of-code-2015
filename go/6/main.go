package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type coordinate struct {
	x, y int
}

const (
	OFF = 0
	ON = 1
)

func main() { 
	input := strings.Split(readFile("input.txt"), "\n")
	start := time.Now()
	partOne(input)
	fmt.Printf("Part 1 took %s\n", time.Since(start))
	start = time.Now()
	partTwo(input)
	fmt.Printf("Part 2 took %s\n", time.Since(start))
}

func partOne(input []string) {
	// store the lights state
	lights := make(map[coordinate]int, 1_000_000)
	
	// parse input
	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ");
		start := convertStringToCoordinates(command[len(command)-3])
		end := convertStringToCoordinates(command[len(command)-1])
		cmdLen := len(command)
		
		for x := start.x; x <= end.x; x++ { // start and end are inclusive
			for y := start.y; y <= end.y; y++ {
				currentCoordinate := coordinate{x, y}
				if cmdLen == 4 {
					lights[currentCoordinate] ^= ON
					continue
				}
				if command[1] == "on" {
					lights[currentCoordinate] = ON
				} else {
					lights[currentCoordinate] = OFF
				}
			}
		}
		}
		lightsOn := 0
		for _, v := range lights {
			if v == ON {
				lightsOn++
			}
		}

		fmt.Printf("Part 1: The amount of lights on is %d\n", lightsOn)
	}

func partTwo(input []string) {
	// store the lights state
	lights := make(map[coordinate]int, 1_000_000)
	
	// parse input
	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ");
		var start coordinate
		var end coordinate
		var operation string
		if len(command) == 4 {
			start = convertStringToCoordinates(command[1])
			end = convertStringToCoordinates(command[3])
			operation = command[0]
		} else if len(command) == 5 {
			start = convertStringToCoordinates(command[2])
			end = convertStringToCoordinates(command[4])
			operation = command[0] + " " + command[1]
		}
		for x := start.x; x <= end.x; x++ {
			for y := start.y; y <= end.y; y++ {
				currentCoordinate := coordinate{x, y}
				switch operation {
				case "toggle":
					lights[currentCoordinate] += 2
				case "turn on":
					lights[currentCoordinate]++
				case "turn off":
					if lights[currentCoordinate] > 0 {
						lights[currentCoordinate]--
					}
				default:
					panic("unrecognised operation")
				}
			}
		}
	}
	var totalBrightness int = 0
	for _, v := range lights {
		totalBrightness += v
	}
	fmt.Printf("Part 2: The total brightness of all the lights is %d\n", totalBrightness)
}

func filterLightsMap(m map[coordinate]int, filter func(any) bool) (map[coordinate]int) {
	newMap := map[coordinate]int{}
	for key, value := range m {
		if filter(value) {
			newMap[key] = value
		}
	}
	return newMap
}

func convertStringToCoordinates(s string) coordinate {
	coorsString := strings.Split(s, ",") 
	x, err := strconv.Atoi(coorsString[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coorsString[1])
	if err != nil {
		panic(err)
	}
	return coordinate{x, y}
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}