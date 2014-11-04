package game

import (
	"board"
	"rules"
)

// Game is the simplist way to represent a Go game
type Game struct {
	prev      board.Board //previous state (used to check for ko)
	curr      board.Board //current state of board
	blackTurn bool        //true if black's turn to play
}

//return empty game of dimensions dim*dim
func initGame(dim int) Game {
	return Game{prev: board.InitBoard(dim), curr: board.InitBoard(dim),
		blackTurn: true}
}

// returns true,0 is move is legal
// otherwise, return false,<reason>
// define int constants to provide more info
func (g *Game) isLegal(i int, j int) (bool, int) {
	// it's illegal to play on a non-empty point
	if g.curr[i][j] != board.EMPTY {
		return false, 1
	}

	// preview state of board after playing this move
	tmp := board.CloneBoard(g.curr)
	if g.blackTurn {
		tmp[i][j] = board.BLACK
	} else {
		tmp[i][j] = board.WHITE
	}
	rules.ProcessBoard(&tmp, g.blackTurn)

	// it's illegal to commit suicide
	if tmp[i][j] == board.EMPTY {
		return false, 2
	}

	// it's illegal to repeat state, check if ko
	if board.EqualBoard(tmp, g.prev) {
		return false, 3
	}

	// if we don't violate previous rules, it's legal
	return true, 0
}

// pre: (i,j) is a legal move
// post: move has been played, state of game is fully upto date
func (g *Game) playMove(i int, j int) {
	tmp := board.CloneBoard(g.curr)
	if g.blackTurn {
		tmp[i][j] = board.BLACK
	} else {
		tmp[i][j] = board.WHITE
	}
	rules.ProcessBoard(&tmp, g.blackTurn)
	g.prev = g.curr
	g.curr = tmp
}
