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
	partTwo(input)
}

func partOne(input []string) {
	// store the lights state
	lights := make(map[coordinate]int, 1_000_000)
	
	// parse input
	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ");
		var start coordinate
		var end coordinate
		if len(command) == 4 {
			start = convertStringToCoordinates(command[1])
			end = convertStringToCoordinates(command[3])
		} else if len(command) == 5 {
			start = convertStringToCoordinates(command[2])
			end = convertStringToCoordinates(command[4])
		}
		
		for x := start.x; x <= end.x; x++ { // start and end are inclusive
			for y := start.y; y <= end.y; y++ {
				currentCoordinate := coordinate{x, y}
				if len(command) == 4 {
					light := lights[currentCoordinate]
					if light == OFF {
						lights[currentCoordinate] = ON
					} else {
						lights[currentCoordinate] = OFF
					}
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
		lightsOn := filterLightsMap(lights, 
			func (value any) bool {
				return value == ON
			},
		)
		fmt.Printf("Part 1: The amount of lights on is %d\n", len(lightsOn))
	}

func partTwo(input []string) {
	
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