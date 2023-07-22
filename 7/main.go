package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type wires = map[string]uint16

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

func partTwo(giveToB uint16, input []string) {
	wires := make(wires, len(input))
	wires["b"] = giveToB
	wiresToBeFound := input
	for {
		wiresToBeFoundNext := []string{}
		for i := 0; i < len(wiresToBeFound); i++ {
			command := strings.Split(wiresToBeFound[i], " ")
			if command[len(command)-1] == "b" {
				continue
			}
			var succesful bool;
			switch len(command) {
			case 3:
				succesful = threeLenCommand(command, &wires)
			case 4:
				succesful = singleInputCommand(command, &wires)
			case 5:
				succesful = twoInputCommand(command, &wires)
			}
			if !succesful {
				wiresToBeFoundNext = append(wiresToBeFoundNext, wiresToBeFound[i])
			}
		}
		_, ok := wires["a"]
		if len(wiresToBeFoundNext) == 0 || ok {
			break
		}
		wiresToBeFound = wiresToBeFoundNext
	}
	fmt.Printf("Part 2: The wire 'a' now has the signal %d\n", wires["a"])
}

func partOne(input []string) {
	wires := make(wires, len(input))
	wiresToBeFound := input
	for {
		wiresToBeFoundNext := []string{}
		for i := 0; i < len(wiresToBeFound); i++ {
			command := strings.Split(wiresToBeFound[i], " ")
			var succesful bool;
			switch len(command) {
			case 3:
				succesful = threeLenCommand(command, &wires)
			case 4:
				succesful = singleInputCommand(command, &wires)
			case 5:
				succesful = twoInputCommand(command, &wires)
			}
			if !succesful {
				wiresToBeFoundNext = append(wiresToBeFoundNext, wiresToBeFound[i])
			}
		}
		_, ok := wires["a"]
		if len(wiresToBeFoundNext) == 0 || ok {
			break
		}
		wiresToBeFound = wiresToBeFoundNext
	}
	fmt.Printf("Part 1: The wire 'a' had the signal %d\n", wires["a"])
	partTwo(wires["a"], input)
}

func threeLenCommand(command []string, pWires *wires) bool {
	num, err := strconv.Atoi(command[0]);
	wires := *pWires
	if err != nil {
		// if there is an error that means the wire is being assigned from another wire
		wire, ok := wires[command[0]]
		if !ok {
			return false
		}
		wires[command[2]] = wire
		return true
	}
	wires[command[2]] = uint16(num)
	return true
}

func singleInputCommand(command []string, pWires *wires) bool {
	operation := command[0]
	wires := *pWires
	num, err := strconv.Atoi(command[1])
	if err != nil {
		// it is another wire
		wire, ok := wires[command[1]]
		if !ok {
			return false
		}
		num = int(wire)
	}
	switch operation {
	case "NOT":
		wires[command[3]] = ^uint16(num)
	}
	return true
}

func twoInputCommand(command []string, pWires *wires) bool {
	wires := *pWires
	num1, err := strconv.Atoi(command[0])
	if err != nil {
		wire, ok := wires[command[0]]
		if !ok {
			return false
		}
		num1 = int(wire)
	}
	num2, err := strconv.Atoi(command[2])
	if err != nil {
		wire, ok := wires[command[2]]
		if !ok {
			return false
		}
		num2 = int(wire)
	}
	operation := command[1]
	switch operation {
	case "AND":
		wires[command[len(command)-1]] = uint16(num1 & num2)
	case "OR":
		wires[command[len(command)-1]] = uint16(num1 | num2)
	case "RSHIFT":
		wires[command[len(command)-1]] = uint16(num1 >> num2)
	case "LSHIFT":
		wires[command[len(command)-1]] = uint16(num1 << num2)
	}
	return true
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}