package tour

import "github.com/hajimehoshi/ebiten/v2"

type Geom struct{}

func (g *Geom) Run() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Geometry Matrix")
	return ebiten.RunGame(g)
}

func (g *Geom) Update() error {
	return nil
}

func (g *Geom) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(1.5, 1)
	screen.DrawImage(img, op)
}

func (g *Geom) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
