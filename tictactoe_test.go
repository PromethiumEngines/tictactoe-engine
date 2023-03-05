package tictactoe

import (
	"log"
	"testing"
	"time"
)

// Test3x3<>Match calls tictactoe.New() with the default options of ToWin 3, Height 3, Weight 3
// for a valid return value, 0 should win in horizontal, vertical and 1 should win in diagonal
func Test3x3HorizontalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Weight: 3})
	if err != nil {
		t.Fatalf("New(), options are invalid, horizontal match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Horizontal game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Horizontal game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 3)

	tttGame.AddMove(0, 1)

	tttGame.AddMove(1, 4)

	// Here X should win
	tttGame.AddMove(0, 2)
}

func Test3x3VerticalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Weight: 3})
	if err != nil {
		t.Fatalf("New(), options are invalid, vertical match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Vertical game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Vertical game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 1)

	tttGame.AddMove(0, 3)

	tttGame.AddMove(1, 2)

	// Here X should win
	tttGame.AddMove(0, 6)
}

func Test3x3DiagonalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Weight: 3})
	if err != nil {
		t.Fatalf("New(), options are invalid, diagonal match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Diagonal game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Diagonal game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// O turn
	tttGame.AddMove(1, 0)

	// X turn
	tttGame.AddMove(0, 1)

	tttGame.AddMove(1, 4)

	tttGame.AddMove(0, 2)

	// Here O should win
	tttGame.AddMove(1, 8)
}

func Test3x3DrawMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Weight: 3})
	if err != nil {
		t.Fatalf("New(), options are invalid, diagonal match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		t.Fatalf(`Draw game concluded as a win, "%v" wins, in "%v" turns.`, *g.Winner, g.Turn)
	})

	tttGame.OnDraw(func(g *Game) {
		log.Printf(`Draw game concluded correctly, in "%v" turns. (In %v)`, g.Turn, time.Since(start))
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 1)

	tttGame.AddMove(0, 2)

	tttGame.AddMove(1, 4)

	tttGame.AddMove(0, 3)

	tttGame.AddMove(1, 5)

	tttGame.AddMove(0, 7)

	tttGame.AddMove(1, 6)

	tttGame.AddMove(0, 8)
}

func Test5x4HorizontalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Weight: 5})
	if err != nil {
		t.Fatalf("New(), options are invalid, horizontal 5x4 match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Horizontal 5x4 game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Horizontal 5x4 game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 5)

	tttGame.AddMove(0, 1)

	tttGame.AddMove(1, 6)

	tttGame.AddMove(0, 2)

	tttGame.AddMove(1, 7)

	// Here X should win
	tttGame.AddMove(0, 3)
}

func Test5x4VerticalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Weight: 5})
	if err != nil {
		t.Fatalf("New(), options are invalid, horizontal 5x4 match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Vertical 5x4 game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Vertical 5x4 game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 1)

	tttGame.AddMove(0, 5)

	tttGame.AddMove(1, 2)

	tttGame.AddMove(0, 10)

	tttGame.AddMove(1, 3)

	// Here X should win
	tttGame.AddMove(0, 15)
}

func Test5x4DiagonalMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Weight: 5})
	if err != nil {
		t.Fatalf("New(), options are invalid, horizontal 5x4 match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		log.Printf(`Diagonal 5x4 game concluded, "%v" wins, "%v" turns. (In %v)`, *g.Winner, g.Turn, time.Since(start))
	})

	tttGame.OnDraw(func(g *Game) {
		t.Fatalf(`Diagonal 5x4 game concluded as a draw, in "%v" turns.`, g.Turn)
	})

	// O turn
	tttGame.AddMove(1, 0)

	// X turn
	tttGame.AddMove(0, 1)

	tttGame.AddMove(1, 6)

	tttGame.AddMove(0, 2)

	tttGame.AddMove(1, 12)

	tttGame.AddMove(0, 3)

	// Here O should win
	tttGame.AddMove(1, 18)
}

func Test5x4DrawMatch(t *testing.T) {
	start := time.Now()
	tttGame, err := New(&GameOptions{ToWin: 4, Height: 4, Weight: 5})
	if err != nil {
		t.Fatalf("New(), options are invalid, horizontal 5x4 match, %v, want valid GameOptions", err)
	}

	tttGame.OnWinner(func(g *Game) {
		t.Fatalf(`Draw 5x4 game concluded as a win, "%v" wins, in "%v" turns.`, *g.Winner, g.Turn)
	})

	tttGame.OnDraw(func(g *Game) {
		log.Printf(`Draw 5x4 game concluded correctly, in "%v" turns. (In %v)`, g.Turn, time.Since(start))
	})

	// X turn
	tttGame.AddMove(0, 0)

	// O turn
	tttGame.AddMove(1, 1)

	tttGame.AddMove(0, 2)

	tttGame.AddMove(1, 3)

	tttGame.AddMove(0, 4)

	tttGame.AddMove(1, 5)

	tttGame.AddMove(0, 6)

	tttGame.AddMove(1, 7)

	tttGame.AddMove(0, 8)

	tttGame.AddMove(1, 9)

	tttGame.AddMove(0, 11)

	tttGame.AddMove(1, 10)

	tttGame.AddMove(0, 13)

	tttGame.AddMove(1, 12)

	tttGame.AddMove(0, 15)

	tttGame.AddMove(1, 14)

	tttGame.AddMove(0, 15)

	tttGame.AddMove(1, 16)

	tttGame.AddMove(0, 17)

	tttGame.AddMove(1, 18)

	// The match will end as draw here
	tttGame.AddMove(0, 19)
}
