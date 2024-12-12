package game

// updateSpeed adjusts the speed of the game based on the current level.
// It decreases the move delay as the player progresses through levels.
// - Level 1: every 1200 frames, reduce delay by 1, but not below 15.
// - Level 2: every 600 frames, reduce delay by 1, but not below 10.
// - Level 3: every 300 frames, reduce delay by 1, but not below 5.
func (g *Game) updateSpeed() {
	switch g.level {
	case 1:
		if g.frameCount%1200 == 0 {
			g.moveDelay = max(g.moveDelay-1, 15)
		}
	case 2:
		if g.frameCount%600 == 0 {
			g.moveDelay = max(g.moveDelay-1, 10)
		}
	case 3:
		if g.frameCount%300 == 0 {
			g.moveDelay = max(g.moveDelay-1, 5)
		}
	}
}

// moveTetromino attempts to move the current Tetromino by the specified dx (horizontal) and dy (vertical) values.
// Returns true if the move is successful, false if the move is blocked by the game boundaries or other Tetrominoes.
func (g *Game) moveTetromino(dx, dy int) bool {
	oldX, oldY := g.tetrominoX, g.tetrominoY
	newX, newY := g.tetrominoX+dx, g.tetrominoY+dy
	if g.canMove(newX, newY, g.currentTetromino) {
		g.tetrominoX = newX
		g.tetrominoY = newY
		if oldX != newX || oldY != newY {
			g.playTurnSound() // Play sound if the Tetromino moved
		}
		return true
	}
	return false
}

// canMove checks if a Tetromino can move to the given x, y position without colliding with the boundaries
// or other blocks already placed on the board.
func (g *Game) canMove(x, y int, tetromino [][]int) bool {
	for i, row := range tetromino {
		for j, cell := range row {
			if cell == 1 {
				// Check boundaries and collision with placed blocks
				if x+j < 0 || x+j >= g.boardWidth() || y+i >= g.boardHeight() || (y+i >= 0 && g.board[y+i][x+j] == 1) {
					return false
				}
			}
		}
	}
	return true
}

// mergeTetromino places the current Tetromino onto the game board when it can no longer move.
// The Tetromino is "merged" by setting the corresponding cells on the board to 1.
func (g *Game) mergeTetromino() {
	for i, row := range g.currentTetromino {
		for j, cell := range row {
			if cell == 1 && g.tetrominoY+i >= 0 {
				g.board[g.tetrominoY+i][g.tetrominoX+j] = 1 // Set the current position on the board
			}
		}
	}
}

// clearLines checks the game board for any complete lines and clears them.
// It returns the number of lines cleared.
func (g *Game) clearLines() int {
	linesCleared := 0
	for i := 0; i < g.boardHeight(); i++ {
		full := true
		// Check if the current line is full (all cells are occupied)
		for j := 0; j < g.boardWidth(); j++ {
			if g.board[i][j] == 0 {
				full = false
				break
			}
		}
		if full {
			linesCleared++
			// Remove the full line and add a new empty line at the top
			g.board = append(g.board[:i], g.board[i+1:]...)
			newLine := make([]int, g.boardWidth())
			g.board = append([][]int{newLine}, g.board...)
		}
	}
	return linesCleared
}

// updateScore updates the score based on the number of lines cleared.
// Each line cleared adds a bonus to the score based on the square of the number of lines cleared.
// It also plays a sound effect for clearing the lines.
func (g *Game) updateScore(linesCleared int) {
	if linesCleared > 0 {
		g.linesCleared += linesCleared
		g.score += linesCleared * linesCleared * 100 // Score formula based on the number of lines cleared
		g.playPopSound()                             // Play sound when lines are cleared
	}
}
