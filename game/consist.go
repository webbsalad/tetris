package game

import "image/color"

// Constants related to the game configuration, such as tile size, delays, and other parameters.
const (
	tileSize         = 32 // Size of each tile in pixels.
	initialMoveDelay = 30 // Initial delay (in frames) between moves of the Tetromino.
	moveDelay        = 5  // Delay (in frames) between moves of the Tetromino after the initial move.
	rotateDelay      = 15 // Delay (in frames) for rotating the Tetromino.
)

// easyTetrominoes defines the shape configurations for Tetrominoes at the "easy" difficulty level.
// Each Tetromino is represented as a 2D array of integers, where 1 represents a filled block.
var easyTetrominoes = [][][]int{
	{
		{1, 1, 1, 1}, // I Tetromino
	},
	{
		{1, 1, 1},
		{0, 1, 0}, // T Tetromino
	},
	{
		{1, 1},
		{1, 1}, // O Tetromino
	},
	{
		{1, 1, 0},
		{0, 1, 1}, // Z Tetromino
	},
	{
		{0, 1, 1},
		{1, 1, 0}, // S Tetromino
	},
	{
		{1, 1, 1},
		{1, 0, 0}, // L Tetromino
	},
	{
		{1, 1, 1},
		{0, 0, 1}, // J Tetromino
	},
}

// mediumTetrominoes defines the shape configurations for Tetrominoes at the "medium" difficulty level.
// Each Tetromino is represented as a 2D array of integers, where 1 represents a filled block.
var mediumTetrominoes = [][][]int{
	{
		{1, 1, 1, 1},
		{0, 1, 1, 0}, // Expanded T Tetromino
	},
	{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 1}, // "L"-shaped Tetromino
	},
	{
		{1}, // Single block Tetromino
	},
}

// hardTetrominoes defines the shape configurations for Tetrominoes at the "hard" difficulty level.
// Each Tetromino is represented as a 2D array of integers, where 1 represents a filled block.
var hardTetrominoes = [][][]int{
	{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}, // Complex Tetromino
	},
	{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0}, // Cross-shaped Tetromino
	},
	{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 1}, // Staircase-shaped Tetromino
	},
}

// tetrominoColors defines the colors associated with each Tetromino shape.
// The colors are represented in RGBA format where R, G, B, and A correspond to the red, green, blue, and alpha channels respectively.
var tetrominoColors = []color.RGBA{
	{0, 255, 255, 255},   // I Tetromino - Cyan
	{255, 0, 255, 255},   // T Tetromino - Magenta
	{255, 255, 0, 255},   // O Tetromino - Yellow
	{255, 0, 0, 255},     // Z Tetromino - Red
	{0, 255, 0, 255},     // S Tetromino - Green
	{255, 165, 0, 255},   // L Tetromino - Orange
	{0, 0, 255, 255},     // J Tetromino - Blue
	{128, 128, 128, 255}, // Cross-shaped Tetromino - Grey
	{128, 0, 128, 255},   // "L"-shaped Tetromino - Purple
	{0, 0, 0, 0},         // Invisible Tetromino (unused)
	{255, 69, 0, 255},    // Complex Tetromino - Red-Orange
	{0, 128, 128, 255},   // Expanded T Tetromino - Teal
	{73, 92, 122, 255},   // Staircase-shaped Tetromino - Dark Blue
}
