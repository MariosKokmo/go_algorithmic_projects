package main

import (
	"fmt"
	"strconv"
)

func fibonacciOnTheFly(n int64) int64 {
	var value int64
	// if the fib(n) hasn't been calculated before
	if n > int64(len(fibonacciValues)-1) {
		if n == 1 {
			value = int64(n) + fibonacciOnTheFly(n-1)
			fibonacciValues = append(fibonacciValues, value)
		}
		if n == 0 {
			value = int64(n)
			fibonacciValues = append(fibonacciValues, value)
		}
		if n >= 2 {
			value = fibonacciOnTheFly(n-2) + fibonacciOnTheFly(n-1)
			fibonacciValues = append(fibonacciValues, value)
		}
	}
	return fibonacciValues[n]
}

func fibonacciBottomUp(n int64) int64 {
	if n <= 1 {
		return int64(n)
	}

	var fibI, fibIMinus1, fibIMinus2 int64
	fibIMinus2 = 0
	fibIMinus1 = 1
	fibI = fibIMinus1 + fibIMinus2
	for i := int64(1); i < n; i++ {
		// Calculate this value of fibI.
		fibI = fibIMinus1 + fibIMinus2

		// Set fibIMinus2 and fibIMinus1 for the next value.
		fibIMinus2 = fibIMinus1
		fibIMinus1 = fibI
	}
	return fibI
}

var fibonacciValues []int64

func main() {
	// Fill-on-the-fly.

	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)

		// Uncomment one of the following.
		fmt.Printf("fibonacciOnTheFly(%d) = %d\n", n, fibonacciOnTheFly(n))
		fmt.Printf("fibonacciBottomUp(%d)  = %d\n", n, fibonacciBottomUp(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacciValues); i++ {
		fmt.Printf("%d: %d\n", i, fibonacciValues[i])
	}
}
