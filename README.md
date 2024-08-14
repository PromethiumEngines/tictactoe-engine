![TicTacToe Engine Thumbnail](https://raw.githubusercontent.com/PromethiumEngines/.github/main/assets/ttt-engine-large.png)

[![Go Documentation](https://pkg.go.dev/badge/github.com/PromethiumEngines/tictactoe-engine.svg)](https://pkg.go.dev/github.com/PromethiumEngines/tictactoe-engine)
[![Go Report](https://goreportcard.com/badge/github.com/PromethiumEngines/tictactoe-engine)](https://goreportcard.com/report/github.com/PromethiumEngines/tictactoe-engine)
[![Go Version](https://img.shields.io/github/go-mod/go-version/PromethiumEngines/tictactoe-engine)](https://golang.org/doc/devel/release.html)
[![TicTacToe-Engine Version](https://img.shields.io/github/v/tag/PromethiumEngines/tictactoe-engine?label=release)](https://github.com/PromethiumEngines/tictactoe-engine/)
[![License](https://img.shields.io/github/license/PromethiumEngines/tictactoe-engine)](https://www.apache.org/licenses/LICENSE-2.0)

<h1></h1>

<h2 align="center">Highly configurable tictactoe engine</h2>

<h2>Installation</h2>

You can install the go package with this command in your project.

```
go get github.com/PromethiumEngines/tictactoe-engine
```

<h2>Documentation</h2>
The documentation and settings of this engine can be found on the go reference

<h2>Usage</h2>

This is just an example of how to run a game of tictactoe using the engine, all of this settings can be changed

```go
package main

import (
	"github.com/PromethiumEngines/tictactoe-engine"
	"log"
)

// Here we create a new game of tictactoe.
tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Width: 3})
if err != nil {
    // An error could occur if the settings are invalid
    panic(err)
}

// You can specify a function to be executed when someone wins
// In the engine X means a 0, and a O means a 1.
tttGame.OnWinner(func (g *Game) {
	log.Printf(`Someone won!! Winner: "%v""`, *g.Winner)
})

tttGame.OnDraw(func (g *Game) {
	// Turn is the current turn, so its the amount turns in the game.
	log.Printf(`It seems the match ended in a draw in "%v" turns`, g.Turn)
})

// For simplicity in adding moves the engine uses the x and y coordinates.
tttGame.AddMove(0, 0)

tttGame.AddMove(0, 1)
// The game will end when the engine detects a win or a draw emitting the correct event.
```
