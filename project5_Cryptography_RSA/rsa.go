package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

var Reader io.Reader

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q *big.Int) *big.Int {
	one := big.NewInt(1)
	dummy1 := big.NewInt(1)
	dummy2 := big.NewInt(1)
	mul_dummy := big.NewInt(1)
	pm1 := dummy1.Sub(p, one)
	qm1 := dummy2.Sub(q, one)
	g := gcd(pm1, qm1)
	mul_dummy = mul_dummy.Mul(pm1, qm1)
	return dummy1.Div(mul_dummy, g)
}

// Return a pseudo random number in the range [min, max).
func randRange(min *big.Int, max *big.Int) *big.Int {
	interval := big.NewInt(0)
	interval = interval.Sub(max, min)
	v, _ := rand.Int(rand.Reader, interval)
	v = v.Add(v, min)
	return v
}

func gcd(a *big.Int, b *big.Int) (c *big.Int) {
	var aa, bb, r *big.Int
	zero := big.NewInt(0)
	if a.Cmp(zero) == -1 {
		aa.Neg(a)
	} else {
		aa = a
	}
	if b.Cmp(zero) == -1 {
		bb.Neg(b)
	} else {
		bb = b
	}
	if aa.Cmp(zero) == 0 {
		return bb
	}
	if bb.Cmp(zero) == 0 {
		return aa
	}
	_, r = new(big.Int).DivMod(aa, bb, new(big.Int))
	return gcd(bb, r)
}

// Pick a random exponent e in the range (2, λn)
// such that gcd(e, λn) = 1.
func randomExponent(λn *big.Int) *big.Int {
	one := big.NewInt(1)
	three := big.NewInt(3)
	for {
		exponent := randRange(three, λn)
		g := gcd(exponent, λn)
		if g.Cmp(one) == 0 {
			return exponent
		}
	}
}

// Calculates the inverse of a in the modulus n
func inverseMod(a, n *big.Int) *big.Int {
	t := big.NewInt(0)
	newt := big.NewInt(1)
	r := new(big.Int).Set(n)
	newr := new(big.Int).Set(a)

	zero := big.NewInt(0)
	one := big.NewInt(1)

	quotient := big.NewInt(1)
	dummy_quot := big.NewInt(1)
	dummy1 := big.NewInt(1)
	dummy2 := big.NewInt(1)

	for {
		if newr.Cmp(zero) == 0 {
			break
		}
		quotient = dummy_quot.Div(r, newr)
		d1 := dummy1.Mul(quotient, newt)
		d2 := dummy1.Sub(t, d1)
		t, newt = new(big.Int).Set(newt), new(big.Int).Set(d2)

		d3 := dummy2.Mul(quotient, newr)
		d4 := dummy2.Sub(r, d3)
		r, newr = new(big.Int).Set(newr), new(big.Int).Set(d4)
	}

	if r.Cmp(one) == 1 {
		return nil
	}
	if t.Cmp(zero) == -1 {
		t = t.Add(t, n)
	}
	return t
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
// It performs numTests number of tests by selecting a random
// integer every time. It performs the modulo exponentiation.
// If it doesn't result in 1, it is not a prime and returns false.
// If it passes the test then the probability of it not being a prime
// is less than (1/2)^numTests
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
// randomly selects a number and runs the primality test.
func findPrime(min *big.Int, max *big.Int, numTests int) *big.Int {
	zero := big.NewInt(0)
	two := big.NewInt(2)
	for {
		number := randRange(min, max)
		mod := big.NewInt(0)
		if mod.Mod(number, two).Cmp(zero) == 0 {
			continue
		}
		isPrime := isProbablyPrime(number, numTests)
		if isPrime {
			return number
		}
	}
}

func main() {
	max := big.NewInt(1)
	min := big.NewInt(1)
	min, _ = min.SetString("1000000", 10)
	max, _ = max.SetString("10000000", 10)
	n := big.NewInt(1)
	two := big.NewInt(2)

	p := findPrime(min, max, 20)
	q := findPrime(min, max, 20)
	n = n.Mul(p, q)
	ln := totient(p, q)
	exp := randomExponent(ln)
	d := inverseMod(exp, ln)

	fmt.Println()
	fmt.Println("*** Public ***")
	fmt.Println("Public key modulus: ", n)
	fmt.Println("Public key exponent: ", exp)
	fmt.Println()
	fmt.Println("*** Private ***")
	fmt.Println("Primes: ", p, q)
	fmt.Println("λ(n): ", ln)
	fmt.Println("d: ", d)
	fmt.Println()
	for {
		var numStr string
		number := new(big.Int)
		fmt.Printf("Message less than %v (give 1 to stop): ", n)
		fmt.Scan(&numStr)
		number, _ = number.SetString(numStr, 10)
		if number.Cmp(two) == -1 {
			break
		}
		encrypt := fastExpMod(number, exp, n)
		fmt.Println("Message: ", number)
		fmt.Println("Ciphertext: ", encrypt)
		decrypt := fastExpMod(encrypt, d, n)
		fmt.Println("Plaintext: ", decrypt)
		fmt.Println()
	}
}
