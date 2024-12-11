package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Update() error {
	if g.isChoosingDifficulty {
		if ebiten.IsKeyPressed(ebiten.Key1) {
			g.Init(1)
		}
		if ebiten.IsKeyPressed(ebiten.Key2) {
			g.Init(2)
		}
		if ebiten.IsKeyPressed(ebiten.Key3) {
			g.Init(3)
		}
		return nil
	}

	if g.isGameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.isGameOver = false
			g.isChoosingDifficulty = true
		}
		return nil
	}

	g.frameCount++
	g.moveTimer++
	g.rotateTimer++

	g.updateSpeed()

	if g.frameCount%g.moveDelay == 0 {
		if !g.moveTetromino(0, 1) {
			g.mergeTetromino()
			linesCleared := g.clearLines()
			g.updateScore(linesCleared)
			g.spawnTetromino()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && !g.movePressed && g.moveTimer >= moveDelay {
		g.moveTetromino(-1, 0)
		g.movePressed = true
		g.moveTimer = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) && !g.movePressed && g.moveTimer >= moveDelay {
		g.moveTetromino(1, 0)
		g.movePressed = true
		g.moveTimer = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.moveTimer >= moveDelay {
		g.moveTetromino(0, 1)
		g.moveTimer = 0
	}

	if !g.rotatePressed {
		if ebiten.IsKeyPressed(ebiten.KeySpace) && g.rotateTimer >= rotateDelay {
			g.rotateTetromino()
			g.rotateTimer = 0
			g.rotatePressed = true
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.movePressed = false
	}
	if !ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.rotatePressed = false
	}

	return nil
}
