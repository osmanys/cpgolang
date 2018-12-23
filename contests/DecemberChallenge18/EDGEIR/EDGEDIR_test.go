package main_test

import (
	"math/rand"
	"testing"
	"time"

	"./"
)

func TestEDGEDIR(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var j, i, u, v, n, m int
	var r []bool
	var edges []main.Pair
	var g *main.Graph
	var found bool

	for tc := 1; tc > 0; tc-- {
		n = rand.Intn(100000) + 1
		m = rand.Intn(100000) + 1
		edges = make([]main.Pair, 0)
		for i = 0; i < m; i++ {
			u = rand.Intn(n)
			v = rand.Intn(n)
			if u == v {
				continue
			}
			for j = 0; j < len(edges); j++ {
				if (edges[j].U == u && edges[j].V == v) || (edges[j].U == v && edges[j].V == u) {
					break
				}
			}
			if j == len(edges) {
				edges = append(edges, main.Pair{U: u, V: v})
			}
		}
		m = len(edges)
		if m%2 == 1 {
			edges = edges[:m-1]
			m--
		}

		g = main.NewGraph(n)
		for _, edge := range edges {
			u = edge.U
			v = edge.V
			g.AddEdge(u, v)
			g.AddEdge(v, u)
		}
		if m%2 == 0 {
			found = true
		} else {
			found = false
		}
		r = main.Solve(g, n, m, edges)
		if r == nil {
			if found {
				t.Errorf("Wrong answer. Config not found. %d %d", n, m)
			}
		} else {
			indegree := make([]int, n)
			for i = 0; i < m; i++ {
				if r[i] {
					indegree[edges[i].V]++
				} else {
					indegree[edges[i].U]++
				}
			}
			for i = 0; i < n; i++ {
				if indegree[i]%2 == 1 {
					break
				}
			}
			if i < n {
				t.Errorf("Wrong answer. Exist vertex odd. %d %d %d %v", n, m, edges, r)
			}
			if !found {
				t.Errorf("Wrong answer. Config found incorrectly. %d %d %d %v", n, m, edges, r)
			}
		}
	}
}

func findSolve(n int, m int, edges []main.Pair) bool {
	indegree := make([]int, n)
	return findSolveRecursive(n, m, edges, 0, indegree)
}

func findSolveRecursive(n int, m int, edges []main.Pair, idx int, indegree []int) bool {
	if idx == m {
		var i int
		for i = 0; i < n; i++ {
			if indegree[i]%2 == 1 {
				break
			}
		}
		if i < n {
			return false
		} else {
			return true
		}
	} else {
		var r bool
		indegree[edges[idx].V]++
		r = findSolveRecursive(n, m, edges, idx+1, indegree)
		indegree[edges[idx].V]--
		if r {
			return true
		}
		indegree[edges[idx].U]++
		r = findSolveRecursive(n, m, edges, idx+1, indegree)
		indegree[edges[idx].U]--
		return r
	}
}
