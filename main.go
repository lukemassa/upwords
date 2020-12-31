package main

import "github.com/lukemassa/upwords/pkg/game"

func main() {
	player1 := game.Player("Luke")
	player2 := game.Player("Nick")

	g := game.New(&player1, &player2)

	g.Play()
}
