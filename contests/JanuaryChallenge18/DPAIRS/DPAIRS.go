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

	var d, n, m int
	scanf("%d %d\n", &n, &m)
	nmin := 1000000000
	imin := -1
	for i := 0; i < n; i++ {
		scan(&d)
		if nmin >= d {
			nmin = d
			imin = i
		}
	}
	mmax := -1000000000
	imax := 0
	for i := 0; i < m; i++ {
		scan(&d)
		if mmax <= d {
			mmax = d
			imax = i
		}
	}
	for i := 0; i < m; i++ {
		printf("%d %d\n", imin, i)
	}
	for i := 0; i < n; i++ {
		if i != imin {
			printf("%d %d\n", i, imax)
		}
	}
}
