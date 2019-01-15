package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func scan(a ...interface{})             { fmt.Fscan(reader, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func print(a ...interface{})            { fmt.Fprint(writer, a...) }

func main() {
	defer writer.Flush()

	var x, y, n, m int
	scanf("%d %d\n", &n, &m)
	points := make([]Point, 2*m)
	for i := 0; i < 2*m; i++ {
		scanf("%d %d\n", &x, &y)
		points[i] = Point{x, y}
	}
	cells := make([][]int, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scan(&cells[i][j])
		}
	}
	res := Solve(n, m, points, cells)
	for i := 0; i < m; i++ {
		printf("%d ", len(res[i]))
		for j := 0; j < len(res[i]); j++ {
			printf("%d %d ", res[i][j].x, res[i][j].y)
		}
		printf("\n")
	}
}

// Point struct
type Point struct {
	x int
	y int
}

// ByPoint alias
type ByPoint []Point

func (a ByPoint) Len() int           { return len(a) }
func (a ByPoint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoint) Less(i, j int) bool { return a[i].y < a[j].y }

// Solve function
func Solve(n, m int, points []Point, cells [][]int) [][]Point {
	idxPoints := make([][]Point, 1000)
	for i := 0; i < 2*m; i++ {
		idxPoints[points[i].x] = append(idxPoints[points[i].x], points[i])
	}
	sortPoints := make([][]Point, 0)
	for i := 0; i < 1000; i++ {
		if len(idxPoints[i]) > 0 {
			sort.Sort(ByPoint(idxPoints[i]))
			sortPoints = append(sortPoints, idxPoints[i])
		}
	}
	res := make([][]Point, m)
	var cur Point
	loaded := false
	idx := 0
	for i := 0; i < len(sortPoints); i++ {
		if i%2 == 0 {
			for j := 0; j < len(sortPoints[i]); j++ {
				if !loaded {
					loaded = true
					cur = sortPoints[i][j]
				} else {
					if j == 0 {
						res[idx] = buildPath(cur, sortPoints[i][j], false)
						idx++
					} else {
						res[idx] = buildPath(cur, sortPoints[i][j], true)
						idx++
					}
					loaded = false
				}
			}
		} else {
			for j := len(sortPoints[i]) - 1; j >= 0; j-- {
				if !loaded {
					loaded = true
					cur = sortPoints[i][j]
				} else {
					res[idx] = buildPath(cur, sortPoints[i][j], true)
					idx++
					loaded = false
				}
			}
		}
	}
	return res
}

func buildPath(p1, p2 Point, bottom bool) []Point {
	res := make([]Point, 0)
	if bottom {
		if p1.y < p2.y {
			for i := p1.y; i < p2.y; i++ {
				res = append(res, Point{p1.x, i})
			}
			for i := p1.x; i <= p2.x; i++ {
				res = append(res, Point{i, p2.y})
			}
		} else {
			for i := p1.x; i < p2.x; i++ {
				res = append(res, Point{i, p1.y})
			}
			for i := p1.y; i >= p2.y; i-- {
				res = append(res, Point{p2.x, i})
			}
		}
	} else {
		if p1.y < p2.y {
			for i := p1.x; i < p2.x; i++ {
				res = append(res, Point{i, p1.y})
			}
			for i := p1.y; i <= p2.y; i++ {
				res = append(res, Point{p2.x, i})
			}
		} else {
			for i := p1.y; i > p2.y; i-- {
				res = append(res, Point{p1.x, i})
			}
			for i := p1.x; i <= p2.x; i++ {
				res = append(res, Point{i, p2.y})
			}
		}
	}
	return res
}
