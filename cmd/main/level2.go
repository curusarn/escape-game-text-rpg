package main

import (
	"fmt"
	"time"

	"github.com/curusarn/escape-game-text-rpg/game"
	"github.com/curusarn/escape-game-text-rpg/position"
	"github.com/curusarn/escape-game-text-rpg/terrain"
)

const treasureStr = "xxx"
const congats = `
  ____                            _         _       _   _                 
 / ___|___  _ __   __ _ _ __ __ _| |_ _   _| | __ _| |_(_) ___  _ __  ___ 
| |   / _ \| '_ \ / _' | '__/ _' | __| | | | |/ _' | __| |/ _ \| '_ \/ __|
| |__| (_) | | | | (_| | | | (_| | |_| |_| | | (_| | |_| | (_) | | | \__ \
 \____\___/|_| |_|\__, |_|  \__,_|\__|\__,_|_|\__,_|\__|_|\___/|_| |_|___/
                  |___/                                                   
`

func EvalDigAction(g *game.Game) (bool, string) {
	fmt.Printf("You started digging at current location ")
	for i := 0; i < 5; i++ {
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\n")
	t := g.GetTerrainOnPosition(g.PlayerPosition)
	if t.ActionStr == treasureStr {
		g.GameState = 2
		return false, "You have found something!" + congats
	}
	return true, "You have saerched the whole place and found nothing."
}

func level2() int {
	peekPlains := "free space (plains)"
	xx := terrain.Terrain{
		PeakMsg:   peekPlains,
		ActionStr: treasureStr,
	}

	cantMoveSea := "You cannot go to the sea."
	ww := terrain.Terrain{
		PeakMsg:     "the sea",
		CantMoveMsg: cantMoveSea,
	}
	ff := terrain.Terrain{
		PeakMsg:     "a fortress surrounded with walls",
		CantMoveMsg: "You cannot go to the fortress.",
	}
	cantMoveMountains := "You cannot go to the mountains."
	mm := terrain.Terrain{
		PeakMsg:     "mountains",
		CantMoveMsg: cantMoveMountains,
	}
	mo := terrain.Terrain{
		PeakMsg: "rocky elevated terrain",
	}

	oo := terrain.Terrain{
		PeakMsg: peekPlains,
	}
	mw := terrain.Terrain{
		// TODO: better message
		PeakMsg:     "the sea (with cliffs above it)",
		CantMoveMsg: cantMoveSea,
	}
	rm := terrain.Terrain{
		PeakMsg:     "a river leading to the mountains",
		CantMoveMsg: cantMoveMountains,
	}
	rr := terrain.Terrain{
		PeakMsg: "a river (looks shallow)",
	}
	rx := terrain.Terrain{
		PeakMsg: "a river passing through a weird sun-shaped thing (maybe a crater)",
	}
	ho := terrain.Terrain{
		PeakMsg: "a house",
	}
	hr := terrain.Terrain{
		PeakMsg: "a house near a river",
	}
	hh := terrain.Terrain{
		PeakMsg: "a group of houses (part of a small town)",
	}
	hc := terrain.Terrain{
		PeakMsg: "a church",
	}
	hf := terrain.Terrain{
		PeakMsg: "a tiny fishing camp near a river",
	}

	t1 := terrain.Terrain{
		PeakMsg: "a single tree",
	}
	t2 := terrain.Terrain{
		PeakMsg: "a group of two trees",
	}
	t3 := terrain.Terrain{
		PeakMsg: "a group of three trees",
	}
	t5 := terrain.Terrain{
		PeakMsg: "a circle of five trees",
	}

	arr := [][]terrain.Terrain{
		{ww, ww, ww, ww, ww, ww, ww, ww, mw, mw, mw, mw, mw, mw, mw, ww, ww, ww, ww, ww, ww, ww, ww},
		{ww, ww, ww, ww, ww, ww, mw, mw, oo, oo, oo, oo, oo, oo, oo, mw, mw, ww, ww, ww, ww, ww, ww},
		{ww, ww, ww, ww, ww, mw, mw, oo, oo, oo, oo, oo, oo, oo, oo, oo, mw, mw, mw, oo, ww, ww, ww},
		{ww, ww, ww, mw, mw, oo, oo, oo, oo, mm, mm, mm, mm, oo, xx, oo, oo, oo, oo, oo, oo, ww, ww},
		{ww, ww, ww, oo, oo, oo, oo, oo, mo, mm, mm, mm, mm, mo, oo, oo, oo, mm, mm, oo, oo, oo, ww},
		{ww, ww, oo, hc, oo, oo, oo, oo, oo, mo, mo, mo, mo, t3, oo, mo, mm, mm, mm, mm, mo, oo, ww},
		{ww, ww, oo, oo, oo, oo, oo, mo, mm, mm, mm, mm, oo, oo, oo, oo, oo, oo, oo, oo, oo, oo, ww},
		{oo, oo, oo, oo, oo, oo, oo, oo, mm, oo, oo, oo, oo, mo, mm, mm, mm, mm, mm, oo, oo, oo, ww},
		{oo, oo, oo, oo, oo, oo, t2, t1, oo, oo, mm, mm, mm, mo, mo, mm, mm, mm, mm, mm, oo, oo, ww},
		{oo, oo, t2, oo, t2, oo, oo, oo, mm, mm, mm, mm, mm, mm, mo, mo, mo, mo, mo, mo, mo, oo, ww},
		{oo, oo, t3, oo, oo, oo, t1, t2, mm, mm, rm, mm, mm, mm, mm, mm, mo, t2, t2, oo, oo, rr, oo},
		{ww, oo, oo, oo, t1, t1, oo, oo, oo, oo, rr, oo, t1, t1, t1, t2, t2, t1, t1, oo, t5, rr, rr},
		{ww, oo, oo, oo, oo, oo, oo, oo, oo, oo, rr, oo, oo, oo, oo, oo, oo, oo, oo, oo, oo, oo, ww},
		{ww, oo, rx, rr, oo, oo, oo, oo, oo, oo, rr, oo, oo, oo, oo, oo, oo, oo, oo, oo, oo, oo, ww},
		{ww, ww, rx, oo, oo, oo, oo, oo, oo, oo, hr, oo, oo, hh, hh, hh, oo, mw, mw, oo, oo, mw, ww},
		{ww, ww, ww, oo, oo, oo, oo, rr, oo, ho, ww, ww, ww, hh, hh, hh, oo, mw, ww, oo, mw, ww, ww},
		{ww, ww, ww, ww, oo, oo, oo, oo, hf, ww, ww, ww, ww, ww, oo, mw, mw, ww, ww, ww, ww, ww, ww},
		{ww, ww, ww, ww, ww, ww, oo, ff, ww, ww, ww, ww, ww, ww, ww, ww, ww, ww, oo, ww, ww, ww, ww},
	}
	gameMap := game.BuildGameMapFromSlice(arr)
	intro := `
  ____ _            _                   _            _     ____  ____   ____ 
 / ___| | ___  _ __(_) ___  _   _ ___  | |_ _____  _| |_  |  _ \|  _ \ / ___|
| |  _| |/ _ \| '__| |/ _ \| | | / __| | __/ _ \ \/ / __| | |_) | |_) | |  _ 
| |_| | | (_) | |  | | (_) | |_| \__ \ | ||  __/>  <| |_  |  _ <|  __/| |_| |
 \____|_|\___/|_|  |_|\___/ \__,_|___/  \__\___/_/\_\\__| |_| \_\_|    \____|

 _     _______     _______ _     ____  
| |   | ____\ \   / / ____| |   |___ \ 
| |   |  _|  \ \ / /|  _| | |     __) |
| |___| |___  \ V / | |___| |___ / __/ 
|_____|_____|  \_/  |_____|_____|_____|

You are a sailor standing on a beach. 
Where do you want to go next?
Type your next command. (Type 'help' for list of available commands.)
-----------------------------------------
### Type 'dig' to search for treasure ###
-----------------------------------------
	`
	help := `
Search for treasure with 'dig'
Walk around with 'north'/'east'/'south'/'west'
Short version of directions also works: n/e/s/w
Show this help with 'help'.
Use 'look' to look around.
Exit with 'exit'.
	`
	g := game.Game{
		GameMap:        gameMap,
		PlayerPosition: position.Position{X: 21, Y: 13},
		NothingTerrain: ww,
		MoveMsg:        "You went ",
		Intro:          intro,
		Help:           help,
		ActionStr:      "dig",
		ActionFunc:     EvalDigAction,
	}

	return g.Start()
}
