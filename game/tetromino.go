package game

import (
	"math/rand"
	"time"
)

func (g *Game) spawnTetromino() {
	rand.Seed(time.Now().UnixNano())
	var allTetrominoes [][][]int

	switch g.level {
	case 2:
		allTetrominoes = append(easyTetrominoes, mediumTetrominoes...)
	case 3:
		allTetrominoes = append(easyTetrominoes, mediumTetrominoes...)
		allTetrominoes = append(allTetrominoes, hardTetrominoes...)
	default:
		allTetrominoes = easyTetrominoes
	}

	index := rand.Intn(len(allTetrominoes))
	g.currentTetromino = allTetrominoes[index]
	g.currentColor = tetrominoColors[index]
	g.tetrominoX = g.boardWidth()/2 - len(g.currentTetromino[0])/2
	g.tetrominoY = 0

	if !g.canMove(g.tetrominoX, g.tetrominoY, g.currentTetromino) {
		g.gameOver()
	}
}

func (g *Game) rotateTetromino() {
	newTetromino := rotateMatrix(g.currentTetromino)
	if g.canMove(g.tetrominoX, g.tetrominoY, newTetromino) {
		g.currentTetromino = newTetromino
		g.playTurnSound()
	}
}

func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	newMatrix := make([][]int, m)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
		for j := range newMatrix[i] {
			newMatrix[i][j] = matrix[n-j-1][i]
		}
	}
	return newMatrix
}
