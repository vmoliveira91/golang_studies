package main

import "fmt"

func nextMove(currentMove string) string {
	var nextMove string
	
	switch currentMove {
		case "LeftToRight":
			nextMove = "TopToBottom"
		case "TopToBottom":
			nextMove = "RightToLeft"
		case "RightToLeft":
			nextMove = "BottomToTop"
		case "BottomToTop":
			nextMove = "LeftToRight"
	}
	
	return nextMove
}

func generateMatrix(n int) [][]int {
	// Initialize
	var matrix = make([][]int, n)    
	for row := range matrix {
		matrix[row] = make([]int, n)
	}
	numberOfSpotsToFill := n * n
	numberOfSpotsFilled := 0
	move := "LeftToRight"
	axisX := 0
	axisY := 0
	number := 1
	
	for numberOfSpotsFilled != numberOfSpotsToFill {
		if move == "LeftToRight"{
			for j := axisY; j < (n - axisY); j++ {
				matrix[axisX][j] = number
				number++
				numberOfSpotsFilled++
			}
			move = nextMove(move)
		} else if move == "TopToBottom"{
			for i:= axisX + 1; i < (n - axisX); i++ {
				matrix[i][n - axisY - 1] = number
				number++
				numberOfSpotsFilled++
			}
			axisY++
			move = nextMove(move)
		} else if move == "RightToLeft" {
			for j := n - axisY - 1; j >= axisY - 1 ; j-- {
				matrix[n - axisX - 1][j] = number
				number++
				numberOfSpotsFilled++
			}
			move = nextMove(move)
			axisX++
		} else {
			for  i := n - axisX - 1; i >= axisX ; i-- {
				matrix[i][axisY - 1] = number
				number++
				numberOfSpotsFilled++
			}
			move = nextMove(move)
		}
	}
	
	return matrix
}

func main() {
	fmt.Println(generateMatrix(5))
}