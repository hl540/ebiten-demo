package examples

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hl540/ebiten-demo/resources/images"
)

type viewport struct {
	x16       int
	y16       int
	width     int
	height    int
	widthMax  int
	heightMax int
}

func (v *viewport) Position() (int, int) {
	return v.x16, v.y16
}

func (v *viewport) Move() {
	// 移动
	v.x16 += v.width / 32
	v.y16 += v.height / 32
	// 移动到末尾后从头开始
	v.x16 %= v.widthMax
	v.y16 %= v.heightMax
}

type InfiniteScroll struct {
	screenWidth  int
	screenHeight int
	bgImage      *ebiten.Image
	viewport     *viewport
}

func (i *InfiniteScroll) Run() error {
	i.screenWidth = 640
	i.screenHeight = 480

	// 加载图片
	img, _, err := image.Decode(bytes.NewBuffer(images.Tile_png))
	if err != nil {
		log.Fatal(err)
	}
	i.bgImage = ebiten.NewImageFromImage(img)

	// 初始化视口
	w, h := i.bgImage.Size()
	i.viewport = &viewport{
		width:     w,
		height:    h,
		widthMax:  w * 16,
		heightMax: h * 16,
	}

	ebiten.SetWindowSize(i.screenWidth, i.screenHeight)
	ebiten.SetWindowTitle("Infinite Scroll (Ebiten Demo)")
	return ebiten.RunGame(i)
}

func (i *InfiniteScroll) Update() error {
	i.viewport.Move()
	return nil
}

func (i *InfiniteScroll) Draw(screen *ebiten.Image) {
	x16, y16 := i.viewport.Position()
	offsetX, offsetY := float64(-x16)/16, float64(-y16)/16
	w, h := i.bgImage.Size()
	for j := 0; j < 3; j++ {
		for k := 0; k < 4; k++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*k), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(i.bgImage, op)
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualFPS()))
}

func (i *InfiniteScroll) Layout(outsideWidth, outsideHeight int) (int, int) {
	return i.screenWidth, i.screenHeight
}
