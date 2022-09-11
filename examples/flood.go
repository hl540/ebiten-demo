package examples

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hl540/ebiten-demo/resources/images"
)

type Floog struct {
	screenWidth  int
	screenHeight int
	ebitenImage  *ebiten.Image
}

func (f *Floog) Run() error {
	f.screenWidth = 320
	f.screenHeight = 240
	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	f.ebitenImage = ebiten.NewImageFromImage(img)
	ebiten.SetWindowSize(f.screenWidth*2, f.screenHeight*4)
	ebiten.SetWindowTitle("Flood fill with solid colors (Ebiten Demo)")
	return ebiten.RunGame(f)
}

func (f *Floog) Update() error {
	return nil
}

func (f *Floog) Draw(screen *ebiten.Image) {
	// 填充背景
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	// 原图
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(f.ebitenImage, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 50)
	// 重置RGB
	op.ColorM.Scale(0, 0, 0, 1)
	// flood fill
	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	op.ColorM.Translate(float64(c.R), float64(c.G), float64(c.B), 0)
	screen.DrawImage(f.ebitenImage, op)
}

func (f *Floog) Layout(outsideWidth, outsideHeight int) (int, int) {
	return f.screenWidth, f.screenHeight
}
