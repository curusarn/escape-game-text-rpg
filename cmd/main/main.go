package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("ERROR: Failed to get user home dir\n")
		return
	}
	lvl1File := path.Join(home, "text-rpg-level1-done")
	lvl2File := path.Join(home, "text-rpg-level2-done")
	for {
		if _, err := os.Stat(lvl1File); errors.Is(err, os.ErrNotExist) {
			state := level1()
			if state == 0 {
				// exited
				break
			}
			if state == 1 {
				time.Sleep(3 * time.Second)
				fmt.Printf("Restarting the game ...\n")
				continue
			}
			if state == 2 {
				// Mark level1 as completed
				fmt.Printf("Saving progress ...\n")
				_, err := os.Create(lvl1File)
				if err != nil {
					fmt.Printf("ERROR: Failed to save progress\n")
					break
				}
				fmt.Printf("Progress saved successfully\n")
				fmt.Printf("You won't have to complete the level again if you exit the session.\n")
				time.Sleep(3 * time.Second)
				continue
			}
			break
		}
		if _, err := os.Stat(lvl2File); errors.Is(err, os.ErrNotExist) {
			state := level2()
			if state == 0 {
				// exited
				break
			}
			if state == 1 {
				fmt.Printf("ERROR: Illegal state: 1 (level2)\n")
				break
			}
			if state == 2 {
				// Mark level2 as completed
				fmt.Printf("Saving progress ...\n")
				_, err := os.Create(lvl2File)
				if err != nil {
					fmt.Printf("ERROR: Failed to save progress\n")
					break
				}
				fmt.Printf("Progress saved successfully\n")
				// fmt.Printf("You won't have to complete the level again if you exit the session.\n")
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
		fmt.Printf("\n\nThe final treasure code is %s\n\n", code)
		break
	}
	return
}
