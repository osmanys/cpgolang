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

	var l, r, k, q, u, v, n, m, tc int
	scanf("%d\n", &tc)
	for ; tc > 0; tc-- {
		scanf("%d %d\n", &n, &m)
		for i := 0; i < m; i++ {
			scanf("%d %d\n", &u, &v)

		}
		scanf("%d\n", &q)
		for i := 0; i < q; i++ {
			scanf("%d ", &k)
			for j := 0; j < k-1; j++ {
				scanf("%d %d ", &l, &r)
			}
			scanf("%d %d\n", &l, &r)
		}
	}
}
