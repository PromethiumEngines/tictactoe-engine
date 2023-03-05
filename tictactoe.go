package tictactoe

import (
	"errors"
)

// A vector is a map of the symbols (X and O), can be treated as weight (X axis)
type Vector = map[int]*int
// A matrix is an array of vectors, the position of a vector can be treated as Height (Y axis)
type Matrix = map[int]Vector

// The representation of a tictactoe game
type Game struct {
	GameOptions `json:"game_options"`
	Ended bool `json:"ended"`
	Winner *int `json:"winner,omitempty"`
	Turn int `json:"turn"`
	Map Matrix `json:"map"`
	
	maxTiles int

	onWinner func(g *Game)
	onDraw func(g *Game)
}

// Sets a function to get called when someone wins
func (g *Game) OnWinner(callback func(g *Game)) *Game {
	g.onWinner = callback

	return g
}

// Sets a function to get called when someone wins
func (g *Game) OnDraw(callback func(g *Game)) *Game {
	g.onDraw = callback

	return g
}

// Add a move to the map
// player: A 0 represents the X and a 1 represents a O
// toWhere: Starts from 0, if number is more than the multiplication of the height and weight returns an error
func (g *Game) AddMove(player int, toWhere int) error {
	if g.Ended {
		return errors.New("TicTacToe-Engine: Game has already ended")
	} else if toWhere > g.maxTiles - 1 {
		return errors.New("TicTacToe-Engine: ToWhere is more than the max tiles of the game (multiplication of the height and weight)")
	}

	// toWhere/g.Weight we get the position of the vector (Y axis)
	var toVector int = toWhere/g.Weight

	// int(toWhere-(toVector*g.Weight)) then we get the position (X axis) of the section in that vector
	// that section is where we will put our symbol (X or O)
	g.Map[toVector][int(toWhere-(toVector*g.Weight))] = &player
	g.Turn++

	g.checkMap()

	return nil
}

// func maplog(d Matrix) {
// 	a := map[int]map[int]int{}

// 	for i, aa := range d {
// 		u := map[int]int{}
// 		for o, ee := range aa {
// 			u[o] = *ee
// 		}
// 		a[i] = u
// 	}

// 	jsd, _ := json.Marshal(a)
// 	log.Println(string(jsd))
// }

// Using a simple algorithm checks for the winner or draw in this turn
func (g *Game) checkMap() {
	//This is to check if there is a draw
	totalLength := 0
	
	// We iterate each vector of the matrix
	for pos, vector := range g.Map {
		totalLength = totalLength + len(vector)
		// Here we create this var to stop the whole for in case someone wins
		breaked := false

		if totalLength == g.maxTiles {
			g.Ended = true
			g.onDraw(g)

			break
		}

		// g.Weight is the X axis in the map, thats why we iterate through each section of this vector
		for secPos := 0; secPos < g.Weight; secPos++ {
			// Here we check if there is space to win in the X axis 
			// We substract the current position of the section in the vector of weight (X axis) and check if its more or equal to the ToWin
			// Here we search for the next section in the current vector, if there is one and is the same symbol as the current section
			// we will enter the if, if not we will continue to the next section
			if nextPos := vector[secPos + 1]; (g.Weight - secPos) >= g.ToWin &&
			nextPos != nil &&
			vector[secPos] != nil &&
			*nextPos == *vector[secPos] {
				// Here we declare the symbol in the last section
				var last *int = nextPos

				// Now we iterate for the next sections to search for a connection of n (ToWin) symbols
				// we substract 1 from here because we want to make ToWin the slice index format (0,1,2) not (1,2,3)
				// we substract another 1 (2 in total) because we already got the next section (nextPos)
				for i := 0; i < g.ToWin - 2; i++ {
					// Here we calculate the NextPos of i substracting to secPos the 2 we already explain
					if nNextPos := vector[secPos + 2 + i]; nNextPos != nil && *nNextPos == *last {
						// Here we also substract 3 from toWin (in total)
						// the 1 alone is for making ToWin index format
						// the 2 and the i, is because we at least already got 2 positions, one outside of the for
						// and one inside the for, we do this because i starts at 0
						if i >= (g.ToWin - 1) - 2 {
							g.Winner = vector[secPos]
							g.Ended = true
							g.onWinner(g)

							// someone wins and we break the whole for
							breaked = true

							break
						}
					} else {
						break
					}

					last = vector[secPos + i]

					// here we just end the for normally because we want to check the next sections
				}
			// Here we check if the length of the matrix is enought to make a connection of n (toWin) symbols but diagonaly
			// also check if the nextvector exists
			// We check if in the next vector there is the same symbol but diagonaly
			} else if nextVectorPos := g.Map[pos + 1]; len(g.Map) - pos >= g.ToWin &&
			nextVectorPos[secPos + 1] != nil &&
			vector[secPos] != nil &&
			*nextVectorPos[secPos + 1] == *vector[secPos] {
				// Here we declare the symbol in the last section
				var last *int = nextVectorPos[secPos + 1]
				
				// Now we iterate for the next vectors to search for a connection of n (ToWin) symbols diagonaly
				// we substract 1 from here because we want to make ToWin the slice index format (0,1,2) not (1,2,3)
				// we substract another 1 (2 in total) because we already got the next vector (nextVectorPos)
				for i := 0; i < g.ToWin - 2; i++ {
					// Here we get the position of the next vector in the i index
					// we sum 2 because we already got 2 vectors (vector index 0 and vector 1)
					nNextVector := g.Map[pos + 2 + i]
					// Here we sum the position of the current section and 1 because we need to advance 1 section to do a diagonal line
					// and sum the "i" to get if there is the 3 (or more) symbol in the line
					if nNextPos := nNextVector[secPos + 2 + i]; nNextPos != nil && *nNextPos == *last {
						// Substract 1 to use the slice format (0, 1, 2) and substract another 2 because we already got 2
						if i >= (g.ToWin - 1) - 2 {
							g.Winner = vector[secPos]
							g.Ended = true
							g.onWinner(g)

							// someone wins and we break the whole for
							breaked = true

							break
						}
					} else {
						break
					}

					last = nNextVector[secPos + 2 + i]
				}
			// The algorithm to check if there is an vertical connection (Y axis) is almost the same as the diagonal one
			// we dont add 2 in positions because we want the same position as in secPos
			} else if nextVectorPos := g.Map[pos + 1]; len(g.Map) - pos >= g.ToWin &&
			nextVectorPos[secPos] != nil &&
			vector[secPos] != nil &&
			*nextVectorPos[secPos] == *vector[secPos] {
				var last *int = nextVectorPos[secPos]

				for i := 0; i < g.ToWin - 2; i++ {
					nNextVector := g.Map[pos + 2 + i]

					if nNextPos := nNextVector[secPos]; nNextPos != nil && *nNextPos == *last {
						if i >= (g.ToWin - 1) - 2 {
							g.Winner = vector[secPos]
							g.Ended = true
							g.onWinner(g)

							breaked = true

							break
						}
					} else {
						break
					}

					last = nNextVector[secPos]
				}
			}
		}

		if breaked {
			break
		}
	}
}

// The current options of a game
// ToWin is the quantity of connected symbols needed for a player to win
// Height and weight cant be equal or less than 1
type GameOptions struct {
	ToWin int `json:"to_win"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

// Create a new instance of a game
// Returns an error if invalid options
func New(options *GameOptions) (*Game, error) {
	if options.Height <= 2 {
		return &Game{}, errors.New("TicTacToe-Engine: Invalid height, less or equal than 2")
	} else if options.Weight <= 2 {
		return &Game{}, errors.New("TicTacToe-Engine: Invalid weight, less or equal than 2")
	} else if options.ToWin <= 2 {
		return &Game{}, errors.New("TicTacToe-Engine: Invalid towin, less or equal than 2")
	}

	Map := Matrix{}

	for i := 0; i < options.Height; i++ {
		Map[i] = Vector{}
	}

	return &Game{
		GameOptions: *options,
		Ended: false,
		Turn: 0,
		Map: Map,
		maxTiles: (options.Height * options.Weight),
	}, nil
}