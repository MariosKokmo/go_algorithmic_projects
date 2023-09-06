package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Customer struct {
	id           string
	numPurchases int
}

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []Customer {
	s := make([]Customer, numItems)
	var initial string
	initial = "C"
	for i := 0; i < numItems; i++ {
		s[i].id = initial + strconv.Itoa(i)
		s[i].numPurchases = rand.Intn(max)
	}
	return s
}

// Print at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) > numItems {
		fmt.Println(slice[:numItems])
	} else {
		fmt.Println(slice)
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []Customer) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

// Countingsort algorithm
func countingSort(slice []Customer, max int) []Customer {
	count := make([]int, max)
	result := make([]Customer, len(slice))
	for _, value := range slice {
		count[value.numPurchases] += 1
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		count[slice[i].numPurchases] -= 1
		result[count[slice[i].numPurchases]] = slice[i]
	}
	return result
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
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
