package main_test

import (
	"math/rand"
	"testing"

	"./"
)

func TestINTXOR(t *testing.T) {
	rand := rand.New(rand.NewSource(100))
	tc := 20000
	var a, res []int
	var n, i int
	for ; tc > 0; tc-- {
		n = rand.Intn(50000) + 8
		a = make([]int, n)
		for i = 0; i < n; i++ {
			a[i] = rand.Intn(1<<30) + 1
		}
		res = main.Solve(n, func(idx1 int, idx2 int, idx3 int) int {
			return a[idx1-1] ^ a[idx2-1] ^ a[idx3-1]
		})
		for i = 0; i < n; i++ {
			if a[i] != res[i] {
				break
			}
		}
		if i < n {
			t.Errorf("Wrong answer %d | %d", n, i)
		}
	}
}
