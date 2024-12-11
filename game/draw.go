package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{54, 54, 54, 54})

	for i := 0; i <= g.boardWidth(); i++ {
		ebitenutil.DrawLine(screen, float64(i*tileSize), 0, float64(i*tileSize), float64(g.screenHeight), color.RGBA{27, 27, 27, 65})
	}
	for i := 0; i <= g.boardHeight(); i++ {
		ebitenutil.DrawLine(screen, 0, float64(i*tileSize), float64(g.screenWidth), float64(i*tileSize), color.RGBA{27, 27, 27, 65})
	}

	for i, row := range g.board {
		for j, cell := range row {
			if cell == 1 {
				ebitenutil.DrawRect(screen, float64(j*tileSize), float64(i*tileSize), float64(tileSize), float64(tileSize), g.currentColor)
			}
		}
	}

	for i, row := range g.currentTetromino {
		for j, cell := range row {
			if cell == 1 {
				ebitenutil.DrawRect(screen, float64((g.tetrominoX+j)*tileSize), float64((g.tetrominoY+i)*tileSize), float64(tileSize), float64(tileSize), g.currentColor)
			}
		}
	}

	if g.isGameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press R to Restart\nScore: "+strconv.Itoa(g.score))
	} else {
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score))
	}

	if g.isChoosingDifficulty {
		ebitenutil.DebugPrint(screen, "Choose Difficulty:\n1. Easy\n2. Medium\n3. Hard")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g.screenWidth <= 0 || g.screenHeight <= 0 {
		return 320, 640
	}
	return g.screenWidth, g.screenHeight
}
