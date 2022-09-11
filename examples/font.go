package examples

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hl540/ebiten-demo/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	sampleText = `The quick brown fox jumps over the lazy dog.`
	texts      = `
白日依山尽
黄河入海流
欲穷千里目
更上一层楼
`
)

type Font struct {
	screenWidth     int
	screenHeight    int
	counter         int
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	kanjiTextColor  color.RGBA
}

func (f *Font) Run() error {
	// 加载字体
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	const dpi = 72
	f.mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	f.mplusBigFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	f.screenWidth = 640
	f.screenHeight = 480
	ebiten.SetWindowSize(f.screenWidth, f.screenHeight)
	ebiten.SetWindowTitle("Font (Ebiten Demo)")
	return ebiten.RunGame(f)
}

func (f *Font) Update() error {
	if f.counter%ebiten.TPS() == 0 {
		f.kanjiTextColor.R = uint8(rand.Intn(0x7f))
		f.kanjiTextColor.G = uint8(rand.Intn(0x7f))
		f.kanjiTextColor.B = uint8(rand.Intn(0x7f))
		f.kanjiTextColor.A = 0xff
	}
	f.counter++
	return nil
}

func (f *Font) Draw(screen *ebiten.Image) {
	const x = 20
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	text.Draw(screen, msg, f.mplusNormalFont, x, 40, color.White)
	text.Draw(screen, sampleText, f.mplusNormalFont, x, 80, color.White)
	for i, line := range strings.Split(texts, "\n") {
		text.Draw(screen, line, f.mplusBigFont, x, 120+50*i, f.kanjiTextColor)
	}

}

func (f *Font) Layout(outsideWidth, outsideHeight int) (int, int) {
	return f.screenWidth, f.screenHeight
}
