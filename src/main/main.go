package main

import (
  //"fmt" //currently unused
  "board"
  "rules"
  "game"
)



func main() {
    var dim int =9
    var m_board board.Board = board.InitBoard(dim)
    rules.ProcessBoard(&m_board)
    board.PrintStdBoard(m_board)

    // remove this when PlayGame becomes fully implemented
    if false{
        game.PlayGame(dim)
    }
}
