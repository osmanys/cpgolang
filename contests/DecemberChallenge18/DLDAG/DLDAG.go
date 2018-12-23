package main

import (
	"bufio"
	"container/heap"
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

	var n, m, u, v int
	scanf("%d %d\n", &n, &m)
	g := NewGraph(n)
	revg := NewGraph(n)
	for i := 0; i < m; i++ {
		scanf("%d %d\n", &u, &v)
		g.AddEdge(u-1, v-1)
		revg.AddEdge(v-1, u-1)
	}
	r, cnt := Solve(g, revg)
	printf("%d\n", cnt)
	for i := 0; i < cnt; i++ {
		printf("%d ", r[i][0])
		for j := 1; j <= r[i][0]; j++ {
			printf("%d ", r[i][j]+1)
		}
		printf("\n")
	}
}

// Solve func
func Solve(g *Graph, revg *Graph) ([][]int, int) {
	n := g.V
	p := make([]int, n)
	ridx := 0
	g.DFS(func(v int) {
		p[v] = ridx
		ridx++
	}, func(int) {}, func(int, int, bool) {}, func(int, int, bool) {})
	var add int
	childs := make([]int, n)
	pq := make(PriorityQueue, 0)
	toAdd := make([]int, 0)
	for i := 0; i < n; i++ {
		childs[i] = len(g.adj[i])
		if childs[i] == 0 {
			toAdd = append(toAdd, i)
		}
	}
	for _, add = range toAdd {
		heap.Push(&pq, &heapItem{value: add, priority: p[add]})
	}
	var v *heapItem
	r := make([][]int, n)
	ridx = 0
	for cnt := pq.Len(); cnt > 0; cnt = pq.Len() {
		r[ridx] = make([]int, 3)
		v = heap.Pop(&pq).(*heapItem)
		r[ridx][0] = 1
		r[ridx][1] = v.value
		toAdd = make([]int, 0)
		for i := 0; i < len(revg.adj[v.value]); i++ {
			childs[revg.adj[v.value][i]]--
			if childs[revg.adj[v.value][i]] == 0 {
				toAdd = append(toAdd, revg.adj[v.value][i])
			}
		}
		if cnt > 1 {
			v = heap.Pop(&pq).(*heapItem)
			r[ridx][0] = 2
			r[ridx][2] = v.value
			for i := 0; i < len(revg.adj[v.value]); i++ {
				childs[revg.adj[v.value][i]]--
				if childs[revg.adj[v.value][i]] == 0 {
					toAdd = append(toAdd, revg.adj[v.value][i])
				}
			}
		}
		for _, add = range toAdd {
			heap.Push(&pq, &heapItem{value: add, priority: p[add]})
		}
		ridx++
	}
	return r, ridx
}

type heapItem struct {
	value    int
	priority int
	index    int
}

// PriorityQueue struct
type PriorityQueue []*heapItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push function
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*heapItem)
	item.index = n
	*pq = append(*pq, item)
}

// Pop function
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Pair struct
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

func (g *Graph) dfsUtil(v int, visited []bool, fprev func(int), fpostv func(int), fpree func(int, int, bool), fposte func(int, int, bool)) {
	visited[v] = true
	fprev(v)
	var visn bool
	for _, n := range g.adj[v] {
		visn = visited[n]
		fpree(v, n, visn)
		if !visited[n] {
			g.dfsUtil(n, visited, fprev, fpostv, fpree, fposte)
		}
		fposte(v, n, visn)
	}
	fpostv(v)
}

// DFS traversal of the vertices
func (g *Graph) DFS(fprev func(int), fpostv func(int), fpree func(int, int, bool), fposte func(int, int, bool)) {
	visited := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.dfsUtil(i, visited, fprev, fpostv, fpree, fposte)
		}
	}
}
