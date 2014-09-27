package main

import "fmt"
import "strconv"

func main() {
    var dim int =9
	var board [][]rune = boardSetup(dim)

	printBoard(board)
	
}

/*
CORNERS:
UL 250c
LL 2514
UR 2510
LR 2518

SIDES: 
R 2524
L 251c
T 252c
B - 2534
Mid - 253c
WHT - 25cb
BLK - 25c9
*/
func boardSetup(dim int) [][]rune{


	board := make([][]rune, dim)
    for i := range board {
        board[i] = make([]rune, dim)
    }
    
   	for a:=0; a<dim; a++{
		for b:=0; b<dim;b++{
			
			switch{
				case a == b && b == 0: {
					board[a][b] ='\u250c' //Upper Left
				}
				case a == 0 && b == dim-1:{
					board[a][b] = '\u2510' //Upper Right
				}
				case a == dim-1 && b == 0:{
					board[a][b] = '\u2514' //Lower Left
				}
				case a == b && b == dim-1:{
					board[a][b] = '\u2518' //Lower Right
				}
				case a == 0:{
					board[a][b] = '\u252c' //Top
				}
				case a == dim-1:{
					board[a][b] = '\u2534' //Bottom
				}
				case b == 0:{
					board[a][b] = '\u251c' //Left
				}
				case b == dim-1:{
					board[a][b] = '\u2524' //Right
				}
				default:{
					board[a][b] = '\u253c' //Middle Tile
				}
			}
		}
	}
	return board

}

func printBoard(board [][]rune){
	var str string = "   "//three spaces
	var dim  = len(board)
	alphabet :=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S"}
	
	for a:= 0; a<dim; a++{
		str += alphabet[a] +" "
		}

	fmt.Println(str)
	str = ""
	
	for a:=0; a<dim; a++{
		if(a<10){
			str += strconv.Itoa(a) +"  "//two spaces
		}else {
			str += strconv.Itoa(a) +" ";
		}
		
		for b:=0; b<dim; b++{
			var straw, _ = strconv.Unquote(strconv.QuoteRune(board[a][b]))
			str +=straw;
			if b < dim-1{
				str+= "\u2500"//horizontal extenders
			}
			
	}
	fmt.Println(str);
	str = ""
	}
	
}
