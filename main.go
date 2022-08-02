package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/curusarn/escape-game-text-rpg/direction"
)

const help = `
Move around to expore with 'go <direction>'.
Directions: 'north'/'east'/'south'/'west'
Show this help with 'help'.
Exit with 'exit'.
`

var xPos int = 0
var yPos int = 0

const xLen int = 5
const yLen int = 5

var gameMap [][]int = [][]int{
	{0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 1, 1, 1, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}

func main() {
	fmt.Println("You are standing in the middle of nowhere.\nType your next action.")
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

func lookAround() string {
	return fmt.Sprintf("When you look North: %s\n", peekStr(direction.North)) +
		fmt.Sprintf("When you look East: %s\n", peekStr(direction.East)) +
		fmt.Sprintf("When you look South: %s\n", peekStr(direction.South)) +
		fmt.Sprintf("When you look West: %s\n", peekStr(direction.West))
}

func peekStr(dir direction.Direction) string {
	str, _ := peek(dir)
	return str
}
func peekBool(dir direction.Direction) bool {
	_, free := peek(dir)
	return free
}

func peek(dir direction.Direction) (string, bool) {
	x := xPos + dir.GetXDelta()
	y := yPos + dir.GetYDelta()

	if x < 0 || y < 0 || x >= xLen || y >= yLen {
		return "You see the edge of the map.", false
	}
	val := gameMap[y][x]
	switch val {
	case 0:
		return "You see free space.", true
	case 1:
		return "You see a wall.", false
	default:
		return "You see PANIC!", false
	}
}

func move(dir direction.Direction) string {
	var msg string

	if !peekBool(dir) {
		msg = "You cannot go that way!"
	} else {
		xPos += dir.GetXDelta()
		yPos += dir.GetYDelta()
		switch dir {
		case direction.North:
			msg = "You went north."
		case direction.South:
			msg = "You went south."
		case direction.East:
			msg = "You went east."
		case direction.West:
			msg = "You went west."
		}
	}
	return msg + "\n" + lookAround()
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
		dir, err := direction.ParseDirection(words[1])
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
		dir, err := direction.ParseDirection(words[0])
		if err != nil {
			return "I don't know what that means. Maybe try shouting for help?"
		}
		return move(dir)
	}

	panic("No action matched")
}
