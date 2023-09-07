package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	L := 0
	R := len(slice) - 1
	tests := 0
	for {
		if L >= R {
			break
		}
		m := (L + R) / 2
		tests += 1
		if slice[m] < target {
			L = m + 1
		} else if slice[m] > target {
			R = m - 1
		} else {
			return m, tests
		}
	}
	return -1, tests
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
	quicksort(slice, 0, len(slice)-1)
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
		index, tests := binarySearch(slice, target)
		if index == -1 {
			fmt.Printf("Target %v not found, %v tests\n", target, tests)
		} else {
			fmt.Printf("values[%v] = %v, %v tests\n", index, target, tests)
		}
	}
}
