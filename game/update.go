// Package game provides functionality for managing the game process,
// including sound handling, game state management, and player interactions.
package game

import "github.com/hajimehoshi/ebiten/v2"

// Update is the main game loop that updates the game state on each frame. It handles user input, game logic, and screen updates.
// The function handles three main states of the game:
// 1. **Choosing Difficulty**: If the player is in the difficulty selection screen, pressing keys 1, 2, or 3 starts the game with the corresponding difficulty level.
// 2. **Game Over**: If the game is over, pressing 'R' will restart the game and bring the player back to the difficulty selection screen.
// 3. **Gameplay**: During normal gameplay, it handles movement and rotation of the current Tetromino, game speed adjustments, line clearing, score updates, and spawning new Tetrominoes.
func (g *Game) Update() error {
	// Difficulty selection screen
	if g.isChoosingDifficulty {
		// Choose difficulty based on key press
		if ebiten.IsKeyPressed(ebiten.Key1) {
			g.Init(1) // Initialize game with difficulty level 1
		}
		if ebiten.IsKeyPressed(ebiten.Key2) {
			g.Init(2) // Initialize game with difficulty level 2
		}
		if ebiten.IsKeyPressed(ebiten.Key3) {
			g.Init(3) // Initialize game with difficulty level 3
		}
		return nil
	}

	// Game over state
	if g.isGameOver {
		// Restart game when 'R' key is pressed
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.isGameOver = false
			g.isChoosingDifficulty = true // Return to difficulty selection screen
		}
		return nil
	}

	// Normal game loop
	g.frameCount++
	g.moveTimer++
	g.rotateTimer++

	// Update game speed based on level
	g.updateSpeed()

	// Move Tetromino down at a regular interval
	if g.frameCount%g.moveDelay == 0 {
		if !g.moveTetromino(0, 1) {
			// If the Tetromino cannot move, merge it with the board and check for cleared lines
			g.mergeTetromino()
			linesCleared := g.clearLines()
			g.updateScore(linesCleared)
			// Spawn a new Tetromino after the current one is placed
			g.spawnTetromino()
		}
	}

	// Handle left movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && !g.movePressed && g.moveTimer >= moveDelay {
		g.moveTetromino(-1, 0)
		g.movePressed = true
		g.moveTimer = 0
	}

	// Handle right movement
	if ebiten.IsKeyPressed(ebiten.KeyRight) && !g.movePressed && g.moveTimer >= moveDelay {
		g.moveTetromino(1, 0)
		g.movePressed = true
		g.moveTimer = 0
	}

	// Handle down movement (accelerate Tetromino fall)
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.moveTimer >= moveDelay {
		g.moveTetromino(0, 1)
		g.moveTimer = 0
	}

	// Handle rotation
	if !g.rotatePressed {
		if ebiten.IsKeyPressed(ebiten.KeySpace) && g.rotateTimer >= rotateDelay {
			g.rotateTetromino() // Rotate the Tetromino
			g.rotateTimer = 0
			g.rotatePressed = true
		}
	}

	// Reset move and rotate states when keys are released
	if !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.movePressed = false
	}
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.rotatePressed = false
	}

	return nil
}
