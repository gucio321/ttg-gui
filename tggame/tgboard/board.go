package tgboard

import (
	"bytes"
	"image"
	_ "image/png" // it is used by image.Decode (without this the Decode crashes)
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gucio321/tic-tac-go/ttgcore/ttgboard"
	"github.com/gucio321/tic-tac-go/ttgcore/ttgletter"

	"github.com/gucio321/ttg-gui/tgassets"
)

const (
	imgW, imgH = 50, 50
	separator  = 20
)

// Board represents a game board
type Board struct {
	*ttgboard.Board
	imgX, imgO *ebiten.Image
}

// New creates a new board
func New(w, h, c int) *Board {
	result := &Board{}
	result.Board = ttgboard.NewBoard(w, h, c)

	img, _, err := image.Decode(bytes.NewReader(tgassets.ImageX))
	if err != nil {
		log.Fatal(err)
	}

	result.imgX = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(tgassets.ImageO))
	if err != nil {
		log.Fatal(err)
	}

	result.imgO = ebiten.NewImageFromImage(img)

	return result
}

// Draw draws the board
func (b *Board) Draw(screen *ebiten.Image) {
	for y := 0; y < b.Width(); y++ {
		for x := 0; x < b.Height(); x++ {
			var img *ebiten.Image

			switch b.GetIndexState(b.Width()*y + x) {
			case ttgletter.LetterX:
				img = b.imgX
			case ttgletter.LetterO:
				img = b.imgO
			default:
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*imgW+separator), float64(y*imgH+separator))
			screen.DrawImage(img, op)
		}
	}
}
