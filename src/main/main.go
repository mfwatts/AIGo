package main

import (
  //"fmt" //currently unused
  "board"
  "rules"
)



func main() {
    var dim int =9
	var m_board [][]board.Piece = board.InitBoard(dim)
	board.PrintStdBoard(m_board)
}
