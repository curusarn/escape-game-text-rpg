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
	l := terrain.Terrain{
		PeakMsg:     "land",
		CantMoveMsg: "You cannot onto land.",
	}
	var openSea string = "an open sea"
	o := terrain.Terrain{
		PeakMsg: openSea,
	}
	x := terrain.Terrain{
		PeakMsg:  openSea,
		DeathMsg: "You have sailed onto a mine, your ship exploded into a thousand pieces and you died.",
	}
	f := terrain.Terrain{
		PeakMsg:  openSea,
		DeathMsg: "TODO: You have sailed into a harbour and you have won.",
	}
	arr := [][]terrain.Terrain{
		{o, x, l, l, l, l, l, l, l, l, o, l, l, o, l, o, x, o, o, o, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, o, o},
		{x, o, o, o, o, o, l, l, l, l, o, l, l, x, o, o, o, x, l, l, l, l, l, x, l, l, l, l, l, l, l, l, l, l, o, l, o, o, x, o},
		{o, o, x, o, o, o, l, l, l, l, l, l, l, o, o, x, o, o, f, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, o, x, o, o, o, x},
		{o, o, o, o, o, o, l, l, x, l, l, l, o, o, x, o, o, o, l, l, o, l, l, l, l, l, l, l, l, l, l, l, l, l, o, o, o, o, o, o},
		{o, o, o, o, o, x, o, l, l, l, l, o, o, x, o, o, x, o, l, l, l, l, l, l, l, l, l, l, l, l, l, l, l, x, o, o, o, o, o, o},
		{o, x, o, o, o, o, o, o, l, l, o, x, o, o, x, x, o, x, l, l, l, l, x, l, l, l, x, l, l, l, l, l, o, o, o, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, l, l, l, o, o, o, o, x, l, l, l, l, l, l, l, o, o, l, o, l, l, o, l, x, o, x, o, o, o, o},
		{o, o, o, o, o, x, o, o, o, o, o, l, l, l, x, o, o, o, o, o, l, l, l, l, x, o, o, o, x, l, l, l, o, o, x, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, l, l, l, l, x, o, o, o, o, x, l, l, l, o, o, o, o, o, o, l, l, o, l, l, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, o, l, l, l, o, o, o, o, o, o, l, l, l, l, o, o, o, o, o, o, x, l, l, l, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, o, l, l, o, o, o, o, x, o, o, l, l, x, o, o, o, o, o, o, o, l, l, l, l, l, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, o, l, l, o, o, o, o, o, o, o, o, o, o, o, o, o, x, o, o, o, o, o, o, l, o, x, o, o, x},
		{o, o, o, o, o, o, o, o, o, o, o, o, l, x, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, l, l, o, o, o, o, x, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, x, o, o, o, o, o, o, o, o, o, o, o, o, x, o, o, o, o, o, x, o, o, o, o, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o},
		{o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, o, x, o, o, o, o, o, o, o, o, x, o, o, o, o, o, o, o, o, o, o, o, o},
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
		PlayerPosition: position.Position{X: 21, Y: 14},
		NothingTerrain: e,
		MoveMsg:        "You have sailed ",
		Intro:          intro,
		Help:           help,
	}
	g.Start()
}
