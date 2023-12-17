package main

import (
	"fmt"
	"math/big"
)

// factorizes as many numbers as possible into a product
// of two smaller numbers and prints the result
func print_prime_factors(number *big.Int, odd_prime *big.Int) int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	limit := big.NewInt(611953)

	// now I don't care what really happens, let us go until we can't
	if odd_prime.Cmp(limit) >= 0 {
		limit.Set(big.NewInt(30597650))
	}

	// Handle the case when the number is less than or equal to 1
	if number.Cmp(one) <= 0 {
		return 0
	}

	// check for cached values
	cached_result, found := result_cache[number.String()]
	if found {
		fmt.Printf("%s=%s\n", number.String(), cached_result)
		return 0
	}

	// Handle the case when the number is divisible by 2
	if new(big.Int).Mod(number, two).Cmp(zero) == 0 {
		quotient := new(big.Int)
		quotient.Quo(number, two)
		result_cache[number.String()] = quotient.String() + "*2"

		// print the result and exit from here
		fmt.Printf("%s=%s*2\n", number.String(), quotient.String())
		return 0
	}

	sqrt_number := new(big.Int).Set(number)
	sqrt_number.Sqrt(sqrt_number)

	loop_counter := 0
	for odd_prime.Cmp(sqrt_number) <= 0 {
		if new(big.Int).Mod(number, odd_prime).Cmp(zero) == 0 {
			quotient := new(big.Int)

			// this is an expensive operation so we'd save the result for later
			quotient.Quo(number, odd_prime)

			// let's save the result for later
			result_cache[number.String()] = quotient.String() + "*" +
				odd_prime.String()
			fmt.Printf("%s=%s*%s\n", number.String(), quotient.String(),
				odd_prime.String())
			return 0
		}
		// skip this number if we go past this prime number without a match
		if odd_prime.Cmp(limit) > 0 || loop_counter == 1000000 {
			// fmt.Printf("Limit reached for [%s]\n", number.String())
			return 1
		}
		odd_prime.Add(odd_prime, two) // odd_prime += 2
		loop_counter++
	}

	// the number is a prime
	fmt.Printf("%s=%s*1\n", number.String(), number.String())
	return 0
}
