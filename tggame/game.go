package tggame

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
	"github.com/gucio321/ttg-gui/tggame/tgboard"
)

// Game represents a tic-tac-go (GIU) game
type Game struct {
	board            *tgboard.Board
	screenW, screenH int
}

// Create creates a new game
func Create(w, h int) *Game {
	result := &Game{
		screenW: w,
		screenH: h,
	}

	// nolint:gomnd // test data
	result.board = tgboard.New(3, 3, 3)
	result.board.SetIndexState(0, ttgletter.LetterX)
	result.board.SetIndexState(1, ttgletter.LetterO)
	// nolint:gomnd // test data
	result.board.SetIndexState(2, ttgletter.LetterX)

	return result
}

// Draw renders stuff
func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

// Layout returns scene size
func (g *Game) Layout(outsideW, outsideH int) (w, h int) {
	return g.screenW, g.screenH
}

// Update updates scene
func (g *Game) Update() error {
	// noop
	return nil
}
