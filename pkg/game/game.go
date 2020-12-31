package game

//	"strconv"
//"math/rand"
//"os"
//"time"
//"github.com/gdamore/tcell/v2"

// ShowMove simple move just to show the current status
var ShowMove = Move{
	value: -1,
}

// Player one of the players of the game
type Player string

// Move A particular move. Typically a number, but might be different instruction
type Move struct {
	value int
}

// Game a full game
type Game struct {
	player1     *Player
	player2     *Player
	player1turn bool
	scores      map[*Player][]int
	ui          UI
}

// UI an interface
type UI interface {
	Show(*Game)
	InputScore(*Player, *Game) Move
}

func (g *Game) turn() {

	for {
		move := g.ui.InputScore(g.whoseTurn(), g)
		if move == ShowMove {
			g.ui.Show(g)
			continue
		}
		g.scores[g.whoseTurn()] = append(g.scores[g.whoseTurn()], move.value)
		return
	}
}

// New a new game
func New(player1, player2 *Player, ui UI) Game {
	g := Game{
		player1: player1,
		player2: player2,
		scores: map[*Player][]int{
			player1: make([]int, 0),
			player2: make([]int, 0),
		},
		ui: ui,
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

// Play the game
func (g *Game) Play() {
	g.ui.Show(g)
	for {
		g.turn()
		g.ui.Show(g)
		g.player1turn = !g.player1turn
	}
}
