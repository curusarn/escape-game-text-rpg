package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/curusarn/escape-game-text-rpg/game"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("ERROR: Failed to get user home dir: %v\n", err)
		return
	}
	screenLogPath := path.Join(home, "text-rpg-screen.log")
	statusLogPath := path.Join(home, "text-rpg-status.log")
	// If the file doesn't exist, create it, or append to the file
	screenLog, err := os.OpenFile(screenLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("ERROR: Could not open/create log file: %v\n", err)
		return
	}
	defer screenLog.Close()
	// If the file doesn't exist, create it, or append to the file
	statusLog, err := os.OpenFile(statusLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("ERROR: Could not open/create log file: %v\n", err)
		return
	}
	defer statusLog.Close()
	lvl1File := path.Join(home, "text-rpg-level1-done")
	lvl2File := path.Join(home, "text-rpg-level2-done")
	for {
		if _, err := os.Stat(lvl1File); errors.Is(err, os.ErrNotExist) {
			state := level1(screenLog, statusLog)
			if state == 0 {
				// exited
				break
			}
			if state == 1 {
				time.Sleep(3 * time.Second)
				game.LogPrintf(screenLog, "Restarting the game ...\n")
				continue
			}
			if state == 2 {
				// Mark level1 as completed
				game.LogPrintf(screenLog, "Saving progress ...\n")
				_, err := os.Create(lvl1File)
				if err != nil {
					game.LogPrintf(screenLog, "ERROR: Failed to save progress\n")
					break
				}
				game.LogPrintf(screenLog, "Progress saved successfully\n")
				game.LogPrintf(screenLog, "You won't have to complete the level again if you exit the session.\n")
				time.Sleep(3 * time.Second)
				continue
			}
			break
		}
		if _, err := os.Stat(lvl2File); errors.Is(err, os.ErrNotExist) {
			state := level2(screenLog, statusLog)
			if state == 0 {
				// exited
				break
			}
			if state == 1 {
				game.LogPrintf(screenLog, "ERROR: Illegal state: 1 (level2)\n")
				break
			}
			if state == 2 {
				// Mark level2 as completed
				game.LogPrintf(screenLog, "Saving progress ...\n")
				_, err := os.Create(lvl2File)
				if err != nil {
					game.LogPrintf(screenLog, "ERROR: Failed to save progress\n")
					break
				}
				game.LogPrintf(screenLog, "Progress saved successfully\n")
				// game.LogPrintf(logFile, "You won't have to complete the level again if you exit the session.\n")
				time.Sleep(3 * time.Second)
				continue
			}
			break

		}
		code := `
_____   _  _     _____    ___  
|___  | | || |   |___ /   / _ \ 
   / /  | || |_    |_ \  | (_) |
  / /   |__   _|  ___) |  \__, |
 /_/       |_|   |____/     /_/ 
		`
		game.LogPrintf(screenLog, "\n\nThe final treasure code is %s\n\n", code)
		break
	}
	return
}
