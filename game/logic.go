package game

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

func (g *Game) moveTetromino(dx, dy int) bool {
	oldX, oldY := g.tetrominoX, g.tetrominoY
	newX, newY := g.tetrominoX+dx, g.tetrominoY+dy
	if g.canMove(newX, newY, g.currentTetromino) {
		g.tetrominoX = newX
		g.tetrominoY = newY
		if oldX != newX || oldY != newY {
			g.playTurnSound()
		}
		return true
	}
	return false
}

func (g *Game) canMove(x, y int, tetromino [][]int) bool {
	for i, row := range tetromino {
		for j, cell := range row {
			if cell == 1 {
				if x+j < 0 || x+j >= g.boardWidth() || y+i >= g.boardHeight() || (y+i >= 0 && g.board[y+i][x+j] == 1) {
					return false
				}
			}
		}
	}
	return true
}

func (g *Game) mergeTetromino() {
	for i, row := range g.currentTetromino {
		for j, cell := range row {
			if cell == 1 && g.tetrominoY+i >= 0 {
				g.board[g.tetrominoY+i][g.tetrominoX+j] = 1
			}
		}
	}
}

func (g *Game) clearLines() int {
	linesCleared := 0
	for i := 0; i < g.boardHeight(); i++ {
		full := true
		for j := 0; j < g.boardWidth(); j++ {
			if g.board[i][j] == 0 {
				full = false
				break
			}
		}
		if full {
			linesCleared++
			g.board = append(g.board[:i], g.board[i+1:]...)
			newLine := make([]int, g.boardWidth())
			g.board = append([][]int{newLine}, g.board...)
		}
	}
	return linesCleared
}

func (g *Game) updateScore(linesCleared int) {
	if linesCleared > 0 {
		g.linesCleared += linesCleared
		g.score += linesCleared * linesCleared * 100
		g.playPopSound()
	}
}
