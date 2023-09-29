package main

import (
	"fmt"
	"time"
)

func findFactors(num int) []int {
	factors := []int{}
	for {
		if num%2 == 0 {
			factors = append(factors, 2)
			num /= 2
		} else {
			break
		}
	}
	factor := 3
	for {
		if factor > num {
			break
		}
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		} else {
			factor += 2
		}
	}
	return factors
}

// Finds the factors using a predefined sieve
func findFactorsSieve(num int) []int {
	factors := []int{}
	i := 0
	factor := primes[i]
	for {
		if factor > num {
			break
		}
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		} else {
			i += 1
			factor = primes[i]
		}
	}
	return factors
}

func multiplySlice(factors []int) int {
	number := 1
	for _, factor := range factors {
		number *= factor
	}
	return number
}

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

var primes []int

func main() {
	// create a large sieve
	bools := eulersSieve(20_000_000)
	primes = sieveToPrimes(bools)
	for {
		var num int
		fmt.Printf("Number: ")
		fmt.Scan(&num)
		if num < 2 {
			break
		}
		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(num)
		elapsed := time.Since(start)
		fmt.Printf("findFactors:       %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(num)
		elapsed = time.Since(start)
		fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()
	}
}
