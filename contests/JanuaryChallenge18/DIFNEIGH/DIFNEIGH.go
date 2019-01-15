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

	var k, tc, n, m int
	var r [][]int
	scanf("%d\n", &tc)
	for ; tc > 0; tc-- {
		scanf("%d %d\n", &n, &m)
		k, r = Solve(n, m)
		printf("%d\n", k)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				printf("%d ", r[i][j])
			}
			printf("\n")
		}
	}
}

// Solve func
func Solve(n, m int) (int, [][]int) {
	k1, r1 := solveV(n, m)
	k2, r2 := solveH(n, m)
	if k1 < k2 {
		return k1, r1
	}
	return k2, r2
}

func solveV(n, m int) (int, [][]int) {
	k := 0
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, m)
	}
	mapp := make(map[int]int)
	var vl, v int
	var has bool
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if n == 2 || m == 2 {
				if n == 2 {
					vl = (j + 1) % 3
				} else {
					vl = (i + 1) % 3
				}
			} else {
				vl = val(i, j)
			}
			if v, has = mapp[vl]; has {
				r[i][j] = v
			} else {
				k++
				mapp[vl] = k
				r[i][j] = k
			}
		}
	}
	return k, r
}

func solveH(n, m int) (int, [][]int) {
	k := 0
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, m)
	}
	mapp := make(map[int]int)
	var vl, v int
	var exist bool
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if n == 2 || m == 2 {
				if n == 2 {
					vl = (j + 1) % 3
				} else {
					vl = (i + 1) % 3
				}
			} else {
				vl = val(j, i)
			}
			if v, exist = mapp[vl]; exist {
				r[i][j] = v
			} else {
				k++
				mapp[vl] = k
				r[i][j] = k
			}
		}
	}
	return k, r
}

func val(x, y int) int {
	if x%4 == 0 {
		if y%4 < 2 {
			return 1
		}
		return 2
	} else if x%4 == 1 {
		if y%4 < 2 {
			return 3
		}
		return 4
	} else if x%4 == 2 {
		if y%4 < 2 {
			return 2
		}
		return 1
	}
	if y%4 < 2 {
		return 4
	}
	return 3
}
