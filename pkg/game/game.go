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

// Player one of the players of the game
type Player string

// Game a full game
type Game struct {
	player1     *Player
	player2     *Player
	player1turn bool
	scores      map[*Player][]int
}

// UI an interface
type UI interface {
	Show()
	InputScore(*Player) int
}

func (g *Game) inputScore(player *Player) int {
	for {
		fmt.Printf("Points for %s: ", *player)
		var entry string
		fmt.Scanln(&entry)
		if entry == "show" {
			g.Show()
			continue
		}
		val, err := strconv.Atoi(entry)
		if err == nil {
			return val
		}
	}
}
func (g *Game) turn() {

	score := g.inputScore(g.whoseTurn())
	g.scores[g.whoseTurn()] = append(g.scores[g.whoseTurn()], score)
}

// New a new game
func New(player1, player2 *Player) Game {
	g := Game{
		player1: player1,
		player2: player2,
		scores: map[*Player][]int{
			player1: make([]int, 0),
			player2: make([]int, 0),
		},
	}
	return g

}

func (g *Game) score(player *Player) int {
	total := 0
	for i := 0; i < len(g.scores[player]); i++ {
		total += g.scores[player][i]
	}
	return total
}

func (g *Game) whoseTurn() *Player {
	if g.player1turn {
		return g.player1
	}
	return g.player2
}

// Show the current situation
func (g *Game) Show() {
	fmt.Printf("%s: %d\n", *g.player1, g.score(g.player1))
	fmt.Printf("%s: %d\n", *g.player2, g.score(g.player2))
}

// Play the game
func (g *Game) Play() {
	for {
		g.turn()
		g.Show()
		g.player1turn = !g.player1turn
	}
}
