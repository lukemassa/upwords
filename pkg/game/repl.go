package game

import (
	"fmt"
	"strconv"
	//	"strconv"
	//"math/rand"
	//"os"
	//"time"
	//"github.com/gdamore/tcell/v2"
)

// REPL a repl-based UI
type REPL struct{}

// InputScore enter the score
func (r REPL) InputScore(player *Player) int {
	for {
		fmt.Printf("Points for %s: ", *player)
		var entry string
		fmt.Scanln(&entry)
		val, err := strconv.Atoi(entry)
		if err == nil {
			return val
		}
	}
}

// Show the status of the game
func (r REPL) Show(game *Game) {
	fmt.Printf("%s: %d\n", *game.player1, game.score(game.player1))
	fmt.Printf("%s: %d\n", *game.player2, game.score(game.player2))
}
