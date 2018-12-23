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

	var d, c, t, q, k int
	ans := 0
	scanf("%d %d\n", &q, &k)
	color := make([]int, 0)
	del := make([]int, 0)
	for i := 0; i < q; i++ {
		scanf("%d %d", &t, &c)
		c = c ^ ans
		if t == 1 {
			scanf(" %d\n", &d)
			color = append([]int{c}, color...)
			del = append([]int{d}, del...)
		} else if t == 2 {
			scanf(" %d\n", &d)
			color = append(color, c)
			del = append(del, d)
		} else {
			scanf("\n")
			ans = maxSubArraySum(del, color, k, c)
			printf("%d\n", ans)
		}
	}
}

func maxSubArraySum(del []int, color []int, k int, c int) int {
	maxSoFar := 0
	maxEndingHere := 0
	for i := 0; i < len(del); i++ {
		if color[i] < c-k || color[i] > c+k {
			continue
		}
		maxEndingHere = maxEndingHere + del[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		} else if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}
