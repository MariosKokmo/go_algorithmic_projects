package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
)

const numTests = 20

var Reader io.Reader

// Return a pseudo random number in the range [min, max).
func randRange(min *big.Int, max *big.Int) *big.Int {
	interval := big.NewInt(0)
	interval = interval.Sub(max, min)
	v, _ := rand.Int(rand.Reader, interval)
	v = v.Add(v, min)
	return v
}

// Use fast exponentiation to calculate num ^ pow mod modulo.
func fastExpMod(num *big.Int, pow *big.Int, mod *big.Int) (c *big.Int) {
	var r *big.Int
	result := big.NewInt(1)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	num_, _ := new(big.Int).SetString(num.String(), 10)
	pow_, _ := new(big.Int).SetString(pow.String(), 10)
	mod_, _ := new(big.Int).SetString(mod.String(), 10)
	for {
		if pow_.Cmp(zero) == -1 || pow_.Cmp(zero) == 0 {
			break
		}
		// if pow % 2 == 1
		_, r = new(big.Int).DivMod(pow_, two, new(big.Int))
		if r.Cmp(one) == 0 {
			result = result.Mul(result, num_)
			result = result.Mod(result, mod_)
		}
		// pow /= 2
		pow_ = pow_.Div(pow_, two)
		// num = (num * num) mod modulo
		// (num * num) mod modulo => (num mod modulo) * (num mod modulo) mod modulo
		_ = num_.Mul(num_, num_)
		_ = num_.Mod(num_, mod_)
	}
	return result
}

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p *big.Int, numTests int) bool {
	one := big.NewInt(1)
	pm1 := big.NewInt(0)
	pm1 = pm1.Sub(p, one)
	for i := 0; i < numTests; i++ {
		n := randRange(one, pm1)
		if fastExpMod(n, pm1, p).Cmp(one) != 0 {
			return false
		}
	}
	return true
}

// Probabilistically find a prime number within the range [min, max).
func findPrime(min *big.Int, max *big.Int, numTests int) *big.Int {
	zero := big.NewInt(0)
	two := big.NewInt(2)
	for {
		number := randRange(min, max)
		if number.Mod(number, two).Cmp(zero) == 0 {
			continue
		}
		isPrime := isProbablyPrime(number, numTests)
		if isPrime {
			return number
		}
	}
}

func testKnownValues() {
	primes := []*big.Int{
		big.NewInt(10009), big.NewInt(11113), big.NewInt(11699),
		big.NewInt(12809), big.NewInt(14149), big.NewInt(15643),
		big.NewInt(17107), big.NewInt(17881), big.NewInt(19301),
		big.NewInt(19793),
	}
	composites := []*big.Int{
		big.NewInt(10323), big.NewInt(11397), big.NewInt(12212),
		big.NewInt(13503), big.NewInt(14599),
		big.NewInt(16113), big.NewInt(17547), big.NewInt(17549),
		big.NewInt(18893), big.NewInt(19999),
	}
	fmt.Printf("Probability: %.7f\n", (1-math.Pow(0.5, numTests))*100)
	fmt.Println()
	fmt.Println("Primes:")
	for _, prime := range primes {
		if isProbablyPrime(prime, numTests) {
			fmt.Printf("%-10v Prime\n", prime)
		} else {
			fmt.Printf("%-10v Composite\n", prime)
		}
	}
	fmt.Println()
	fmt.Println("Composites:")
	for _, comp := range composites {
		if !isProbablyPrime(comp, numTests) {
			fmt.Printf("%-10v Composite\n", comp)
		} else {
			fmt.Printf("%-10v Prime\n", comp)
		}
	}
}

func main() {
	testKnownValues()
}
