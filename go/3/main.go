// https://adventofcode.com/2015/day/3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type coordinate struct {
	x int
	y int
}

type position struct {
	currentPosition, previousPosition coordinate;
}

func main() {
	input := strings.Split(readTextFile("input.txt"), "")
	start := time.Now()
	partOne(input)
	fmt.Printf("Part 1 took %s time to run\n", time.Since(start))
	afterFirst := time.Now() 
	partTwo(input)
	fmt.Printf("Part 2 took %s time to run\n", time.Since(afterFirst))
}

func partOne(input []string) {
	housesVisited := []coordinate{{0, 0}};
	previousHouse := coordinate{0, 0}
	for _,symbol := range input {
		directionToMove := getDirection(symbol)
		currentHouse := coordinate{previousHouse.x + directionToMove.x, previousHouse.y + directionToMove.y}
		if visited := isHouseVisited(currentHouse, housesVisited); !visited {
			housesVisited = append(housesVisited, currentHouse)
		}
		previousHouse = currentHouse
	}
	fmt.Println(len(housesVisited))
}

func partTwo(input []string) {
	housesVisited := []coordinate{{0, 0}}
	santa := position{coordinate{0, 0}, coordinate{0, 0}}
	robot := position{coordinate{0, 0}, coordinate{0, 0}}
	for i, symbol := range input {
		direction := getDirection(symbol)
		if i % 2 == 0 { // santa does every even move
			santa.currentPosition = coordinate{santa.previousPosition.x + direction.x, santa.previousPosition.y + direction.y}
			if visted := isHouseVisited(santa.currentPosition, housesVisited); !visted {
				housesVisited = append(housesVisited, santa.currentPosition)
			}
			santa.previousPosition = santa.currentPosition
		} else { // robo-santa does every odd move
			robot.currentPosition = coordinate{robot.previousPosition.x + direction.x, robot.previousPosition.y + direction.y}
			if visted := isHouseVisited(robot.currentPosition, housesVisited); !visted {
				housesVisited = append(housesVisited, robot.currentPosition)
			}
			robot.previousPosition = robot.currentPosition
		}
	}
	fmt.Println(len(housesVisited))
}

func isHouseVisited(currHouse coordinate, housesVisited []coordinate) bool {
	for _, coordinate := range housesVisited {
		if coordinate == currHouse {
			return true
		}
	}
	return false
}

func getDirection(symbol string) coordinate {
	switch symbol {
	case ">":
		return coordinate{1, 0}
	case "<":
		return coordinate{-1, 0}
	case "^":
		return coordinate{0, 1}
	case "v":
		return coordinate{0, -1}
	default:
		panic("unknown symbol in input")
	}
}

func readTextFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}