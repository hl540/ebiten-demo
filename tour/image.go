package tour

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Image struct{}

func (i *Image) Run() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	return ebiten.RunGame(i)
}

func (i *Image) Update() error {
	return nil
}

func (i *Image) Draw(screem *ebiten.Image) {
	screem.DrawImage(img, nil)
}

func (i *Image) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
