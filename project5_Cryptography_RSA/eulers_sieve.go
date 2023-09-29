package main

import (
	"fmt"
	"math"
	"time"
)

// Build an Euler's sieve.
func eulersSieve(max int) []bool {
	sieve := make([]bool, max)
	for p := 3; p <= max; p = p + 2 {
		maxQ := int(max / p)
		if maxQ%2 == 0 {
			maxQ -= 1
		}
		for q := maxQ; q >= p; q = q - 2 {
			//if q is marked as prime in the table,
			//then cross out the entry for q * p to show it is not prime.
			if sieve[q] == false {
				sieve[p*q] = true
			}
		}
	}
	return sieve
}

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
	primes = append(primes, 2)
	for i := 3; i < len(sieve); i = i + 2 {
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
	sieveEuler := eulersSieve(max)
	elapsed := time.Since(start)
	fmt.Printf("Euler Elapsed: %f seconds\n", elapsed.Seconds())

	start = time.Now()
	sieveErat := sieveOfEratosthenes(max)
	elapsed = time.Since(start)
	fmt.Printf("Eratosthenes Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieveErat)
		primes := sieveToPrimes(sieveEuler)
		fmt.Println(primes)
	}
}
