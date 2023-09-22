// Finds ONE solution to the N queens problem
// Uses a brute force search so it is not very efficient
package main

import (
	"fmt"
	"time"
)

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
// Brute force approach checks all the squares
func placeQueens1(board [][]string, numRows, r, c int) bool {
	if r >= numRows {
		return boardIsASolution(board)
	}
	nextC := (c + 1) % numRows
	nextR := r + (c+1)/numRows
	// if we don't place a queen
	if placeQueens1(board, numRows, nextR, nextC) {
		return true
	}
	// if we place a queen
	board[r][c] = "Q"
	if placeQueens1(board, numRows, nextR, nextC) {
		return true
	}
	// revert to empty
	board[r][c] = "."
	return false

}

// Stops if N queens have been placed already
func placeQueens2(board [][]string, numRows, r, c, numPlaced int) bool {
	if r >= numRows {
		return boardIsASolution(board)
	}
	if numPlaced == numRows {
		return boardIsASolution(board)
	}
	nextC := (c + 1) % numRows
	nextR := r + (c+1)/numRows
	// if we don't place a queen
	if placeQueens2(board, numRows, nextR, nextC, numPlaced) {
		return true
	}
	// if we place a queen
	board[r][c] = "Q"
	curNumPlaced := numPlaced + 1
	if placeQueens2(board, numRows, nextR, nextC, curNumPlaced) {
		return true
	}
	// revert to empty
	board[r][c] = "."
	return false
}

// Try to place a queen in this column.
// Return true if we find a legal board.
// Implements a 1 queen per column improvement
func placeQueens4(board [][]string, numRows, c int) bool {
	if c == numRows {
		if boardIsLegal(board, numRows) {
			return true
		}
		return false
	}
	if c < numRows {
		if !boardIsLegal(board, numRows) {
			return false
		}
		// assign a queen to column c
		for r := 0; r < numRows; r++ {
			board[r][c] = "Q"
			if placeQueens4(board, numRows, c+1) {
				return true
			}
			board[r][c] = "."
		}
	}
	return false
}

func main() {
	//test_utilities()

	const numRows = 12
	board := makeBoard(numRows)

	start := time.Now()
	//success := placeQueens1(board, numRows, 0, 0)
	//success := placeQueens2(board, numRows, 0, 0, 0)
	//success := placeQueens3(board, numRows, 0, 0, 0)
	success := placeQueens4(board, numRows, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dumpBoard(board, numRows)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}

func test_utilities() {
	fmt.Println("=========")
	fmt.Println("==Test 1=")
	board := [][]string{{".", "Q", ".", "."},
		{".", ".", ".", "Q"},
		{"Q", ".", ".", "."},
		{".", ".", "Q", "."}}
	dumpBoard(board, 4)
	fmt.Println(seriesIsLegal(board, 4, 0, 0, 1, 1)) // checks diagonal
	fmt.Println(boardIsLegal(board, 4))              // should print true
	fmt.Println(boardIsASolution(board))             // true
	fmt.Println("========")
	fmt.Println("========")
	fmt.Println("==Test 2==")
	board = [][]string{{"Q", ".", ".", "."},
		{".", ".", ".", "Q"},
		{"Q", ".", ".", "."},
		{".", ".", "Q", "."}}
	dumpBoard(board, 4)
	fmt.Println(seriesIsLegal(board, 4, 0, 0, 1, 0)) // checks first column. gives false
	fmt.Println(boardIsLegal(board, 4))              // should print false
	fmt.Println(boardIsASolution(board))             // false
	fmt.Println("========")
	fmt.Println("==Test 3==")
	board = [][]string{{".", ".", "Q", "."},
		{".", ".", ".", "Q"},
		{"Q", ".", ".", "."},
		{".", ".", "Q", "."}}
	dumpBoard(board, 4)
	fmt.Println(seriesIsLegal(board, 4, 0, 0, 1, 1)) // gives true
	fmt.Println(seriesIsLegal(board, 4, 0, 2, 1, 1)) // checks first column. gives false. diagonal fails
	fmt.Println(boardIsLegal(board, 4))              // should print false
	fmt.Println(boardIsASolution(board))             // false
	fmt.Println("========")
	fmt.Println("==Test 4==")
	board = [][]string{{".", ".", ".", "Q"},
		{".", ".", "Q", "."},
		{".", "Q", ".", "."},
		{"Q", ".", ".", "."}}
	dumpBoard(board, 4)
	fmt.Println(seriesIsLegal(board, 4, 0, 3, 1, 0))   // checks last column. gives true
	fmt.Println(seriesIsLegal(board, 4, 0, 0, 1, 1))   // true
	fmt.Println(seriesIsLegal(board, 4, 3, 3, -1, -1)) // . gives false. diagonal fails
	fmt.Println(boardIsLegal(board, 4))                // should print false
	fmt.Println(boardIsASolution(board))               // false
	fmt.Println("========")
}

func makeBoard(numRows int) [][]string {
	board := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		board[i] = make([]string, numRows)
		for j := 0; j < numRows; j++ {
			board[i][j] = "."
		}
	}
	return board
}

// Return true if this series of squares contains at most one queen.
func seriesIsLegal(board [][]string, numRows, r0, c0, dr, dc int) bool {
	r := r0
	c := c0
	numQueens := 0
	for {
		if r >= numRows || c >= numRows || r < 0 || c < 0 {
			break
		}
		if board[r][c] == "Q" {
			numQueens += 1
		}
		if numQueens > 1 {
			return false
		}
		r = r + dr
		c = c + dc
	}
	return true
}

// Return true if the board is legal.
// checks every row column and diagonal
func boardIsLegal(board [][]string, numRows int) bool {
	// for every row start at (i,0) and increase the column
	for row := 0; row < numRows; row++ {
		isLegal := seriesIsLegal(board, numRows, row, 0, 0, 1)
		if !isLegal {
			return false
		}
	}
	// for every col start at (0,col) and increase the row
	for col := 0; col < numRows; col++ {
		isLegal := seriesIsLegal(board, numRows, 0, col, 1, 0)
		if !isLegal {
			return false
		}
	}
	// Check right diagonals
	// for every diagonal start at (0, col) and increase the column and the row
	for col := 0; col < numRows; col++ {
		isLegal := seriesIsLegal(board, numRows, 0, col, 1, 1)
		if !isLegal {
			return false
		}
	}
	// for every diagonal start at (row, 0) and increase the column and the row
	for row := 0; row < numRows; row++ {
		isLegal := seriesIsLegal(board, numRows, row, 0, 1, 1)
		if !isLegal {
			return false
		}
	}
	// Check the left diagonals starting at top row
	for col := numRows - 1; col >= 0; col-- {
		isLegal := seriesIsLegal(board, numRows, 0, col, 1, -1)
		if !isLegal {
			return false
		}
	}
	for row := 0; row < numRows; row++ {
		isLegal := seriesIsLegal(board, numRows, row, numRows-1, 1, -1)
		if !isLegal {
			return false
		}
	}
	return true
}

// Return true if the board is legal and a solution.
func boardIsASolution(board [][]string) bool {
	if !boardIsLegal(board, len(board)) {
		return false
	}
	// Count the queens
	queens := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == "Q" {
				queens += 1
			}
		}
	}
	if queens == len(board) {
		return true
	}
	return false
}

// Displays the board
func dumpBoard(board [][]string, numRows int) {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numRows; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println("")
	}
}
