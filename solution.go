package main

import "fmt"

func main() {
	var board [8]int    //temporary board
	var result [][8]int //all_results
	index := 0          //depth of the tree
	position := 0       //position of the queens
	resindex := 0       //index of the result
	fmt.Println(Solution(board, result, resindex, index, position))
	//fmt.Println(Check(board, 1, 1))
}

func Solution(board [8]int, result [][8]int, resindex, index, position int) [][8]int {
	if position > 7 {
		if index == 0 {
			return result
		} else {
			return Solution(board, result, resindex, index-1, board[index-1]+1)
		}
	} else if index == 7 {
		if Check(board, index, position) == false {
			return Solution(board, result, resindex, index, position+1)
		}
		board[index] = position
		result = append(result, board)
		return Solution(board, result, resindex+1, index, position+1)
	} else {
		if Check(board, index, position) == false {
			return Solution(board, result, resindex, index, position+1)
		}
		board[index] = position
		return Solution(board, result, resindex, index+1, 0)
	}
}

func Check(board [8]int, index, position int) bool {
	for i, v := range board[:index] {
		if (position == v) || (position == v-(index-i)) || (position == v+(index-i)) {
			return false
		}
	}
	return true
}
