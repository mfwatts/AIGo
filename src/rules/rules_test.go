package rules

import (
  "testing"
  "board"
  "fmt"
)

func TestRules(t *testing.T) {
  const succ bool = true //false
  if !succ {
    t.Errorf("This is an intentional failure caused by %v",succ)
  }
}

func TestSmallBoard(t *testing.T) {
  /*
   // We construct this test board
   BWE
   EWE
   EEE
   // We expect the following after ProcessBoard
   EWE
   WBE
   EEE
  */
  var dim int = 3
  //t_board is what we test on
  t_board := board.InitBoard(dim)
  //e_board is what we expect as result
  e_board := board.InitBoard(dim)

  t_board[0][0] = board.BLACK; t_board[0][1] = board.WHITE; t_board[0][2] = board.EMPTY;
  t_board[1][0] = board.WHITE; t_board[1][1] = board.BLACK; t_board[1][2] = board.EMPTY;
  t_board[2][0] = board.EMPTY; t_board[2][1] = board.EMPTY; t_board[2][2] = board.EMPTY;

  e_board[0][0] = board.EMPTY; e_board[0][1] = board.WHITE; e_board[0][2] = board.EMPTY;
  e_board[1][0] = board.WHITE; e_board[1][1] = board.BLACK; e_board[1][2] = board.EMPTY;
  e_board[2][0] = board.EMPTY; e_board[2][1] = board.EMPTY; e_board[2][2] = board.EMPTY;

  ProcessBoard(&t_board)
    for a:=0; a<dim; a++{
       for b:=0; b<dim;b++{
           fmt.Print(t_board[a][b])
           if t_board[a][b] != e_board[a][b]{
             t.Errorf("ProccessBoard fails at indices %v,%v",a,b)
           }
       }
       fmt.Println()
    }
}
