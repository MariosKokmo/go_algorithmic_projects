package main

import "fmt"

func iterativeFactorial(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= int64(n); i++ {
		result *= i
	}

	return result
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}
