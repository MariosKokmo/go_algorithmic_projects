// This implementation does not use the allowed moves function
package main

import (
	"fmt"
	"strconv"
	"time"
)

// The board dimensions.
//const numRows = 5
//const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var moveOffsets []Offset

var numCalls int64

func initializeOffsets() []Offset {
	moveOffsets = []Offset{Offset{-2, -1}, Offset{-1, -2},
		Offset{-2, 1}, Offset{-1, 2},
		Offset{2, -1}, Offset{1, -2},
		Offset{2, 1}, Offset{1, 2}}
	return moveOffsets
}

// Creates the 2D board. If board[2][3] is 4, then that means the knight visited square [2][3] in move 4
func makeBoard(numRows, numCols int) [][]int {
	board := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		board[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			board[i][j] = unvisited
		}
	}
	return board
}

// Displays the board
func dumpBoard(board [][]int, numRows, numCols int) {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			fmt.Printf("%02d ", board[i][j])
		}
		fmt.Println("")
	}
}

// Try to extend a knight's tour starting at (curRow, curCol).
// Return true or false to indicate whether we have found a solution.
// numVisited is the number of the squares the knight has visited in the current tour including the current square
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {
	numCalls += 1
	// If the night has visited all squares
	if numVisited == numRows*numCols {
		if requireClosedTour == false {
			// route is closed and valid
			return true
		} else {
			// we check if we can land back on the initial square
			for _, offset := range moveOffsets {
				r := curRow + offset.dr
				c := curCol + offset.dc

				if r < 0 || r >= numRows {
					continue
				}
				if c < 0 || c >= numCols {
					continue
				}

				if board[r][c] == 0 {
					return true
				}
			}
			// no move would lead us to the starting square
			return false
		}
	}
	// knight has not visited every square
	// check all posible moves from current square
	for _, offset := range moveOffsets {
		r := curRow + offset.dr
		c := curCol + offset.dc

		if r < 0 || r >= numRows {
			continue
		}
		if c < 0 || c >= numCols {
			continue
		}
		// check if it lands on previously visited or outside the board
		if board[r][c] != unvisited {
			continue
		}
		// else search recursively for a tour
		board[r][c] = numVisited
		found := findTour(board, numRows, numCols, r, c, numVisited+1)
		if found {
			return true
		}
		board[r][c] = unvisited
	}
	return false
}

func main() {
	numCalls = 0

	// Initialize the move offsets.
	initializeOffsets()
	var sizeN string
	fmt.Printf("Give size N: ")
	fmt.Scanln(&sizeN)
	numRows, _ := strconv.Atoi(sizeN)
	numCols := numRows

	// Create the blank board.
	board := makeBoard(numRows, numCols)

	// Try to find a tour.
	start := time.Now()
	if findTour(board, numRows, numCols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dumpBoard(board, numRows, numCols)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", numCalls)
}
