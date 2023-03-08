![TicTacToe Engine Thumbnail](https://raw.githubusercontent.com/PromethiumEngines/.github/main/assets/ttt-engine-large.png)

[![Go Report](https://goreportcard.com/badge/github.com/PromethiumEngines/tictactoe-engine)](https://goreportcard.com/report/github.com/PromethiumEngines/tictactoe-engine)
[![Go Version](https://img.shields.io/github/go-mod/go-version/PromethiumEngines/tictactoe-engine)](https://golang.org/doc/devel/release.html)
[![TicTacToe-Engine Version](https://img.shields.io/github/v/tag/PromethiumEngines/tictactoe-engine?label=release)](https://github.com/PromethiumEngines/tictactoe-engine/)
[![License](https://img.shields.io/github/license/PromethiumEngines/tictactoe-engine)](https://www.apache.org/licenses/LICENSE-2.0)
[![Discord](https://discord.com/api/guilds/761370919419117598/widget.png)](https://discord.gg/g3ZbCmShD4)

<h1></h1>

<h2 align="center">An engine to simulate a game of tictactoe highly configurable made by <b>PromethiumEngines</b></h2>

<h2>Installation</h2>

You can install the go package with this command in your project.

```
go get github.com/PromethiumEngines/tictactoe-engine
```

<h2>Usage</h2>

This engine is easy to use:

```go
package main

import (
	"github.com/PromethiumEngines/tictactoe-engine"
	"log"
)

// Here we create a new game of tictactoe.
// (All this settings can be changed)
tttGame, err := New(&GameOptions{ToWin: 3, Height: 3, Weight: 3})
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

// You can start adding moves here
// If your map is 3x3 it means the max-tiles are 9,
// but we start from 0 so be carefully where you put your symbol!
tttGame.AddMove(0, 0)

tttGame.AddMove(0, 1)
// The game will end when the engine detects a win or a draw.
```