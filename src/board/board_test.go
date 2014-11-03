package board

import (
  "testing"
)

func Test_CloneBoard(t *testing.T) {
  a := InitBoard(3)
  b := CloneBoard(a)
  a[0][0] = BLACK
  if b[0][0] == BLACK {
    t.Errorf("CloneBoard doesn't copy properly")
  }
}

func Test_CompareBoards(t *testing.T) {
  a := InitBoard(3)
  b := CloneBoard(a)
  a[0][1] = BLACK
  b[0][1] = BLACK
  if !EqualBoard(a,b) {
    t.Errorf("We can't easy compare two boards")
  }
}
