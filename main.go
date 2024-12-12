package main

import (
	"log"

	"github.com/webbsalad/tetris/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.New()

	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Tetris")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
