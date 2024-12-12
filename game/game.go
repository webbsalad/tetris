package game

import (
	"image/color"
	"log"
)

// Game represents the state of the game. It holds all the variables required to track the game's progress,
// including the game board, current Tetromino, score, game state, audio context, and other game parameters.
type Game struct {
	board                  [][]int              // The game board, represented as a 2D array of integers.
	currentTetromino       [][]int              // The current Tetromino, represented as a 2D array of integers.
	currentColor           color.RGBA           // The color of the current Tetromino.
	tetrominoX, tetrominoY int                  // The coordinates of the current Tetromino's position on the board.
	frameCount             int                  // The frame count, used for timing and game updates.
	moveTimer              int                  // A timer to control the speed of moving the Tetromino.
	rotateTimer            int                  // A timer to control the rotation speed of the Tetromino.
	movePressed            bool                 // A flag indicating whether the move action was pressed.
	rotatePressed          bool                 // A flag indicating whether the rotate action was pressed.
	score                  int                  // The player's current score.
	linesCleared           int                  // The number of lines the player has cleared.
	isGameOver             bool                 // A flag indicating whether the game is over.
	level                  int                  // The current level of the game.
	moveDelay              int                  // The delay between moves, affecting the speed of the game.
	isChoosingDifficulty   bool                 // A flag indicating whether the player is choosing the difficulty.
	boardWidthExtra        int                  // The width of the game board, minus some padding.
	boardHeightExtra       int                  // The height of the game board, minus some padding.
	screenWidth            int                  // The width of the game screen.
	screenHeight           int                  // The height of the game screen.
	audioContext           *AudioContextWrapper // The audio context used for sound playback.
	soundDead              *SoundPlayer         // The sound player for the "game over" sound.
	soundPop               *SoundPlayer         // The sound player for the "pop" sound.
	soundTurn              *SoundPlayer         // The sound player for the "turn" sound.
}

// New initializes a new game instance and returns a pointer to it.
// It sets up the default screen size and the difficulty selection flag,
// then initializes the audio context and sound players.
func New() *Game {
	g := &Game{
		isChoosingDifficulty: true,
		screenWidth:          320, // Default screen width.
		screenHeight:         640, // Default screen height.
	}
	g.initAudio() // Initialize the audio context and load sounds.
	return g
}

// gameOver is called when the game ends. It sets the game-over state, logs the final score,
// and plays the "game over" sound.
func (g *Game) gameOver() {
	log.Println("Game Over! Final Score:", g.score)
	g.isGameOver = true
	g.playDeadSound() // Play the "dead" sound when the game is over.
}
