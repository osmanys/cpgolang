package main_test

import (
	"testing"

	"./"
)

func TestEDGEDIR(t *testing.T) {
	g := main.NewGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	g.DFS()
}
