package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// xPos := 0
// yPos := 0
const help = `
Move around to expore with 'go <direction>'.
Show this help with 'help'.
Exit with 'exit'.
`

type Direction int64

const (
	North Direction = iota
	South
	East
	West
)

func parseDirection(str string) (Direction, error) {
	switch str {
	case "n":
		fallthrough
	case "north":
		return North, nil

	case "s":
		fallthrough
	case "south":
		return South, nil

	case "e":
		fallthrough
	case "east":
		return East, nil

	case "w":
		fallthrough
	case "west":
		return West, nil
	}
	return North, errors.New("Unknown direction")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		resp := execInput(input)
		fmt.Println(resp)
	}
}

func move(dir Direction) string {
	var msg string

	switch dir {
	case North:
		msg = "You went north."
	case South:
		msg = "You went south."
	case East:
		msg = "You went east."
	case West:
		msg = "You went west."
	}
	return msg + "\nYou see wastelands all around you"
}

func execInput(input string) string {
	input = strings.TrimSuffix(input, "\n")
	words := strings.Split(input, " ")
	length := len(words)

	switch words[0] {
	case "go":
		if length != 2 {
			return "go where?"
		}
		dir, err := parseDirection(words[1])
		if err != nil {
			return "I don't know that direction"
		}
		return move(dir)
	case "exit":
		os.Exit(0)
	case "what":
		fallthrough
	case "help":
		return help
	default:
		if length != 1 {
			return "what do you want to do?"
		}
		dir, err := parseDirection(words[0])
		if err != nil {
			return "I don't know what that means. Maybe try shouting fo help?"
		}
		return move(dir)
	}

	panic("No action matched")
}
