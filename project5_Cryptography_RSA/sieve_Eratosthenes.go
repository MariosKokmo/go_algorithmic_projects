package main

import (
	"fmt"
	"math"
	"time"
)

// Build a sieve of Eratosthenes.
func sieveOfEratosthenes(max int) []bool {
	sieve := make([]bool, max)
	// we only need to check up to sqrt(n)
	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		if sieve[i] == false {
			for j := i * i; j < max; j = j + i {
				sieve[j] = true
			}
		}
	}
	return sieve
}

// Print out the primes in the sieve.
func printSieve(sieve []bool) {
	fmt.Print("2 ")
	for i := 3; i < len(sieve); i = i + 2 {
		if !sieve[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}

// Convert the sieve into a slice holding prime numbers.
func sieveToPrimes(sieve []bool) []int {
	var primes []int
	for i := 2; i < len(sieve); i++ {
		if sieve[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
