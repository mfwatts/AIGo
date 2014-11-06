//package boardState
package main

import(
	"board"
)


func updateInfluence(board [][]int, x,y int) board [][]int{
	int dim = len(board)
		
	neighbors := []int{80,60,30,1}
	board[x][y] = null;

	copX := x+1;
	x = x-1;
	storeX := x

	for b, index:= y-4 ,3; b <dim; b, index = y+1,abs(index-b-1) {
		for a:=index ; copX < dim && a < 4; a, copX = a+1, copX+1{
			if(board[copX][y] ==0){
				board[copX][y] = neighbors[a]
			}else if board[copX][y]!=null { 
				board[copX][y] = max(board[copX][y], neighbors[a])+5;
			}
		}//positive for
	
		for a:=index ; x>=0 && a<4; a, x = a+1, x-1{
			if(board[x][y] ==0){
				board[x][y] = neighbors[a]
			}else if board[x][y]!=null { 
				board[x][y] = max(board[x][y], neighbors[a])+5;
			}			
		}//negative for
		copX, x = storeX
		//if index = 
	}//overarching 4
	return board
}

func main(){
	dim = 9
	board := make([][]int, dim)
    for i := range board {
        board[i] = make([]int, dim)
    }
    for a:=0; a<dim; a++{
       for b:=0; b<dim;b++{
           board[a][b] = 0;
       }
    }
	board = updateInfluence(board, 5,5)
	
	
	str := ""

	for a:=0; a<dim; a++{
		for b:=0; b<dim; b++{
			str +=Itoa(board[a][b])
		}
		str+= "\n"
	}
	fmt.Println(str)
}

