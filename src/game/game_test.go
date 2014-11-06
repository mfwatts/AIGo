package game

import (
	"testing"
	//"fmt"
	"board"
)

func Test_isLegal(t *testing.T) {
	//we test for these:
	//	playing in non-empty position
	//	suicide
	//	ko //unimplemented
	g3 := initGame(4)

	// Initial board
	// EWE
	// WEE
	// EEE
	(g3.curr)[0][1] = board.WHITE
	(g3.curr)[1][0] = board.WHITE

	leg, code := g3.isLegal(0, 1)
	if leg == true {
		t.Errorf("We can't play in non-empty point. 1=%v", code)
	}

	leg, code = g3.isLegal(0, 0)
	if leg == true {
		t.Errorf("We can't allow suicide. 2=%v", code)
	}

	// Initial board - ko
	// EWBE
	// WEWB
	// EWBE
	(g3.curr)[1][1] = board.EMPTY
	(g3.curr)[2][2], (g3.curr)[1][3], (g3.curr)[0][2] = board.BLACK, board.BLACK, board.BLACK
	(g3.curr)[1][2], (g3.curr)[2][1] = board.WHITE, board.WHITE

	board.PrintStdBoard(g3.curr)
	leg, code = g3.isLegal(1, 1)
	if leg == true {
		leg2, code2 := g3.isLegal(1, 2)
		if leg2 == true {
			t.Errorf("We can't allow ko. 3=%v", code2)
		}
	} else {
		t.Errorf("We don't allow ko!. code=%v", code)
	}

}

func Test_playMove(t *testing.T) {
	if false {
		t.Errorf("playMove failure %v", true)
	}
}

