package game

import (
	"board"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//return true,false,0,0 if player resigned
//returns false,true,0,0 if player passed
// otherwise return false,false,<i>,<j> for move in position [i,j]
func getMove() (bool, bool, int, int) {
	reader := bufio.NewReader(os.Stdin)
	for { //loop until valid input is provided
		fmt.Println("Play a move? yes/no")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.TrimSpace(text))
		//fmt.Println(text) //debug line
		if text == "yes" {
			fmt.Println("Enter pair row,column (e.g 3,5)")
			pairStr, _ := reader.ReadString('\n')
			pair := strings.Split(strings.TrimSpace(pairStr), ",")
			//fmt.Println(pair) //debug line
			if len(pair) != 2 {
				// TODO complain at user
			} else {
				i, _ := strconv.Atoi(strings.TrimSpace(pair[0]))
				j, _ := strconv.Atoi(strings.TrimSpace(pair[1]))
				//fmt.Println(i,j) //debug line
				return false, false, i, j
			}
		} else if text == "no" {
			fmt.Println("Enter 'yes' to pass")
			text, _ = reader.ReadString('\n')
			text = strings.ToLower(strings.TrimSpace(text))
			//fmt.Println(text) //debug line
			if text == "yes" {
				return false, true, 0, 0
			}

			fmt.Println("enter 'yes' to resign")
			text, _ = reader.ReadString('\n')
			text = strings.ToLower(strings.TrimSpace(text))
			//fmt.Println(text) //debug line
			if text == "yes" {
				return true, false, 0, 0
			}
		}
		fmt.Println("Unable to process input, please try again")
	}
	//return false,true,0,0
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
	g := initGame(dim)                      //full game state
	wPassed := false                        //white passed on last move
	bPassed := false                        //black passed on last move
	bTurn := true                           // is it black's turn?
	resign, pass, i, j := false, true, 0, 0 //the current move
	legalMove := false
	moveStatus := 0 // constants for why a move is legal/illegal
	for !wPassed || !bPassed {

		board.PrintStdBoard(g.curr)
		if bTurn {
			fmt.Println("Black's turn")
		} else {
			fmt.Println("White's turn")
		}

		for !legalMove {
			resign, pass, i, j = getMove()
			if resign {
				if bTurn {
					fmt.Println("Black resigned, white wins")
				} else {
					fmt.Println("White resigned, black wins")
				}
				return //game over, escape
			} else if pass {
				if bTurn {
					bPassed = true
				} else {
					wPassed = true
				}
				legalMove = true //turn over, next player's turn now
			} else { //we check if move (i,j) is legal
				if bTurn {
					bPassed = false
				} else {
					wPassed = false
				}
				//if legal, play move
				legalMove, moveStatus = g.isLegal(i, j)
				if legalMove {
					g.playMove(i, j)
				} else { //otherwise tell user why move is illegal
					switch {
					case moveStatus == 1:
						{
						}
					}
				}
			}
		}

		//reset, update any state for next turn
		legalMove = false
		bTurn = !bTurn      //switch players
		g.blackTurn = bTurn //keep game-state current
	}
}
