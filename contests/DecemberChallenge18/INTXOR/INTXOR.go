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
	var x1, x2, x3, x4 int
	i := 0
	for ; i < n/4; i++ {
		x1 = eval(4*i+2, 4*i+3, 4*i+4)
		x2 = eval(4*i+1, 4*i+3, 4*i+4)
		x3 = eval(4*i+1, 4*i+2, 4*i+4)
		x4 = eval(4*i+1, 4*i+2, 4*i+3)

		r[4*i] = x2 ^ x3 ^ x4
		r[4*i+1] = x1 ^ x3 ^ x4
		r[4*i+2] = x1 ^ x2 ^ x4
		r[4*i+3] = x1 ^ x2 ^ x3
	}
	for i := 4 * i; i < n; i++ {
		r[i] = eval(1, 2, i+1) ^ r[0] ^ r[1]
	}
	return r
}
