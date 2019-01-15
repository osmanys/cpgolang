package main_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"."
)

func TestGRAPART(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	tc := 5
	limit := 10
	var k, u, n int
	var perm, neighbor, a []int
	var accept bool
	var ga []bool
	var g *main.Graph
	fmt.Printf("%d\n", tc)
	for ; tc > 0; tc-- {
		n = limit //rand.Intn(limit-1) + 1
		if n%2 == 1 {
			n++
		}
		g = main.NewGraph(n)
		perm = rand.Perm(n)
		for i := 1; i < n; i++ {
			u = rand.Intn(i)
			g.AddEdge(perm[i], perm[u])
			g.AddEdge(perm[u], perm[i])
		}
		k, a = main.Solve(n, g)

		ga = make([]bool, n)
		for i := 0; i < n/2; i++ {
			ga[a[i]] = true
		}
		for i := 0; i < n; i++ {
			neighbor = make([]int, 0)
			accept = false
			for j := 1; j <= g.Adj[i][0]; j++ {
				d := g.Adj[i][j]
				neighbor = append(neighbor, d)
				if (ga[i] && !ga[d]) || (!ga[i] && ga[d]) {
					accept = true
					break
				}
			}
			if !accept && k == 2 {
				for _, ng := range neighbor {
					for j := 1; j <= g.Adj[ng][0]; j++ {
						d := g.Adj[ng][j]
						if (ga[i] && !ga[d]) || (!ga[i] && ga[d]) {
							accept = true
							break
						}
					}
				}
			}
			if !accept {
				/*for i := 0; i < n; i++ {
					for j := 1; j <= g.Adj[i][0]; j++ {
						d := g.Adj[i][j]
						if i < d {
							t.Errorf("%d %d\n", i, d)
						}
					}
				}
				t.Errorf("%v\n", a[:n/2])
				t.Errorf("%v\n", a[n/2:])*/
				t.Errorf("Error case n: %d, k: %d, node: %d", n, k, i)
				break
			}
		}
		fmt.Printf("%d\n", n)
		for i := 0; i < n; i++ {
			for j := 1; j <= g.Adj[i][0]; j++ {
				d := g.Adj[i][j]
				if i < d {
					fmt.Printf("%d %d\n", i+1, d+1)
				}
			}
		}
		fmt.Printf("k: %d\n", k)
		fmt.Printf("%v\n", a[:n/2])
		fmt.Printf("%v\n", a[n/2:])
	}
	t.Errorf("")
}
