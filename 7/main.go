package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	add func(command []string)
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
			threeLenCommand(command, &wires)
		case 4:
			singleInputCommand(command, &wires, &queue)
		case 5:
			twoInputCommand(command, &wires, &queue)
		default:
			panic("Something went wrong")
		}
	}
	fmt.Println(wires)
	fmt.Printf("The total power provided to wire 'a' is %d\n", wires["a"])
}

func threeLenCommand(command []string, pWires *map[string]uint16) {
	target := command[len(command)-1]
	wires := *pWires
	var err error
	num, err := strconv.Atoi(command[0])
	wires[target] = uint16(num)
	if err != nil { // it was an assignment from another wire. eg: lx -> b
		wires[target] = wires[command[0]]
	}
}

func singleInputCommand(command []string, pWires *map[string]uint16, pQueue *queue) {
	target := command[len(command)-1]
	wires := *pWires
	operation := command[0]
	input, ok := wires[command[1]]
	queue := *pQueue
	if !ok {
		queue.add(command)
	}
	switch operation {
	case "NOT":
		result := ^input
		wires[target] = result
		queue.found(target)
	}
}

func twoInputCommand(command []string, pWires *map[string]uint16, queue *queue) {
	target := command[len(command)-1]
	wires := *pWires
	operation := command[1]
	var input1 uint16 = 0
	if condition, err := strconv.Atoi(command[0]); err == nil {
		input1 = uint16(condition)
	} else {
		input1 = wires[command[0]]
	}
	var input2 uint16 = 0
	if condition, err := strconv.Atoi(command[2]); err == nil {
		input2 = uint16(condition)
	} else {
		input2 = wires[command[2]]
	}
	if input1 == 0 || input2 == 0 { // "A gate provides no signal until all of its inputs have a signal."
		return
	}
	switch operation {
	case "AND":
		wires[target] = input1 & input2
	case "OR":
		wires[target] = input1 | input2
	case "LSHIFT":
		wires[target] = input1 << input2
	case "RSHIFT":
		wires[target] = input1 >> input2
	}
}

func createQueue(pWires *map[string]uint16) (queue) {
	commands := make(map[string][]string) // target -> command
	// wires := *pWires
	var myQueue queue

	add := func (command []string) {
		commands[command[len(command)-1]] = command
	}
	found := func (target string) {
		if _, ok := commands[target]; !ok {
			return
		}
		command := commands[target]
		switch len(command) {
		case 3:
			threeLenCommand(command, pWires)
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