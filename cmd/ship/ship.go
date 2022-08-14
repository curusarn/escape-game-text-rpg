package main

import (
	"fmt"

	"github.com/curusarn/escape-game-text-rpg/game"
	"github.com/curusarn/escape-game-text-rpg/position"
	"github.com/curusarn/escape-game-text-rpg/terrain"
)

func run() int {
	e := terrain.Terrain{
		PeakMsg:     "the edge of the world",
		CantMoveMsg: "You cannot sail over the edge of the world.",
	}
	l := terrain.Terrain{
		PeakMsg:     "land",
		CantMoveMsg: "You cannot sail onto land.",
	}
	var openSea string = "an open sea"
	o := terrain.Terrain{
		PeakMsg: openSea,
	}
	youDied := `
__   _____  _   _   ____ ___ _____ ____  
\ \ / / _ \| | | | |  _ \_ _| ____|  _ \ 
 \ V / | | | | | | | | | | ||  _| | | | |
  | || |_| | |_| | | |_| | || |___| |_| |
  |_| \___/ \___/  |____/___|_____|____/ 

	`
	x := terrain.Terrain{
		PeakMsg:   openSea,
		DeathMsg:  "You have sailed onto a mine, your ship exploded into thousand pieces and ..." + youDied,
		GameState: 1,
	}
	victory := `
__     _____ ____ _____ ___  ______   __
\ \   / /_ _/ ___|_   _/ _ \|  _ \ \ / /
 \ \ / / | | |     | || | | | |_) \ V / 
  \ V /  | | |___  | || |_| |  _ < | |  
   \_/  |___\____| |_| \___/|_| \_\|_|  

	`
	f := terrain.Terrain{
		PeakMsg:   openSea,
		DeathMsg:  "You have found a harbour." + victory,
		GameState: 2,
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
__        __   _                            _        
\ \      / /__| | ___ ___  _ __ ___   ___  | |_ ___  
 \ \ /\ / / _ \ |/ __/ _ \| '_ ' _ \ / _ \ | __/ _ \ 
  \ V  V /  __/ | (_| (_) | | | | | |  __/ | || (_) |
   \_/\_/ \___|_|\___\___/|_| |_| |_|\___|  \__\___/ 
                                                     
  ____ _            _                   _            _     ____  ____   ____ 
 / ___| | ___  _ __(_) ___  _   _ ___  | |_ _____  _| |_  |  _ \|  _ \ / ___|
| |  _| |/ _ \| '__| |/ _ \| | | / __| | __/ _ \ \/ / __| | |_) | |_) | |  _ 
| |_| | | (_) | |  | | (_) | |_| \__ \ | ||  __/>  <| |_  |  _ <|  __/| |_| |
 \____|_|\___/|_|  |_|\___/ \__,_|___/  \__\___/_/\_\\__| |_| \_\_|    \____|
	
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

	return g.Start()
}

func main() {
	for {
		state := run()
		if state == 1 {
			fmt.Printf("Restarting the game ...")
			continue
		}
		if state == 2 {
			fmt.Printf("TODO: Victory  (level 2) ...")
			// Mark level1 as completed
			break
		}
		break
	}
	return
}
