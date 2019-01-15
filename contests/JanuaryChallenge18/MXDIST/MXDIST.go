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

	var f, t, q, x, y, n int
	scanf("%d\n", &n)
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		scanf("%d %d\n", &x, &y)
		points[i] = Point{x, y}
	}
	scanf("%d\n", &q)
	for i := 0; i < q; i++ {
		scanf("%d %d\n", &f, &t)
		printf("%d\n", Solve(points[f-1:t]))
	}
}

// Point struct
type Point struct {
	x int
	y int
}

// ByPoint alias
type ByPoint []Point

func (a ByPoint) Len() int      { return len(a) }
func (a ByPoint) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//func (a ByPoint) Less(i, j int) bool { return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y) }
func (a ByPoint) Less(i, j int) bool { return a[i].x < a[j].x }

// Solve func
func Solve(points []Point) int {
	l, u := hull(points)
	fmt.Printf("hull l: %v\n", l)
	fmt.Printf("hull u: %v\n", u)
	i := 0
	j := len(l) - 1
	m := len(u) - 1
	res := -1
	var dist int
	for i < m || j > 0 {
		dist = distance(u[i], l[j])
		fmt.Printf("dist: %d\n", dist)
		if res < dist {
			res = dist
		}
		if i == m {
			j--
		} else if j == 0 {
			i++
		} else {
			if (u[i+1].y-u[i].y)*(l[j].x-l[j-1].x) > (l[j].y-l[j-1].y)*(u[i+1].x-u[i].x) {
				i++
			} else {
				j--
			}
		}
	}
	return res
}

func cross(p, q, r Point) int {
	return (q.x-p.x)*(r.y-p.y) - (r.x-p.x)*(q.y-p.y)
}

func distance(p, q Point) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
}

func hull(p []Point) ([]Point, []Point) {
	n := len(p)
	j := 0
	k := 0
	l := make([]Point, 2*n)
	u := make([]Point, 2*n)
	sort.Sort(ByPoint(p))
	for i := 0; i < n; i++ {
		for j >= 2 && cross(l[j-2], l[j-1], p[i]) <= 0 {
			j--
		}
		for k >= 2 && cross(u[k-2], u[k-1], p[i]) >= 0 {
			k--
		}
		l[j] = p[i]
		j++
		u[k] = p[i]
		k++
	}
	return u[:k], l[:j]
}
