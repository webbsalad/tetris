package game

import (
	"math/rand"
	"time"
)

// spawnTetromino spawns a new Tetromino at the top of the board. The type of Tetromino spawned depends on the current game level.
// - Level 1: Only easy Tetrominoes are used.
// - Level 2: Easy and medium Tetrominoes are used.
// - Level 3: Easy, medium, and hard Tetrominoes are used.
// The new Tetromino is placed in the middle of the board horizontally, with its top row at the top of the board.
// If there is no room for the Tetromino to spawn, the game is over.
func (g *Game) spawnTetromino() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	var allTetrominoes [][][]int

	// Add Tetrominoes based on the current game level
	switch g.level {
	case 2:
		allTetrominoes = append(easyTetrominoes, mediumTetrominoes...)
	case 3:
		allTetrominoes = append(easyTetrominoes, mediumTetrominoes...)
		allTetrominoes = append(allTetrominoes, hardTetrominoes...)
	default:
		allTetrominoes = easyTetrominoes
	}

	// Select a random Tetromino from the list
	index := rand.Intn(len(allTetrominoes))
	g.currentTetromino = allTetrominoes[index]
	g.currentColor = tetrominoColors[index]

	// Position the Tetromino at the top of the board and in the center horizontally
	g.tetrominoX = g.boardWidth()/2 - len(g.currentTetromino[0])/2
	g.tetrominoY = 0

	// If there is no room to spawn the Tetromino, the game is over
	if !g.canMove(g.tetrominoX, g.tetrominoY, g.currentTetromino) {
		g.gameOver()
	}
}

// rotateTetromino rotates the current Tetromino 90 degrees clockwise.
// It checks if the rotated Tetromino can fit in the current position on the board.
// If the move is valid, the rotation is applied and a turn sound is played.
func (g *Game) rotateTetromino() {
	// Rotate the current Tetromino
	newTetromino := rotateMatrix(g.currentTetromino)

	// If the rotated Tetromino can fit, apply the rotation
	if g.canMove(g.tetrominoX, g.tetrominoY, newTetromino) {
		g.currentTetromino = newTetromino
		g.playTurnSound() // Play sound when the Tetromino is rotated
	}
}

// rotateMatrix rotates the given 2D matrix (Tetromino) 90 degrees clockwise.
// It creates a new matrix where the rows and columns are swapped and reversed.
func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)              // Number of rows
	m := len(matrix[0])           // Number of columns
	newMatrix := make([][]int, m) // Create a new matrix with swapped dimensions
	for i := range newMatrix {
		newMatrix[i] = make([]int, n) // Initialize each row
		for j := range newMatrix[i] {
			// Swap and reverse the matrix elements
			newMatrix[i][j] = matrix[n-j-1][i]
		}
	}
	return newMatrix
}
