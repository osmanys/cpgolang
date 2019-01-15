package main_test

import (
	"math/rand"
	"testing"
	"time"

	"."
)

func TestDIFNEIGH(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	var a, b, c, d, n, m int
	var r [][]int
	var fail bool
	for _t := 0; _t < 100000; _t++ {
		n = rand.Intn(50) + 1
		m = rand.Intn(50) + 1
		_, r = main.Solve(n, m)
		fail = false
		for i := 0; i < n && !fail; i++ {
			for j := 0; j < m; j++ {
				a = -1
				b = -2
				c = -3
				d = -4
				if i > 0 {
					a = r[i-1][j]
				}
				if i < n-1 {
					b = r[i+1][j]
				}
				if j > 0 {
					c = r[i][j-1]
				}
				if j < m-1 {
					d = r[i][j+1]
				}
				if a == b || a == c || a == d || b == c || b == d || c == d {
					fail = true
				}
			}
		}
		if fail {
			t.Errorf("Testing fail n: %d m: %d", n, m)
		}
	}
}
