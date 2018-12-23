package main_test

import (
	"fmt"
	"math/rand"
	"testing"

	"./"
)

func TestMAXEP(t *testing.T) {
	rand := rand.New(rand.NewSource(100))
	var n, c, v int
	for tc := 500; tc > 0; tc-- {
		n = rand.Intn(150000) + 1
		c = rand.Intn(150) + 1
		for i := 1; i <= 1000; i++ {
			v = rand.Intn(n) + 1
			shownParams := 0
			points := 1000
			r := 0
			main.Solve(func() int {
				if shownParams == 0 {
					shownParams = 1
					return n
				} else if shownParams == 1 {
					shownParams = 2
					return c
				} else {
					return r
				}
			}, func(s string) {
				var typ, val int
				fmt.Sscanf(s, "%d %d\n", &typ, &val)
				if points <= 0 {
					r = -1
					t.Errorf("No coins. N: %d, c: %d, v: %d", n, c, v)
				} else if typ == 1 {
					if val < 1 || val > n {
						r = -1
						t.Errorf("Value out of range. N: %d, c: %d, v: %d, val: %d", n, c, v, val)
					} else if val < v {
						r = 0
					} else {
						r = 1
					}
					points--
				} else if typ == 2 {
					r = 0
					points -= c
				} else {
					if val == v {
						r = 0
					} else {
						r = -1
						t.Errorf("Wrong answer. N: %d, c: %d, v: %d, res: %d", n, c, v, val)
					}
				}
			})
		}
	}
}
