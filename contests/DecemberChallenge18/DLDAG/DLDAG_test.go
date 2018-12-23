package main_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"./"
)

func TestDLDAG(t *testing.T) {
	var u, v int
	rand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	total := 0
	for tc := 0; tc < 10; tc++ {
		n := 10000
		preg := main.NewGraph(n)
		for i := 0; i < n-1; i++ {
			//u = i + 1
			//v = i + 2
			u = rand.Intn(n) + 1
			v = rand.Intn(n) + 1
			if u != v {
				preg.AddEdge(u-1, v-1)
			}
		}
		g := main.NewGraph(n)
		revg := main.NewGraph(n)
		visiting := make([]bool, n)
		preg.DFS(func(v int) {
			visiting[v] = true
		}, func(v int) {
			visiting[v] = false
		}, func(u int, v int, visited bool) {
			if !visited || !visiting[v] {
				g.AddEdge(u, v)
				revg.AddEdge(v, u)
			}
		}, func(int, int, bool) {})
		_, cnt := main.Solve(g, revg)
		total += cnt
	}
	fmt.Printf("%d\n", total/10)
}
