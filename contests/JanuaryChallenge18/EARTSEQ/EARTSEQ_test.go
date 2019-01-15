package main_test

import (
	"math/rand"
	"testing"
	"time"

	"."
)

func TestEARTSEQ(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	var _gcd int
	var r []int
	primes := main.Sieve(main.MAX)
	for n := 3; n < 50000; n++ {
		r = main.Solve(n, primes)
		for i := 0; i < n; i++ {
			_gcd = gcd(r[i], r[(i+1)%n])
			if r[i] < 1 || r[i] > 1000000000 || _gcd == 1 || gcd(_gcd, r[(i+2)%n]) > 1 {
				t.Errorf("Error return in n: %d", n)
			}
		}
	}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
