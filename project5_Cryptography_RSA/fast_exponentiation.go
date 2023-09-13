package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// Use fast exponentiation to calculate num ^ pow.
func fastExp(num *big.Int, pow *big.Int) (c *big.Int) {
	var r *big.Int
	result := big.NewInt(1)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	num_, _ := new(big.Int).SetString(num.String(), 10)
	pow_, _ := new(big.Int).SetString(pow.String(), 10)
	for {
		if pow_.Cmp(zero) == -1 || pow_.Cmp(zero) == 0 {
			break
		}
		// if pow % 2 == 1
		_, r = new(big.Int).DivMod(pow_, two, new(big.Int))
		if r.Cmp(one) == 0 {
			// result *= num
			result = result.Mul(result, num_)
		}
		// pow /= 2
		pow_ = pow_.Div(pow_, two)
		// num *= num
		num_ = num_.Mul(num_, num_)
	}
	return result
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

func sacanNumbers(s string) (*big.Int, *big.Int, *big.Int) {
	var n []*big.Int
	if len(strings.Fields(s)) < 3 {
		panic("3 arguments are required")
	}
	for _, f := range strings.Fields(s) {
		i, err := new(big.Int).SetString(f, 10)
		if err == true {
			n = append(n, i)
		}
	}
	return n[0], n[1], n[2]
}

func main() {
	num := big.NewInt(1)
	pow := big.NewInt(1)
	mod := big.NewInt(1)
	zero := big.NewInt(0)
	for {
		fmt.Printf("Enter num, pow and mod: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num, pow, mod = sacanNumbers(scanner.Text())
		}
		// if num or pow is zero, exit
		if num.Cmp(zero) == -1 || pow.Cmp(zero) == -1 || mod.Cmp(zero) == -1 {
			break
		}

		fmt.Printf(" num: %s \n pow: %s \n mod: %s \n (num^pow): %s \n (num^pow mod): %s \n", num.String(), pow.String(), mod.String(), fastExp(num, pow).String(), fastExpMod(num, pow, mod).String())

	}
}
