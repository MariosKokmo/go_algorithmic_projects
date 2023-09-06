package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	for index, value := range slice {
		if value == target {
			return index, index + 1
		}
	}
	return -1, len(slice)
}

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) > numItems {
		fmt.Println(slice[:numItems])
	} else {
		fmt.Println(slice)
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	var reply string
	for {
		fmt.Printf("Target: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			reply = scanner.Text()
		}

		if reply == "" {
			fmt.Println("empty string given, terminating..")
			break
		}

		target, _ := strconv.Atoi(reply)
		fmt.Println("Target given: ", target)

		// call search
		index, tests := linearSearch(slice, target)
		if index == -1 {
			fmt.Printf("Target %v not found, %v tests\n", target, tests)
		} else {
			fmt.Printf("values[%v] = %v, %v tests\n", index, target, tests)
		}
	}
}
