package tour

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fill struct {
}

func (f *Fill) Run() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill")
	return ebiten.RunGame(f)
}

func (f *Fill) Update() error {
	return nil
}

func (f *Fill) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
}

func (f *Fill) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
