package game

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

func (g *Game) updateScreenSize() {
	g.screenWidth = g.boardWidth() * tileSize
	g.screenHeight = g.boardHeight() * tileSize
}

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
