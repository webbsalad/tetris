package game

// Init initializes the game with the specified difficulty level.
// It sets up the game state, including board dimensions based on the level,
// initializes the board, and spawns a new Tetromino.
// The method also updates the screen size and sets the move delay.
//
// Parameters:
// - level: The difficulty level for the game, which affects the board size.
func (g *Game) Init(level int) {
	g.isChoosingDifficulty = false
	g.level = level
	g.boardWidthExtra = g.boardWidth() - 10
	g.boardHeightExtra = g.boardHeight() - 20
	g.board = make([][]int, g.boardHeight())
	for i := range g.board {
		g.board[i] = make([]int, g.boardWidth())
	}
	g.moveDelay = initialMoveDelay
	g.updateScreenSize()
	g.spawnTetromino()
}

// updateScreenSize updates the screen size based on the current board dimensions.
// The screen width is calculated by multiplying the board width by the tile size,
// and the screen height is calculated by multiplying the board height by the tile size.
func (g *Game) updateScreenSize() {
	g.screenWidth = g.boardWidth() * tileSize
	g.screenHeight = g.boardHeight() * tileSize
}

// boardWidth returns the width of the board based on the current game level.
// The level affects the board width, with the following mapping:
// - Level 1: 10 columns
// - Level 2: 12 columns
// - Level 3: 14 columns
//
// Returns:
// - The width of the board for the current level.
func (g *Game) boardWidth() int {
	switch g.level {
	case 2:
		return 12
	case 3:
		return 14
	default:
		return 10
	}
}

// boardHeight returns the height of the board based on the current game level.
// The level affects the board height, with the following mapping:
// - Level 1: 20 rows
// - Level 2: 22 rows
// - Level 3: 24 rows
//
// Returns:
// - The height of the board for the current level.
func (g *Game) boardHeight() int {
	switch g.level {
	case 2:
		return 22
	case 3:
		return 24
	default:
		return 20
	}
}
