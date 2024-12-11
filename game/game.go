package game

import (
	"image/color"
	"log"
)

type Game struct {
	board                  [][]int
	currentTetromino       [][]int
	currentColor           color.RGBA
	tetrominoX, tetrominoY int
	frameCount             int
	moveTimer              int
	rotateTimer            int
	movePressed            bool
	rotatePressed          bool
	score                  int
	linesCleared           int
	isGameOver             bool
	level                  int
	moveDelay              int
	isChoosingDifficulty   bool
	boardWidthExtra        int
	boardHeightExtra       int
	screenWidth            int
	screenHeight           int

	audioContext *AudioContextWrapper
	soundDead    *SoundPlayer
	soundPop     *SoundPlayer
	soundTurn    *SoundPlayer
}

func New() *Game {
	g := &Game{
		isChoosingDifficulty: true,
		screenWidth:          320,
		screenHeight:         640,
	}
	g.initAudio()
	return g
}

func (g *Game) gameOver() {
	log.Println("Game Over! Final Score:", g.score)
	g.isGameOver = true
	g.playDeadSound()
}
