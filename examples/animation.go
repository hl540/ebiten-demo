package examples

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hl540/ebiten-demo/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0  // 帧动画图片x开始位置
	frameOY     = 32 // 帧动画图片y开始位置
	frameWidth  = 32 // 单个帧图片宽度
	frameHeight = 32 // 单个帧图片高度
	frameNum    = 8  // 帧动画图片数量
)

var runnerImage *ebiten.Image

type Animation struct {
	count int
}

func (a *Animation) Run() error {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	return ebiten.RunGame(a)
}

func (a *Animation) Update() error {
	a.count++
	return nil
}

func (a *Animation) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// 居中
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	// 获取下一帧图片index
	i := (a.count / 5) % frameNum
	// 计算下一帧图片偏移量
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (a *Animation) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
