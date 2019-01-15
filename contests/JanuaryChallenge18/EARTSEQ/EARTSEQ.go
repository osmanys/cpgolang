package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func scan(a ...interface{})             { fmt.Fscan(reader, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func print(a ...interface{})            { fmt.Fprint(writer, a...) }

// MAX var
var MAX = 700000

func main() {
	defer writer.Flush()

	primes := Sieve(MAX)
	var n, t int
	var r []int
	scanf("%d\n", &t)
	for ; t > 0; t-- {
		scanf("%d\n", &n)
		r = Solve(n, primes)
		for i := 0; i < len(r); i++ {
			printf("%d ", r[i])
		}
		printf("\n")
	}
}

// Sieve func
func Sieve(n int) []int {
	isprime := make([]bool, n)
	prime := make([]int, 0)
	SPF := make([]int, n)
	for i := 2; i < n; i++ {
		isprime[i] = true
	}
	for i := 2; i < n; i++ {
		if isprime[i] {
			prime = append(prime, i)
			SPF[i] = i
		}

		for j := 0; j < len(prime) &&
			i*prime[j] < n && prime[j] <= SPF[i]; j++ {
			isprime[i*prime[j]] = false
			SPF[i*prime[j]] = prime[j]
		}
	}
	return prime
}

// Solve func
func Solve(n int, primes []int) []int {
	r := make([]int, n)
	r[0] = 2
	r[1] = 2
	for i := 1; i < n-1; i++ {
		if i%3 == 0 {
			r[i] *= 7
			r[i+1] = 7
		} else if i%3 == 1 {
			r[i] *= 3
			r[i+1] = 3
		} else {
			r[i] *= 5
			r[i+1] = 5
		}
	}
	r[n-1] *= 11
	r[0] *= 11
	for i := 0; i < n; i++ {
		r[i] *= primes[i+5]
	}
	return r
}
