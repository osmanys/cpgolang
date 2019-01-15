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

func main() {
	defer writer.Flush()

	var g *Graph
	var k, u, v, n, tc int
	var a []int
	scanf("%d\n", &tc)
	for ; tc > 0; tc-- {
		scanf("%d\n", &n)
		g = NewGraph(n)
		for i := 0; i < n-1; i++ {
			scanf("%d %d\n", &u, &v)
			g.AddEdge(u-1, v-1)
			g.AddEdge(v-1, u-1)
		}
		k, a = Solve(n, g)
		printf("%d\n", k)
		for i := 0; i < n/2; i++ {
			printf("%d ", a[i]+1)
		}
		printf("\n")
		for i := n / 2; i < n; i++ {
			printf("%d ", a[i]+1)
		}
		printf("\n")
	}
}

var res = make([]int, 10000)

// Solve func
func Solve(n int, g *Graph) (int, []int) {
	depth := make([]int, n)
	arr := make([][]int, 4)
	g.DFS(func(v int) {
		if g.Adj[v][0] > 1 {
			arr[2*(depth[v]%2)] = append(arr[2*(depth[v]%2)], v)
		} else {
			arr[2*(depth[v]%2)+1] = append(arr[2*(depth[v]%2)+1], v)
		}
	}, func(u, v int, visited bool) {
		if !visited {
			depth[v] = depth[u] + 1
		}
	}, func(u, v int, visited bool) {
	}, func(v int) {
	})
	copy(res, arr[0])
	copy(res[len(arr[0]):], arr[1])
	copy(res[len(arr[0])+len(arr[1]):], arr[3])
	copy(res[len(arr[0])+len(arr[1])+len(arr[3]):], arr[2])
	var k = 1
	if len(arr[0])+len(arr[1]) != len(arr[2])+len(arr[3]) {
		k = 2
	}
	return k, res[:n]
}

var graph = &Graph{}

// Graph struct
type Graph struct {
	V   int
	Adj [10000][10000]int
}

// NewGraph function
func NewGraph(v int) *Graph {
	graph.V = v
	for i := 0; i < v; i++ {
		graph.Adj[i][0] = 0
	}
	return graph
}

// AddEdge function
func (g *Graph) AddEdge(v int, w int) {
	g.Adj[v][0]++
	g.Adj[v][g.Adj[v][0]] = w
}

// Len function
func (g *Graph) Len(v int) int {
	return g.Adj[v][0]
}

func (g *Graph) dfsUtil(v int, visited []bool, fprev func(int), fpree func(int, int, bool), fposte func(int, int, bool), fpostv func(int)) {
	visited[v] = true
	fprev(v)
	var visn bool
	var n int
	for i := 1; i <= g.Adj[v][0]; i++ {
		n = g.Adj[v][i]
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
