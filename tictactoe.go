package tictactoe

import (
	"errors"
)

// Vector is a map of the symbols (X and O) on the Y axis.
type Vector = map[int]*int

// Matrix is a map of Vectors representing the board on the X axis.
type Matrix = map[int]Vector

// GameOptions The current options of a game
// ToWin is the quantity of connected symbols needed for a player to win
// Height and Width cant be equal or less than 1.
type GameOptions struct {
	ToWin  int `json:"to_win"`
	Height int `json:"height"`
	Width  int `json:"width"`
}

// Game is the representation of all the data in the current game.
type Game struct {
	GameOptions `       json:"game_options"`
	Ended       bool   `json:"ended"`
	Winner      *int   `json:"winner,omitempty"`
	Turn        int    `json:"turn"`
	Map         Matrix `json:"map"`

	maxTurns int

	onWinner func(g *Game)
	onDraw   func(g *Game)
}

// New Creates a new instance of a game
// Returns an error if invalid options.
func New(options *GameOptions) (*Game, error) {
	if options.Height <= 2 {
		return nil, errors.New(
			"TicTacToe-Engine: Invalid Height, less or equal than 2",
		)
	} else if options.Width <= 2 {
		return nil, errors.New("TicTacToe-Engine: Invalid Width, less or equal than 2")
	} else if options.ToWin <= 2 {
		return nil, errors.New("TicTacToe-Engine: Invalid ToWin, less or equal than 2")
	} else if options.Height < options.ToWin && options.Width < options.ToWin {
		return nil, errors.New("TicTacToe-Engine: Invalid ToWin, less than height or weight")
	}

	Map := Matrix{}

	for i := 0; i < options.Width; i++ {
		Map[i] = Vector{}
	}

	return &Game{
		GameOptions: *options,
		Ended:       false,
		Turn:        0,
		Map:         Map,
		maxTurns:    options.Height * options.Width,
	}, nil
}

// OnWinner lets you set a callback that gets executed when the game is
// concluded as someone winning.
func (g *Game) OnWinner(callback func(g *Game)) *Game {
	g.onWinner = callback

	return g
}

// OnDraw lets you set a callback that gets executed when the game is concluded
// as a draw.
func (g *Game) OnDraw(callback func(g *Game)) *Game {
	g.onDraw = callback

	return g
}

// AddMove lets you set a new move of a player.
// Player: A 0 represents the X, and a 1 represents a O.
func (g *Game) AddMove(player int, x int, y int) error {
	if g.Ended {
		return errors.New("TicTacToe-Engine: Game has already ended")
	} else if x > g.Width-1 || y > g.Height-1 {
		return errors.New("TicTacToe-Engine: Coordinates (X or Y) are out of bounds")
	}

	g.Map[x][y] = &player
	g.Turn++

	g.checkMap()

	return nil
}

func (g *Game) checkMap() {
	for x := range g.Map {
		for y := range g.Map[x] {
			if g.checkDirection(x, y, 1, 0) || // Horizontal
				g.checkDirection(x, y, 0, 1) || // Vertical
				g.checkDirection(x, y, 1, 1) || // Diagonal \
				g.checkDirection(x, y, 1, -1) { // Diagonal /
				g.Winner = g.Map[x][y]
				g.Ended = true
				g.onWinner(g)

				return
			} else if g.maxTurns == g.Turn {
				g.Ended = true
				g.onDraw(g)

				return
			}
		}
	}
}

func (g *Game) checkDirection(x, y, dx, dy int) bool {
	player := g.Map[x][y] // Get the player symbol at the starting position
	if player == nil {
		return false // If the starting position is empty, no winning sequence is possible
	}

	count := 1 // Start counting from the current position

	// Check in the direction (dx, dy)
	for pos := 1; pos < g.ToWin; pos++ {
		// Check both directions with a single loop
		for sign := -1; sign <= 1; sign += 2 {
			nextPosX, nextPosY := x+dx*pos*sign, y+dy*pos*sign

			if nextPosX < 0 || nextPosY < 0 || nextPosX >= g.Width || nextPosY >= g.Height ||
				g.Map[nextPosX][nextPosY] == nil || *g.Map[nextPosX][nextPosY] != *player {
				break // Stop if out of bounds or sequence breaks
			}
			count++
		}
	}

	// Return true if a winning sequence is found
	return count >= g.ToWin
}
