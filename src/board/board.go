package board

import (
	"fmt"
	"reflect"
	"strconv"
)

//board constants are wrapped ints
type Piece int

const (
	NULLS    = -1 //not a valid piece
	EMPTY    = 0  //empty intersection
	WHITE    = 1
	BLACK    = 2
	MARKED_B = 3
	MARKED_W = 4
	MARKED_E = 5
)

// output unicode constants
type Tile rune

const (
	//CORNERS:
	UL = '\u250c'
	LL = '\u2514'
	UR = '\u2510'
	LR = '\u2518'

	//SIDES:
	R   = '\u2524'
	L   = '\u251c'
	T   = '\u252c'
	B   = '\u2534'
	Mid = '\u253c'

	//pieces
	WHT = '\u25cb'
	BLK = '\u25c9'
)

//Alias to avoid passing around explicit arrays
type Board [][]Piece

func InitBoard(dim int) Board {
	board := make([][]Piece, dim)
	for i := range board {
		board[i] = make([]Piece, dim)
	}
	for a := 0; a < dim; a++ {
		for b := 0; b < dim; b++ {
			board[a][b] = EMPTY
		}
	}
	return board
}

func CloneBoard(src Board) Board {
	dest := make([][]Piece, len(src))
	for i := range src {
		dest[i] = make([]Piece, len(src[i]))
		copy(dest[i], src[i])
	}
	return dest
}

func EqualBoard(src Board, dest Board) bool {
	return reflect.DeepEqual(src, dest)
}

func PrintStdBoard(board Board) {

	var str string = "   " //three spaces
	var dim = len(board)
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S"}

	for a := 0; a < dim; a++ {
		str += alphabet[a] + " "
	}

	fmt.Println(str)
	str = ""

	for a := 0; a < dim; a++ {
		if a < 10 {
			str += strconv.Itoa(a) + "  " //two spaces
		} else {
			str += strconv.Itoa(a) + " "
		}
		for b := 0; b < dim; b++ {
			var tile rune = pieceToRune(board[a][b], a, b, dim)
			var straw, _ = strconv.Unquote(strconv.QuoteRune(tile))
			str += straw
			if b < dim-1 {
				str += "\u2500" //horizontal extenders
			}

		}
		fmt.Println(str)
		str = ""
	}

}

func pieceToRune(piece Piece, a int, b int, dim int) rune {
	var r rune
	if piece == EMPTY {
		switch {
		case a == b && b == 0:
			{
				r = '\u250c' //Upper Left
			}
		case a == 0 && b == dim-1:
			{
				r = '\u2510' //Upper Right
			}
		case a == dim-1 && b == 0:
			{
				r = '\u2514' //Lower Left
			}
		case a == b && b == dim-1:
			{
				r = '\u2518' //Lower Right
			}
		case a == 0:
			{
				r = '\u252c' //Top
			}
		case a == dim-1:
			{
				r = '\u2534' //Bottom
			}
		case b == 0:
			{
				r = '\u251c' //Left
			}
		case b == dim-1:
			{
				r = '\u2524' //Right
			}
		default:
			{
				r = '\u253c' //Middle Tile
			}
		}
	} else if piece == WHITE {
		r = '\u25cb'
	} else if piece == BLACK {
		r = '\u25c9'
	} else if piece == MARKED_W {
		r = 'W'
	} else if piece == MARKED_B {
		r = 'B'
	}
	//fmt.Println(piece,strconv.QuoteRune(r))
	return r
}
