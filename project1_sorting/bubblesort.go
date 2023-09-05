package main

import (
	"fmt"
	"math/rand"
)

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

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

// Use bubble sort to sort the slice.
func bubbleSort(slice []int) {
	n := len(slice)
	var temp int
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if slice[i-1] > slice[i] {
				temp = slice[i-1]
				slice[i-1] = slice[i]
				slice[i] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
