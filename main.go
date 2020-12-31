package main

import "github.com/lukemassa/upwords/pkg/game"

func main() {
	player1 := game.Player("Luke")
	player2 := game.Player("Nick")
	ui := game.NewTUI()
	//ui := game.REPL{}

	g := game.New(&player1, &player2, ui)

	g.Play()
}
