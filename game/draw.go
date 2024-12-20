// Package game provides functionality for managing the game process,
// including sound handling, game state management, and player interactions.
package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw renders the game scene on the provided screen image.
// It fills the screen with a background color and then draws the game board,
// Tetromino, and other elements such as the score and game-over message.
//
// The game grid is drawn with lines representing the board's columns and rows.
// The active Tetromino is drawn on top of the grid, and the current game state is displayed in the corner.
// If the game is over, the score and a restart prompt are shown, and if the user is choosing difficulty, options are displayed.
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the background with a gray color
	screen.Fill(color.RGBA{54, 54, 54, 54})

	// Draw the vertical grid lines
	for i := 0; i <= g.boardWidth(); i++ {
		ebitenutil.DrawLine(screen, float64(i*tileSize), 0, float64(i*tileSize), float64(g.screenHeight), color.RGBA{27, 27, 27, 65})
	}

	// Draw the horizontal grid lines
	for i := 0; i <= g.boardHeight(); i++ {
		ebitenutil.DrawLine(screen, 0, float64(i*tileSize), float64(g.screenWidth), float64(i*tileSize), color.RGBA{27, 27, 27, 65})
	}

	// Draw the static blocks on the board (cells that are part of the game)
	for i, row := range g.board {
		for j, cell := range row {
			if cell == 1 {
				// Draw a block for each filled cell in the board
				ebitenutil.DrawRect(screen, float64(j*tileSize), float64(i*tileSize), float64(tileSize), float64(tileSize), g.currentColor)
			}
		}
	}

	// Draw the active Tetromino
	for i, row := range g.currentTetromino {
		for j, cell := range row {
			if cell == 1 {
				// Draw a block for each cell of the Tetromino
				ebitenutil.DrawRect(screen, float64((g.tetrominoX+j)*tileSize), float64((g.tetrominoY+i)*tileSize), float64(tileSize), float64(tileSize), g.currentColor)
			}
		}
	}

	// Display the game-over message and score if the game is over
	if g.isGameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press R to Restart\nScore: "+strconv.Itoa(g.score))
	} else {
		// Display the current score
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score))
	}

	// Display difficulty selection menu if the user is choosing difficulty
	if g.isChoosingDifficulty {
		ebitenutil.DebugPrint(screen, "Choose Difficulty:\n1. Easy\n2. Medium\n3. Hard")
	}
}

// Layout defines the layout for the game's screen size.
// It is called by the Ebiten library to determine the dimensions of the window.
// If the screen dimensions are invalid (non-positive), it returns a default size of 320x640.
//
// Parameters:
// - outsideWidth: The width of the window in pixels.
// - outsideHeight: The height of the window in pixels.
//
// Returns:
// - The width and height of the screen in pixels.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return default screen size if dimensions are invalid
	if g.screenWidth <= 0 || g.screenHeight <= 0 {
		return 320, 640
	}
	return g.screenWidth, g.screenHeight
}
