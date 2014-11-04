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
	g3 := initGame(3)

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
}

func Test_playMove(t *testing.T) {
	if false {
		t.Errorf("playMove failure %v", true)
	}
}
