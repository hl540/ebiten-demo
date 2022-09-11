package examples

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hl540/ebiten-demo/resources/images"
)

type Hsv struct {
	screenWidth   int
	screenHeight  int
	hue128        int
	saturation128 int
	value128      int
	inverted      bool
	gophersImage  *ebiten.Image
}

func (h *Hsv) Run() error {
	h.screenWidth = 640
	h.screenHeight = 480

	h.saturation128 = 128
	h.value128 = 128

	// 加载图片
	img, _, err := image.Decode(bytes.NewBuffer(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	h.gophersImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(h.screenWidth, h.screenHeight)
	ebiten.SetWindowTitle("Font (Ebiten Demo)")
	return ebiten.RunGame(h)
}

func clamp(v, min, max int) int {
	if min > max {
		panic("min must <= max")
	}
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func (h *Hsv) Update() error {
	// 监听键盘
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		h.hue128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		h.hue128++
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		h.saturation128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		h.saturation128++
	}
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		h.value128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		h.value128++
	}
	h.hue128 = clamp(h.hue128, -256, 256)
	h.saturation128 = clamp(h.saturation128, 0, 256)
	h.value128 = clamp(h.value128, 0, 256)

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		h.inverted = !h.inverted
	}
	return nil
}

func (h *Hsv) Draw(screen *ebiten.Image) {
	// 居中
	w, heg := h.gophersImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(heg)/2)
	// 缩放
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(h.screenWidth)/2, float64(h.screenHeight)/2)

	// 改变hsv
	hue := float64(h.hue128) * 2 * math.Pi / 128
	saturation := float64(h.saturation128) / 128
	value := float64(h.value128) / 128
	op.ColorM.ChangeHSV(hue, saturation, value)

	// 颜色反转
	if h.inverted {
		op.ColorM.Scale(-1, -1, -1, 1)
		op.ColorM.Translate(1, 1, 1, 0)
	}
	screen.DrawImage(h.gophersImage, op)

	// 显示提示语
	msg := fmt.Sprintf("Hue:        %0.2f [Q][W]\n", hue)
	msg += fmt.Sprintf("Saturation: %0.2f [A][S]\n", saturation)
	msg += fmt.Sprintf("Value:      %0.2f [Z][X]\n", value)
	msg += fmt.Sprintf("Inverted:   %v [I]\n", h.inverted)
	ebitenutil.DebugPrint(screen, msg)
}

func (h *Hsv) Layout(outsideWidth, outsideHeight int) (int, int) {
	return h.screenWidth, h.screenHeight
}
