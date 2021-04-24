package tggame

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gucio321/tic-tac-go/ttggame/ttgletter"
	"github.com/gucio321/ttg-gui/tggame/tgboard"
)

type Game struct {
	board *tgboard.Board
}

func Create() *Game {
	result := &Game{}
	result.board = tgboard.New(3, 3, 3)
	result.board.SetIndexState(0, ttgletter.LetterX)
	result.board.SetIndexState(1, ttgletter.LetterO)
	result.board.SetIndexState(2, ttgletter.LetterX)

	return result
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return 640, 480
}

func (g *Game) Update() error {
	// noop
	return nil
}
