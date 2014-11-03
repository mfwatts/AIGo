package game

import (
  "fmt"
  "board"
  "os"
  "bufio"
  "strings"
  "strconv"
  "rules"
)

//TODO List
// 1. Finish implementing functions
// 2. Define moveStatus codes
// 3. Write good unit tests in twoPlayerGame_test.go

// Game is the simplist way to represent a Go game
type Game struct {
  prev board.Board //previous state (used to check for ko)
  curr board.Board //current state of board
  blackTurn bool //true if black's turn to play
}


//return true,false,0,0 if player resigned
//returns false,true,0,0 if player passed
// otherwise return false,false,<i>,<j> for move in position [i,j]
func getMove() (bool,bool,int,int) {
    reader := bufio.NewReader(os.Stdin)
    for { //loop until valid input is provided
    fmt.Println("Play a move? yes/no")
    text, _ := reader.ReadString('\n')
    text = strings.ToLower(strings.TrimSpace(text))
    //fmt.Println(text) //debug line
    if text == "yes" {
        fmt.Println("Enter pair row,column (e.g 3,5)")
        pairStr, _ := reader.ReadString('\n')
        pair := strings.Split(strings.TrimSpace(pairStr),",")
        //fmt.Println(pair) //debug line
        if len(pair) != 2{
            ; // TODO complain at user
        } else{
            i, _ := strconv.Atoi(strings.TrimSpace(pair[0]))
            j, _ := strconv.Atoi(strings.TrimSpace(pair[1]))
            //fmt.Println(i,j) //debug line
            return false,false,i,j
        }
    } else if text == "no"{
        fmt.Println("Enter 'yes' to pass")
        text, _ = reader.ReadString('\n')
        text = strings.ToLower(strings.TrimSpace(text))
        //fmt.Println(text) //debug line
        if text == "yes"{
            return false,true,0,0
        }

        fmt.Println("enter 'yes' to resign")
        text, _ = reader.ReadString('\n')
        text = strings.ToLower(strings.TrimSpace(text))
        //fmt.Println(text) //debug line
        if text == "yes"{
            return true,false,0,0
        }
    }
    fmt.Println("Unable to process input, please try again");
    }
    //return false,true,0,0
}

// returns true,0 is move is legal
// otherwise, return false,<reason>
// define int constants to provide more info
func (g *Game) isLegal(i int, j int) (bool,int) {
   // it's illegal to play on a non-empty point
   if g.curr[i][j] != board.EMPTY {
       return false,1
   }

   // preview state of board after playing this move
   tmp := board.CloneBoard(g.curr)
   if g.blackTurn{
       tmp[i][j] = board.BLACK
   } else{
       tmp[i][j] = board.WHITE
   }
   rules.ProcessBoard(&tmp)

  // it's illegal to commit suicide
  if tmp[i][j] == board.EMPTY{
      return false,2
  }

  // it's illegal to repeat state, check if ko
  if board.EqualBoard(tmp, g.prev){
      return false,3
  }

  // if we don't violate previous rules, it's legal
  return true,0
}

// pre: (i,j) is a legal move
// post: move has been played, state of game is fully upto date
func (g *Game) playMove(i int, j int) {
  //copy current to previous state
  //place stone at i,j on current
  //process board
   tmp := board.CloneBoard(g.curr)
   if g.blackTurn{
       tmp[i][j] = board.BLACK
   } else{
       tmp[i][j] = board.WHITE
   }
   rules.ProcessBoard(&tmp)
   g.prev = g.curr
   g.curr = tmp
}

//return empty game of dimensions dim*dim
func initGame(dim int) Game {
  return Game{prev : board.InitBoard(dim),curr: board.InitBoard(dim),
         blackTurn: true}
}

// plays game between two players
func PlayGame(dim int) {
    // STRUCTURE
    // while both player's haven't passed
    //   print current state of board
    //   get move from correct player
    //   while move is illegal, ask for new move
    //   play it, update state
    //   switch player's turn
    g := initGame(dim) //full game state
    wPassed := false //white passed on last move
    bPassed := false //black passed on last move
    bTurn := true // is it black's turn?
    resign, pass, i, j := false, true, 0,0 //the current move
    legalMove := false
    moveStatus := 0 // constants for why a move is legal/illegal
    for !wPassed && !bPassed {

        board.PrintStdBoard(g.curr)
        if bTurn{ fmt.Println("Black's turn")
        } else{ fmt.Println("White's turn") }

        for !legalMove {
            resign,pass,i,j = getMove()
            if resign {
              if bTurn{ fmt.Println("Black resigned, white wins")
              } else { fmt.Println("White resigned, black wins") }
              return //game over, escape
            } else if pass {
                if bTurn{ bPassed = true
                } else { wPassed = true }
                legalMove=true //turn over, next player's turn now
            } else{ //we check if move (i,j) is legal
                if bTurn{ bPassed = false
                } else { wPassed = false }
                //if legal, play move
                legalMove, moveStatus = g.isLegal(i,j)
                if legalMove { g.playMove(i,j)
                } else{ //otherwise tell user why move is illegal
                    switch{
                        case moveStatus == 1 :{ }
                    }
                }
            }
        }

        //reset, update any state for next turn
        legalMove = false
        bTurn = !bTurn //switch players
        g.blackTurn = bTurn //keep game-state current
    }
}

