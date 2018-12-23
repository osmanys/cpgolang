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
	t := nextInt()
	var n int
	var r []int
	for ; t > 0; t-- {
		n = nextInt()
		r = Solve(n, func(a int, b int, c int) int {
			printf("1 %d %d %d\n", a, b, c)
			writer.Flush()
			return nextInt()
		})
		print(2)
		for _, v := range r {
			printf(" %d", v)
		}
		print("\n")
		writer.Flush()
		nextInt()
	}
}

// Solve function
func Solve(n int, eval func(int, int, int) int) []int {
	r := make([]int, n)
	i := 0
	for ; i < n/4-2; i++ {
		calc4(4*i, eval, r)
	}
	switch n % 4 {
	case 0:
		calc4(4*i, eval, r)
		calc4(4*i+4, eval, r)
	case 1:
		calc4(4*i, eval, r)
		calc5(4*i+4, eval, r)
	case 2:
		calc5(4*i, eval, r)
		calc5(4*i+5, eval, r)
	case 3:
		calc5(4*i, eval, r)
		calc6(4*i+5, eval, r)
	}
	return r
}

func calc4(i int, eval func(int, int, int) int, r []int) {
	var x1, x2, x3, x4 int
	x1 = eval(i+2, i+3, i+4)
	x2 = eval(i+1, i+3, i+4)
	x3 = eval(i+1, i+2, i+4)
	x4 = eval(i+1, i+2, i+3)

	r[i] = x2 ^ x3 ^ x4
	r[i+1] = x1 ^ x3 ^ x4
	r[i+2] = x1 ^ x2 ^ x4
	r[i+3] = x1 ^ x2 ^ x3
}

func calc5(i int, eval func(int, int, int) int, r []int) {
	x1 := eval(i+1, i+2, i+3)
	x2 := eval(i+2, i+3, i+4)
	x3 := eval(i+3, i+4, i+5)
	x4 := eval(i+1, i+4, i+5)
	x5 := eval(i+1, i+2, i+5)

	r[i] = x2 ^ x3 ^ x5
	r[i+1] = x1 ^ x3 ^ x4
	r[i+2] = x2 ^ x4 ^ x5
	r[i+3] = x1 ^ x3 ^ x5
	r[i+4] = x1 ^ x2 ^ x4
}

func calc6(i int, eval func(int, int, int) int, r []int) {
	x1 := eval(i+1, i+2, i+3)
	x2 := eval(i+3, i+4, i+5)
	x3 := eval(i+1, i+5, i+6)
	x4 := eval(i+3, i+4, i+6)
	x5 := eval(i+2, i+5, i+6)
	x6 := eval(i+1, i+2, i+4)

	r[i] = x2 ^ x3 ^ x4
	r[i+1] = x2 ^ x4 ^ x5
	r[i+2] = x1 ^ x3 ^ x5
	r[i+3] = x3 ^ x5 ^ x6
	r[i+4] = x1 ^ x2 ^ x6
	r[i+5] = x1 ^ x4 ^ x6
}
