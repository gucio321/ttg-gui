package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gucio321/ttg-gui/tggame"
)

const (
	screenW, screenH = 640, 480
)

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Tic-Tac-Go")
	if err := ebiten.RunGame(tggame.Create()); err != nil {
		log.Fatal(err)
	}
}
