package examples

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type world struct {
	area   []bool
	width  int
	height int
}

func NewWorld(width, height, maxCell int) *world {
	w := &world{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}

	for i := 0; i < maxCell; i++ {
		x := rand.Intn(w.width)
		y := rand.Intn(w.height)
		w.area[y*w.width+x] = true
	}
	return w
}

func neighbourCount(a []bool, width, height, x, y int) int {
	c := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || width <= x2 || height <= y2 {
				continue
			}
			if a[y2*width+x2] {
				c++
			}
		}
	}
	return c
}

func (w *world) Update() {
	width := w.width
	height := w.height
	next := make([]bool, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pop := neighbourCount(w.area, width, height, x, y)
			switch {
			case pop < 2:
				next[y*width+x] = false
			case (pop == 2 || pop == 3) && w.area[y*width+x]:
				next[y*width+x] = true
			case pop > 3:
				next[y*width+x] = false
			case pop == 3:
				next[y*width+x] = true
			}
		}
	}
	w.area = next
}

type Life struct {
	screenWidth  int
	screenHeight int
	world        *world
	pixels       []byte
}

func (l *Life) Run() error {
	l.screenWidth = 640
	l.screenHeight = 480
	rand.Seed(time.Now().UnixNano())
	l.world = NewWorld(l.screenWidth, l.screenHeight, int(l.screenWidth*l.screenHeight)/10)

	ebiten.SetWindowSize(l.screenWidth, l.screenHeight)
	ebiten.SetWindowTitle("Game of Life (Ebiten Demo)")
	return ebiten.RunGame(l)
}

func (l *Life) Update() error {
	if len(l.pixels) == 0 {
		l.pixels = make([]byte, l.screenWidth*l.screenHeight*4)
	}
	l.world.Update()
	for i, v := range l.world.area {
		if v {
			l.pixels[4*i] = 0xff
			l.pixels[4*i+1] = 0xff
			l.pixels[4*i+2] = 0xff
			l.pixels[4*i+3] = 0xff
		} else {
			l.pixels[4*i] = 0
			l.pixels[4*i+1] = 0
			l.pixels[4*i+2] = 0
			l.pixels[4*i+3] = 0
		}
	}
	return nil
}

func (l *Life) Draw(screen *ebiten.Image) {
	screen.WritePixels(l.pixels)
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (l *Life) Layout(outsideWidth, outsideHeight int) (int, int) {
	return l.screenWidth, l.screenHeight
}
