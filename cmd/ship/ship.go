package main

import (
	"github.com/curusarn/escape-game-text-rpg/game"
	"github.com/curusarn/escape-game-text-rpg/position"
	"github.com/curusarn/escape-game-text-rpg/terrain"
)

func main() {
	e := terrain.Terrain{
		PeakMsg:     "the edge of the world",
		CantMoveMsg: "You cannot sail to the edge of the world.",
	}
	var openSea string = "an open sea"
	o := terrain.Terrain{
		PeakMsg: openSea,
	}
	x := terrain.Terrain{
		PeakMsg:  openSea,
		DeathMsg: "You have sailed onto a mine, your ship exploded into a thousand pieces and you died.",
	}
	arr := [][]terrain.Terrain{
		{x, x, x, o, o, o},
		{x, o, o, o, o, o},
		{x, o, o, o, o, o},
		{x, o, o, o, o, o},
		{x, o, o, o, o, o},
	}
	gameMap := game.BuildGameMapFromSlice(arr)
	intro := `
You are a captain of a small ship in the middle of the sea.
Where should we sail next?
Type your next command.
	`
	help := `
Sail around with 'north'/'east'/'south'/'west'
Short version of directions also works: n/e/s/w
Show this help with 'help'.
Use 'look' to look around.
Exit with 'exit'.
	`
	g := game.Game{
		GameMap:        gameMap,
		PlayerPosition: position.Position{X: 2, Y: 2},
		NothingTerrain: e,
		MoveMsg:        "You have sailed ",
		Intro:          intro,
		Help:           help,
	}
	g.Start()
}
