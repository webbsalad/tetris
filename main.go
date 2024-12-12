// Package main is the entry point for the Tetris game application.
// It initializes the game window, sets up the game loop using the Ebiten game library,
// and handles the execution of the game.
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
