package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

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

func lcm(a *big.Int, b *big.Int) (c *big.Int) {
	var result *big.Int
	GCD := gcd(a, b)
	result, _ = new(big.Int).DivMod(b, GCD, new(big.Int))
	result = result.Mul(result, a)
	return result
}

func numbers(s string) (*big.Int, *big.Int) {
	var n []*big.Int
	if len(strings.Fields(s)) < 2 {
		panic("2 arguments are required")
	}
	for _, f := range strings.Fields(s) {
		i, err := new(big.Int).SetString(f, 10)
		if err == true {
			n = append(n, i)
		}
	}
	return n[0], n[1]
}

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	one := big.NewInt(1)
	for {
		fmt.Printf("Enter A and B: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			a, b = numbers(scanner.Text())
		}

		if a.Cmp(one) == -1 || b.Cmp(one) == -1 {
			break
		}

		fmt.Printf("A %s, B %s, GCD: %s, LCM: %s\n ", a.String(), b.String(), gcd(a, b).String(), lcm(a, b).String())
	}
}
