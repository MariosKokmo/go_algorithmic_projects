// Implementation of quicksort using additional O(n) space
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
func quicksort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	left_partition, right_partition, pivot := partition(slice)
	result := append(append(quicksort(left_partition), pivot), quicksort(right_partition)...)
	return result
}

// Partitions the current array in
func partition(slice []int) ([]int, []int, int) {
	// get pivot
	pivot := slice[0]
	var left_partition []int
	var right_partition []int
	for i := 1; i < len(slice); i++ {
		if slice[i] <= pivot {
			left_partition = append(left_partition, slice[i])
		} else {
			right_partition = append(right_partition, slice[i])
		}
	}
	return left_partition, right_partition, pivot
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
	result := quicksort(slice)
	printSlice(result, 40)

	// Verify that it's sorted.
	checkSorted(result)
}
