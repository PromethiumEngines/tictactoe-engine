package tictactoe

import (
	"log"
	"testing"
	"time"
)

// Test3x3HorizontalMatch calls tictactoe.New() with the default options of ToWin 3,
// Height 3, Weight 3 for a valid return value, 0 should win in horizontal,
// vertical and 1 should win in diagonal.
func Test3x3HorizontalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, horizontal match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Horizontal game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Horizontal game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)

	// O turn
	tttGame.AddMove(1, 0, 1)

	tttGame.AddMove(0, 1, 0)

	tttGame.AddMove(1, 1, 1)

	// Here X should win
	tttGame.AddMove(0, 2, 0)
}

func Test3x3VerticalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, vertical match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Vertical game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Vertical game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)

	// O turn
	tttGame.AddMove(1, 1, 0)

	tttGame.AddMove(0, 0, 1)

	tttGame.AddMove(1, 1, 1)

	// Here X should win
	tttGame.AddMove(0, 0, 2)
}

func Test3x3DiagonalDecreasingMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, diagonal left to right match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Diagonal left to right game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Diagonal left to right game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// O turn
	tttGame.AddMove(1, 0, 0)

	// X turn
	tttGame.AddMove(0, 0, 1)

	tttGame.AddMove(1, 1, 1)

	tttGame.AddMove(0, 0, 2)

	// Here O should win
	tttGame.AddMove(1, 2, 2)
}

func Test3x3DiagonalIncreasingMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, diagonal right to left match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Diagonal right to left game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Diagonal right to left game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// X turn
	tttGame.AddMove(0, 2, 0)

	// 0 turn
	tttGame.AddMove(1, 2, 1)

	tttGame.AddMove(0, 1, 1)

	tttGame.AddMove(1, 2, 2)

	// Here O should win
	tttGame.AddMove(0, 0, 2)
}

func Test3x3DrawMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, diagonal match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		t.Fatalf(
			`Draw game concluded as a win, "%v" wins, in "%v" turns.`,
			*g.Winner,
			g.Turn,
		)
	})

	tttGame.OnDraw(func(g *Game) {
		log.Printf(
			`Draw game concluded, in "%v" turns. (In %v)`,
			g.Turn,
			time.Since(start),
		)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)

	// O turn
	tttGame.AddMove(1, 0, 1)

	tttGame.AddMove(0, 0, 2)

	tttGame.AddMove(1, 1, 1)

	tttGame.AddMove(0, 1, 0)

	tttGame.AddMove(1, 2, 2)

	tttGame.AddMove(0, 1, 2)

	tttGame.AddMove(1, 2, 0)

	tttGame.AddMove(0, 2, 1)
}

func Test5x4HorizontalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Width: 5})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, horizontal 5x4 match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Horizontal 5x4 game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Horizontal 5x4 game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)

	// O turn
	tttGame.AddMove(1, 0, 1)

	tttGame.AddMove(0, 1, 0)

	tttGame.AddMove(1, 0, 2)

	tttGame.AddMove(0, 2, 0)

	tttGame.AddMove(1, 0, 3)

	// Here X should win
	tttGame.AddMove(0, 3, 0)
}

func Test5x4VerticalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Width: 5})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, vertical 5x4 match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Vertical 5x4 game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Vertical 5x4 game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)

	// O turn
	tttGame.AddMove(1, 1, 0)

	tttGame.AddMove(0, 0, 1)

	tttGame.AddMove(1, 1, 1)

	tttGame.AddMove(0, 0, 2)

	tttGame.AddMove(1, 1, 2)

	// Here X should win
	tttGame.AddMove(0, 0, 3)
}

func Test5x4DiagonalDecreasingMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Width: 5})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, diagonal 5x4 left to right match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Diagonal 5x4 left to right game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Diagonal 5x4 left to right game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// O turn
	tttGame.AddMove(1, 0, 0)

	// X turn
	tttGame.AddMove(0, 1, 0)

	tttGame.AddMove(1, 1, 1)

	tttGame.AddMove(0, 1, 2)

	tttGame.AddMove(1, 2, 2)

	tttGame.AddMove(0, 2, 3)

	// Here O should win
	tttGame.AddMove(1, 3, 3)
}

func Test5x4DiagonalIncreasingMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Width: 5})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, diagonal 5x4 right to left match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(
			`Diagonal 5x4 right to left game concluded, "%v" wins, "%v" turns. (In %v)`,
			*g.Winner,
			g.Turn,
			time.Since(start),
		)
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(
			`Diagonal 5x4 right to left game concluded as a draw, in "%v" turns.`,
			g.Turn,
		)
	})

	// X turn
	tttGame.AddMove(0, 4, 0)

	// O turn
	tttGame.AddMove(1, 0, 1)

	tttGame.AddMove(0, 3, 1)

	tttGame.AddMove(1, 0, 2)

	tttGame.AddMove(0, 2, 2)

	tttGame.AddMove(1, 1, 3)

	// Here O should win
	tttGame.AddMove(0, 1, 3)
}

func Test5x4DrawMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Width: 5})
	if err != nil {
		t.Fatalf(
			"New(), options are invalid, draw 5x4 match, %v, want valid GameOptions",
			err,
		)
	}

	tttGame.OnWinner(func(g *Game) {
		t.Fatalf(
			`Draw 5x4 game concluded as a win, "%v" wins, in "%v" turns.`,
			*g.Winner,
			g.Turn,
		)
	})

	tttGame.OnDraw(func(g *Game) {
		log.Printf(
			`Draw 5x4 game concluded, in "%v" turns. (In %v)`,
			g.Turn,
			time.Since(start),
		)
	})

	// X turn
	tttGame.AddMove(0, 0, 0)
	tttGame.AddMove(0, 0, 1)

	tttGame.AddMove(0, 1, 2)
	tttGame.AddMove(0, 1, 3)

	tttGame.AddMove(0, 2, 0)
	tttGame.AddMove(0, 2, 1)

	tttGame.AddMove(0, 3, 2)
	tttGame.AddMove(0, 3, 3)

	tttGame.AddMove(0, 4, 0)
	tttGame.AddMove(0, 4, 1)

	// O turn
	tttGame.AddMove(1, 0, 2)
	tttGame.AddMove(1, 0, 3)

	tttGame.AddMove(1, 1, 0)
	tttGame.AddMove(1, 1, 1)

	tttGame.AddMove(1, 2, 2)
	tttGame.AddMove(1, 2, 3)

	tttGame.AddMove(1, 3, 0)
	tttGame.AddMove(1, 3, 1)
	tttGame.AddMove(1, 4, 2)
	// The match will end as draw here
	tttGame.AddMove(1, 4, 3)
}
