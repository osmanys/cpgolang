package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func next() string {
	var s string
	for _, err := fmt.Fscanf(reader, "%v", &s); err != nil; _, err = fmt.Fscanf(reader, "%d", &s) {
	}
	return s
}
func nextLine() string {
	s, _ := reader.ReadString('\n')
	return s
}
func nextInt() int {
	var s int
	for _, err := fmt.Fscanf(reader, "%d", &s); err != nil; _, err = fmt.Fscanf(reader, "%d", &s) {
	}
	return s
}
func nextFloat() float64 {
	var s float64
	for _, err := fmt.Fscanf(reader, "%f", &s); err != nil; _, err = fmt.Fscanf(reader, "%d", &s) {
	}
	return s
}
func nextBoolean() bool {
	var s bool
	for _, err := fmt.Fscanf(reader, "%t", &s); err != nil; _, err = fmt.Fscanf(reader, "%d", &s) {
	}
	return s
}

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func scan(a ...interface{})             { fmt.Fscan(reader, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func print(a ...interface{})            { fmt.Fprint(writer, a...) }

func main() {
	defer writer.Flush()

	var u, v, n, m int
	var g *Graph
	var edges []Pair
	var r []bool
	t := nextInt()
	for ; t > 0; t-- {
		n = nextInt()
		m = nextInt()
		g = NewGraph(n)
		edges = make([]Pair, m)
		for i := 0; i < m; i++ {
			u = nextInt()
			v = nextInt()
			g.AddEdge(u-1, v-1)
			g.AddEdge(v-1, u-1)
			edges[i] = Pair{u - 1, v - 1}
		}
		r = Solve(g, n, m, edges)
		if r == nil {
			printf("-1\n")
		} else {
			for _, edge := range r {
				if edge {
					printf("0 ")
				} else {
					printf("1 ")
				}
			}
			printf("\n")
		}
	}
}

//Solve function
func Solve(g *Graph, n int, m int, edges []Pair) []bool {
	if m%2 == 1 {
		return nil
	}
	indegree := make([]int, n)
	edgesDir := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		edgesDir[i] = make(map[int]int, 0)
	}
	g.DFS(func(v int) {
	}, func(u int, v int, visited bool) {
		if edgesDir[u][v] == 0 {
			edgesDir[v][u] = 2
		}
	}, func(u int, v int, visited bool) {
		if edgesDir[u][v] == 0 {
			if visited {
				edgesDir[u][v] = 1
				edgesDir[v][u] = -1
				indegree[v]++
			} else if indegree[v]%2 == 1 {
				edgesDir[u][v] = 1
				edgesDir[v][u] = -1
				indegree[v]++
			} else {
				edgesDir[u][v] = -1
				edgesDir[v][u] = 1
				indegree[u]++
			}
		}
	}, func(v int) {
	})
	i := 0
	for ; i < n; i++ {
		if indegree[i]%2 == 1 {
			break
		}
	}
	if i < n {
		return nil
	}
	r := make([]bool, 0)
	for _, edge := range edges {
		if edgesDir[edge.U][edge.V] == 1 {
			r = append(r, true)
		} else {
			r = append(r, false)
		}
	}
	return r
}

//Pair struct
type Pair struct {
	U int
	V int
}

// Graph struct
type Graph struct {
	V   int
	adj [][]int
}

// NewGraph function
func NewGraph(v int) *Graph {
	return &Graph{v, make([][]int, v)}
}

// AddEdge function
func (g *Graph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
}

func (g *Graph) dfsUtil(v int, visited []bool, fprev func(int), fpree func(int, int, bool), fposte func(int, int, bool), fpostv func(int)) {
	visited[v] = true
	fprev(v)
	var visn bool
	for _, n := range g.adj[v] {
		visn = visited[n]
		fpree(v, n, visn)
		if !visited[n] {
			g.dfsUtil(n, visited, fprev, fpree, fposte, fpostv)
		}
		fposte(v, n, visn)
	}
	fpostv(v)
}

// DFS traversal of the vertices
func (g *Graph) DFS(fprev func(int), fpree func(int, int, bool), fposte func(int, int, bool), fpostv func(int)) {
	visited := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.dfsUtil(i, visited, fprev, fpree, fposte, fpostv)
		}
	}
}
