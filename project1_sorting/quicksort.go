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

// Performs quicksort on a slice
func quicksort(slice []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	// Partition array and get the pivot index
	p := partition(slice, lo, hi)

	// Sort the two partitions
	quicksort(slice, lo, p-1) // Left side of pivot
	quicksort(slice, p+1, hi) // Right side of pivot
}

// Partitions the current array in left and right subarrays
func partition(slice []int, lo, hi int) int {
	pivot := slice[hi] // Choose the last element as the pivot
	// Temporary pivot index
	i := lo - 1
	for j := lo; j < hi; j++ {
		// If the current element is less than or equal to the pivot
		if slice[j] <= pivot {
			// Move the temporary pivot index forward
			i = i + 1
			// Swap the current element with the element at the temporary pivot index
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	// Move the pivot element to the correct pivot position (between the smaller and larger elements)
	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]
	return i // the pivot index
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

	// Sort and display the result.
	quicksort(slice, 0, len(slice)-1)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
