package game

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/curusarn/escape-game-text-rpg/position"
	"github.com/curusarn/escape-game-text-rpg/terrain"
)

type ActionFuncType func(*Game) (bool, string)

type Game struct {
	GameMap        map[position.Position]terrain.Terrain
	PlayerPosition position.Position
	NothingTerrain terrain.Terrain
	MoveMsg        string
	Intro          string
	Help           string
	GameState      int
	ActionStr      string
	ActionFunc     ActionFuncType

	LevelName    string
	ScreenLog    io.Writer
	StatusLog    io.Writer
	statusLogger *log.Logger
}

func LogPrintf(w io.Writer, format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Fprintf(w, format, args...)
}

func (g *Game) Printf(format string, args ...interface{}) {
	LogPrintf(g.ScreenLog, format, args...)
}

func (g *Game) InitLogger() {
	pid := os.Getpid()
	prefix := fmt.Sprintf("[pid: %d, %s] ", pid, g.LevelName)
	g.statusLogger = log.New(g.StatusLog, prefix, log.Default().Flags())
	g.statusLogger.Printf("Game started\n")
}

func (g *Game) LogStatus() {
	t := g.GetTerrainOnPosition(g.PlayerPosition)
	g.statusLogger.Printf("Player position: [%d,%d]; Terrain: %s\n", g.PlayerPosition.X, g.PlayerPosition.Y, t.PeakMsg)
}

func (g *Game) Start() int {
	g.InitLogger()
	g.Printf("%s\n", g.Intro)

	reader := bufio.NewReader(os.Stdin)
	for {

		g.Printf("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			g.Printf("ERROR: Could not read input: %v\n", err)
		}
		fmt.Fprint(g.ScreenLog, input)

		// Handle the execution of the input.
		ok, resp := g.HandleInput(input)
		g.Printf("%s\n", resp)
		g.LogStatus()
		if ok == false {
			g.statusLogger.Printf("Game ended with GameState: %d\n", g.GameState)
			return g.GameState
		}
	}
}

func (g *Game) GetTerrainOnPosition(p position.Position) terrain.Terrain {
	if val, ok := g.GameMap[p]; ok {
		return val
	}
	return g.NothingTerrain
}

func (g *Game) MovePlayer(d position.Direction) (bool, string) {
	newPos := g.PlayerPosition.Move(d)
	newTer := g.GetTerrainOnPosition(newPos)
	if !newTer.IsFree() {
		return false, newTer.CantMoveMsg + "\n"
	}
	g.PlayerPosition = newPos
	return true, ""
}

func (g *Game) Peak(d position.Direction) string {
	t := g.GetTerrainOnPosition(g.PlayerPosition.Move(d))
	return fmt.Sprintf("When you look %s you see %s.\n", d.ToString(), t.PeakMsg)
}

func (g *Game) LookAround() string {
	return g.Peak(position.DirectionNorth) +
		g.Peak(position.DirectionEast) +
		g.Peak(position.DirectionSouth) +
		g.Peak(position.DirectionWest)
}

func (g *Game) SetGameState() {
	t := g.GetTerrainOnPosition(g.PlayerPosition)
	if t.GameState != 0 {
		g.GameState = t.GameState
	}
}

func (g *Game) EvalPlayerPosition() (bool, string) {
	defer g.SetGameState()
	t := g.GetTerrainOnPosition(g.PlayerPosition)
	if t.IsDeadly() {
		return false, t.DeathMsg
	}
	return true, g.LookAround()
}

func (g *Game) HandleInput(input string) (bool, string) {
	input = strings.TrimSuffix(input, "\n")
	words := strings.Split(input, " ")
	length := len(words)

	switch words[0] {
	case "exit":
		return false, "Exited."
	case "help":
		return true, g.Help
	case "look":
		return true, g.LookAround()
	default:
		if length != 1 {
			return true, "You need to type something.\n" + g.Help
		}
		if g.ActionStr != "" && words[0] == g.ActionStr {
			return g.ActionFunc(g)
		}
		dir, err := position.ParseDirection(words[0])
		if err != nil {
			return true, "I don't know what that means.\n" + g.Help
		}
		if ok, msg := g.MovePlayer(dir); !ok {
			return true, msg
		} else {
			moveMsg := g.MoveMsg + dir.ToString() + "\n"
			ok, msg := g.EvalPlayerPosition()
			if !ok {
				return false, msg
			}
			return true, moveMsg + msg
		}
	}
}

func BuildGameMapFromSlice(arr [][]terrain.Terrain) map[position.Position]terrain.Terrain {
	gameMap := map[position.Position]terrain.Terrain{}
	for y, row := range arr {
		for x, t := range row {
			gameMap[position.Position{X: x, Y: y}] = t
		}
	}
	return gameMap
}
