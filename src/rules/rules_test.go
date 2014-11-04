package rules

import (
	"board"
	"fmt"
	"testing"
)

func Test_removeGroupIfDead(t *testing.T) {
	/*
	   // We construct this test board
	   BWE
	   WEE
	   EEE
	*/
	var dim int = 3
	t_board := board.InitBoard(dim)

	t_board[0][0] = board.BLACK
	t_board[0][1] = board.WHITE
	t_board[1][0] = board.WHITE

	removeGroupIfDead(0, 0, board.BLACK, board.MARKED_B, &t_board)
	board.PrintStdBoard(t_board)
	if t_board[0][0] != board.EMPTY {
		t.Errorf("fail")
	}
}

func Test_ProcessBoard(t *testing.T) {
	/*
	   // We construct this test board
	   BWB
	   WBE
	   EEE
	   // We expect the following after ProcessBoard(&t_board_true,true)
	   BEB
	   WBE
	   EEE
	   // We expect the following after ProcessBoard(&t_board_false,false)
	   EWB
	   WBE
	   EEE
	*/
	var dim int = 3
	t_board := board.InitBoard(dim)

	t_board[0][0] = board.BLACK
	t_board[0][1] = board.WHITE
	t_board[0][2] = board.BLACK
	t_board[1][0] = board.WHITE
	t_board[1][1] = board.BLACK
	t_board[1][2] = board.EMPTY
	t_board[2][0] = board.EMPTY
	t_board[2][1] = board.EMPTY
	t_board[2][2] = board.EMPTY

	e_board_true := board.CloneBoard(t_board)
	e_board_true[0][1] = board.EMPTY

	e_board_false := board.CloneBoard(t_board)
	e_board_false[0][0] = board.EMPTY

	t_board_true := board.CloneBoard(t_board)
	t_board_false := board.CloneBoard(t_board)

	ProcessBoard(&t_board_true, true)
	ProcessBoard(&t_board_false, false)

	fmt.Println("t_board_true")
	for a := 0; a < dim; a++ {
		for b := 0; b < dim; b++ {
			fmt.Print(t_board_true[a][b])
			if t_board_true[a][b] != e_board_true[a][b] {
				t.Errorf("ProccessBoard fails at indices %v,%v", a, b)
			}
		}
		fmt.Println()
	}

	fmt.Println("t_board_false")
	for a := 0; a < dim; a++ {
		for b := 0; b < dim; b++ {
			fmt.Print(t_board_false[a][b])
			if t_board_false[a][b] != e_board_false[a][b] {
				t.Errorf("ProccessBoard fails at indices %v,%v", a, b)
			}
		}
		fmt.Println()
	}
}
