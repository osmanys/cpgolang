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
	/*defer writer.Flush()

	var n, m int
	var grade []int
	t := nextInt()
	for ; t > 0; t-- {
		n = nextInt()
		m = nextInt()
		grade = make([]int, n)
		for i := 0; i < m; i++ {

		}
	}*/
}

//Solve function
func Solve() {

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

func (g *Graph) dfsUtil(v int, visited []bool) {
	visited[v] = true

	//code here
	fmt.Printf("%d ", v)

	for _, n := range g.adj[v] {
		if !visited[n] {
			g.dfsUtil(n, visited)
		}
	}
}

// DFS traversal of the vertices reachable from v.
// It uses recursive DFSUtil()
func (g *Graph) DFS() {
	visited := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.dfsUtil(i, visited)
		}
	}
}
