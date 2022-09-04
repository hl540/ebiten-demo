package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hl540/ebiten-demo/examples"
)

type Game interface {
	Run() error
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World!")

	// game := &hello_world.HelloWorld{}
	// game := tour.Fill{}
	// game := tour.Image{}
	// game := tour.Geom{}
	game := examples.Animation{}
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
