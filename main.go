package main

import (
	"fmt"
	"log"

	"github.com/webbsalad/tetris/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.New()

	fmt.Println(g)

	ebiten.SetWindowSize(320, 640)
	ebiten.SetWindowTitle("Tetris")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
