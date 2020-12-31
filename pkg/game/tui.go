package game

import (
	"fmt"
	"os"
	"time"

	//	"strconv"
	//"math/rand"
	//"os"
	//"time"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// TUI a text based ui
type TUI struct {
	screen *tcell.Screen
}

// NewTUI a new initialized
func NewTUI() TUI {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Show()
	t := TUI{
		screen: &s,
	}
	go t.run()
	return t

}

func (t TUI) run() {
	tick := time.Tick(1 * time.Second)
	a := 0
	for {
		select {
		// Got a tick, we should check on doSomething()
		case <-tick:
			emitStr(*t.screen, 0, 0, tcell.StyleDefault, fmt.Sprintf("%d", a))
			(*t.screen).Show()
			a++
		}
	}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

// InputScore enter the score
func (t TUI) InputScore(player *Player, game *Game) Move {

	for {
		ev := (*t.screen).PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {

				(*t.screen).Fini()
				os.Exit(0)
			}
			//t.screen.SetCell(5, 1, tcell.StyleDefault, ev.Rune())
			//t.screen.Show()
			r := ev.Rune()
			if r >= '0' && r <= '9' {
				return Move{
					value: int(r - '0'),
				}
			}

		}
	}
}

// Show the status of the game
func (t TUI) Show(game *Game) {

	/*
		for _, v := range game.scores[game.player1] {
			lastScore = v
		}
		//r := []rune(strconv.Itoa(lastScore))
		//t.screen.Fill(r[0], tcell.StyleDefault)
	*/
	emitStr((*t.screen), 5, 5, tcell.StyleDefault, fmt.Sprintf("%s: %d", *game.player1, game.score(game.player1)))
	emitStr((*t.screen), 5, 6, tcell.StyleDefault, fmt.Sprintf("%s: %d", *game.player2, game.score(game.player2)))
	//emitStr(t.screen, 6, 5, tcell.StyleDefault, strconv.Itoa(lastScore))
	(*t.screen).Show()
	//	t.screen.Fini()
}
