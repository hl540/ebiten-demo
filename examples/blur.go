package examples

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hl540/ebiten-demo/resources/images"
)

var gophersImage *ebiten.Image

type Blur struct {
	screenWidth  int
	screenHeight int
}

func (b *Blur) Run() error {
	b.screenWidth = 640
	b.screenHeight = 480
	img, _, err := image.Decode(bytes.NewReader(images.FiveYears_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)
	ebiten.SetWindowSize(b.screenWidth, b.screenHeight)
	ebiten.SetWindowTitle("Blur (Ebiten Demo)")
	return ebiten.RunGame(b)
}

func (b *Blur) Update() error {
	return nil
}

func (b *Blur) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(gophersImage, op)

	layers := 0
	for j := -3; j <= 3; j++ {
		for i := -3; i <= 3; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i), 244+float64(j))
			layers++
			op.ColorM.Scale(1, 1, 1, 1.0/float64(layers))
			screen.DrawImage(gophersImage, op)
		}
	}
}

func (b *Blur) Layout(outsideWidth, outsideHeight int) (int, int) {
	return b.screenWidth, b.screenHeight
}
