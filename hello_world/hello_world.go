package hello_world

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type HelloWorld struct {
	num int
}

func (h *HelloWorld) Run() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill")
	return ebiten.RunGame(h)
}

func (h *HelloWorld) Update() error {
	h.num += 1
	return nil
}

func (h *HelloWorld) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World![%d]", h.num))
}

func (h *HelloWorld) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
