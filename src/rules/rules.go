package rules

import (
	"board"
)

// The function ProcessBoard will take a state after a legal move has been
// played and adjusts the state to be legal (IE: removes dead groups).

// we use this to grab entry from state whenever bounds-checking is an issue
func getEntry(i int, j int, state *board.Board) board.Piece {
	var dim int = len((*state))
	if i < 0 || i >= dim || j < 0 || j >= dim {
		return board.NULLS //we overshot the board
	} else {
		return (*state)[i][j]
	}
}

// m represents a marked type piece
func removeGroup(i int, j int, m board.Piece, state *board.Board) {
	(*state)[i][j] = board.EMPTY

	if getEntry(i+1, j, state) == m {
		removeGroup(i+1, j, m, state)
	}
	if getEntry(i-1, j, state) == m {
		removeGroup(i-1, j, m, state)
	}
	if getEntry(i, j+1, state) == m {
		removeGroup(i, j+1, m, state)
	}
	if getEntry(i, j-1, state) == m {
		removeGroup(i+1, j-1, m, state)
	}
}

// pre: state[i][j] represents either BLACK or WHITE
// post: all pieces of value p connected to i,j are marked. The number of
//		liberties of the group are returned
func countLiberties(i int, j int, p board.Piece, m board.Piece, state *board.Board) int {
	var lib_count int = 0

	(*state)[i][j] = m //avoiding infinite recursion is good

	if getEntry(i+1, j, state) == board.EMPTY {
		(*state)[i+1][j] = board.MARKED_E
		lib_count += 1
	}
	if getEntry(i-1, j, state) == board.EMPTY {
		(*state)[i-1][j] = board.MARKED_E
		lib_count += 1
	}
	if getEntry(i, j+1, state) == board.EMPTY {
		(*state)[i][j+1] = board.MARKED_E
		lib_count += 1
	}
	if getEntry(i, j-1, state) == board.EMPTY {
		(*state)[i][j-1] = board.MARKED_E
		lib_count += 1
	}

	if getEntry(i+1, j, state) == p {
		lib_count += countLiberties(i+1, j, p, m, state)
	}
	if getEntry(i-1, j, state) == p {
		lib_count += countLiberties(i-1, j, p, m, state)
	}
	if getEntry(i, j+1, state) == p {
		lib_count += countLiberties(i, j+1, p, m, state)
	}
	if getEntry(i, j-1, state) == p {
		lib_count += countLiberties(i, j-1, p, m, state)
	}

	return lib_count
}

// remove temporary marked pieces
func cleanBoard(state *board.Board) {
	var dim = len(*state)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if (*state)[i][j] == board.MARKED_B {
				(*state)[i][j] = board.BLACK
			}
			if (*state)[i][j] == board.MARKED_W {
				(*state)[i][j] = board.WHITE
			}
			if (*state)[i][j] == board.MARKED_E {
				(*state)[i][j] = board.EMPTY
			}
		}
	}
}

//If we desire to implement this function, it will print statistics about the
// state to stdout
func printStats(state *board.Board) {
}

func removeGroupIfDead(i int, j int, p board.Piece, m board.Piece, state *board.Board) {
	if (*state)[i][j] == p {
		if countLiberties(i, j, p, m, state) == 0 {
			removeGroup(i, j, m, state)
		}
	}
}

func ProcessBoard(state *board.Board, bTurn bool) {
	var dim = len(*state)
	// we count the opponent's liberties and remove dead groups before we count
	// our own
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if bTurn {
				removeGroupIfDead(i, j, board.WHITE, board.MARKED_W, state)
			} else {
				removeGroupIfDead(i, j, board.BLACK, board.MARKED_B, state)
			}
		}
	}
	cleanBoard(state)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if bTurn {
				removeGroupIfDead(i, j, board.BLACK, board.MARKED_B, state)
			} else {
				removeGroupIfDead(i, j, board.WHITE, board.MARKED_W, state)
			}
		}
	}
	cleanBoard(state)
}
