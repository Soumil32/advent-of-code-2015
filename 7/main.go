package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	add func(command instruction)
	found func(target string)
}

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
	wires := make(map[string]uint16, len(input))
	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")
		queue := createQueue(&wires)
		switch len(command) {
		case 3:
			instruction := instruction{command[len(command)-1], command, []string{command[0]}}
			threeLenCommand(instruction, &wires, &queue)
		case 4:
			instruction := instruction{command[len(command)-1], command, []string{command[1]}}
			singleInputCommand(instruction, &wires, &queue)
		case 5:
			instruction := instruction{command[len(command)-1], command, []string{command[0], command[2]}}
			twoInputCommand(instruction, &wires, &queue)
		default:
			panic("Something went wrong")
		}
	}
	fmt.Println(wires)
	fmt.Printf("The total power provided to wire 'a' is %d\n", wires["a"])
}

func threeLenCommand(instruction instruction, pWires *map[string]uint16, pQueue *queue) {
	wires := *pWires
	var err error
	num, err := strconv.Atoi(instruction.command[0])
	wires[instruction.target] = uint16(num)
	if err != nil { // it was an assignment from another wire. eg: lx -> b
		_, ok := wires[instruction.command[0]]
		if !ok {
			queue := *pQueue
			queue.add(instruction)
		}
	}
}

func singleInputCommand(instruction instruction, pWires *map[string]uint16, pQueue *queue) {
	wires := *pWires
	operation := instruction.command[0]
	input, ok := wires[instruction.command[1]]
	queue := *pQueue
	if !ok {
		queue.add(instruction)
	}
	switch operation {
	case "NOT":
		result := ^input
		wires[instruction.target] = result
		queue.found(instruction.target)
	}
}

func twoInputCommand(instruction instruction, pWires *map[string]uint16, queue *queue) {
	wires := *pWires
	operation := instruction.command[1]
	var input1 uint16 = 0
	if condition, err := strconv.Atoi(instruction.command[0]); err == nil {
		input1 = uint16(condition)
	} else {
		input1 = wires[instruction.command[0]]
	}
	var input2 uint16 = 0
	if condition, err := strconv.Atoi(instruction.command[2]); err == nil {
		input2 = uint16(condition)
	} else {
		input2 = wires[instruction.command[2]]
	}
	if input1 == 0 || input2 == 0 { // "A gate provides no signal until all of its inputs have a signal."
		return
	}
	switch operation {
	case "AND":
		wires[instruction.target] = input1 & input2
	case "OR":
		wires[instruction.target] = input1 | input2
	case "LSHIFT":
		wires[instruction.target] = input1 << input2
	case "RSHIFT":
		wires[instruction.target] = input1 >> input2
	}
}

func createQueue(pWires *map[string]uint16) (queue) {
	commands := make(map[string]instruction) // target -> command
	wires := *pWires
	var myQueue queue

	add := func (instruction instruction) {
		commands[instruction.target] = instruction
	}
	found := func (target string) {
		if _, ok := commands[target]; !ok {
			return
		}
		instruction := commands[target]
		for _, dependency := range instruction.dependencies {
			_, ok := wires[dependency]
			if !ok {
				return // we haven't found all the dependencies yet
			}
		}
		command := commands[target]
		switch len(command.command) {
		case 3:
			threeLenCommand(command, pWires, &myQueue)
		case 4:
			singleInputCommand(command, pWires, &myQueue)
		case 5:
			twoInputCommand(command, pWires, &myQueue)
		default:
			panic("Something went wrong")
		}
	}
	myQueue = queue{add, found,}
	return myQueue
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}