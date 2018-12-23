package main_test

import (
	"fmt"
	"math/rand"
	"testing"

	"./"
)

func TestINTXOR(t *testing.T) {
	rand := rand.New(rand.NewSource(100))
	tc := 20000
	var cnt, a, res []int
	var n, i int
	for ; tc > 0; tc-- {
		n = rand.Intn(50000) + 8
		a = make([]int, n)
		for i = 0; i < n; i++ {
			a[i] = rand.Intn(1<<31) + 1
		}
		cnt = make([]int, n)
		res = main.Solve(n, func(idx1 int, idx2 int, idx3 int) int {
			cnt[idx1-1]++
			cnt[idx2-1]++
			cnt[idx3-1]++
			return a[idx1-1] ^ a[idx2-1] ^ a[idx3-1]
		})
		for i = 0; i < n; i++ {
			if cnt[i] > 3 {
				fmt.Printf("%d %d\n", i, cnt[i])
				break
			}
		}
		if i < n {
			t.Errorf("Wrong answer, index used more than 3 times n: %d | i: %d", n, i)
		} else {
			for i = 0; i < n; i++ {
				if a[i] != res[i] {
					break
				}
			}
			if i < n {
				t.Errorf("Wrong answer n: %d | i: %d", n, i)
			}
		}
	}
}
