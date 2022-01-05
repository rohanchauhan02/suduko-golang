package main

import "fmt"
import "os"

func display(matrix [][]int){
	for i:=0;i<len(matrix);i++{
		if i!=0 && i%3==0{
			fmt.Println("---------------------")
		}
		for j:=0;j<len(matrix[0]);j++{
			if j!=0 &&  j%3==0{
				fmt.Print("|")
			}
			fmt.Printf("%v%s",matrix[i][j]," ")
		}
		fmt.Println()
	}
}
func isSafePlace(matrix [][]int,row,col, data int) bool{
	for i:=0;i<9;i++{
		if matrix[i][col]==data{
			return false
		}
		if matrix[row][col]==data{
			return false
		}
	}
	nrow:=(row/3)*3
	ncol:=(col/3)*3
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if matrix[nrow+i][ncol+j]==data{
				return false
			}
		}
	}
	return true
}

func suduko(matrix [][]int,idx int){
	if idx==81{
		display(matrix)
		os.Exit(200)
		return  //why return is not working here??
	}
	
	row:=idx/9
	col:=idx%9
	if(matrix[row][col]==0){
		for i:=1;i<=9;i++{
			if isSafePlace(matrix,row,col,i){
				matrix[row][col]=i
				suduko(matrix,idx+1)
				matrix[row][col]=0
			}
		}
	}else{
		suduko(matrix,idx+1)
	}
}
func main(){
	// var matrix [][]int
	//The Worlds hardest Sudoku Puzzle
	matrix:=[][]int{{1,0,0,0,0,7,0,9,0},
	{0,3,0,0,2,0,0,0,8},
	{0,0,9,6,0,0,5,0,0},
	{0,0,5,3,0,0,9,0,0},
	{0,1,0,0,8,0,0,0,2},
	{6,0,0,0,0,4,0,0,0},
	{3,0,0,0,0,0,0,1,0},
	{0,4,0,0,0,0,0,0,7},
	{0,0,7,0,0,0,3,0,0}}
	display(matrix)
	fmt.Println("OutPut =>")
	suduko(matrix,0)
	// display(matrix)
}

