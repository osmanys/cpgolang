package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var n, c, r, i int
	scanf("%d %d\n", &n, &c)

	start := 1
	end := n

	for start < end {
		for i = 0; start+i*((end-start)/10) <= end; i++ {
			printf("1 %d\n", start+i*((end-start)/10))
			writer.Flush()
			scanf("%d\n", &r)
			if r == 1 {
				break
			} else if r == -1 {
				return
			}
		}
		start = start + (i-1)*((end-start)/10)
		if start+i*((end-start)/10) <= end {
			end = start + i*((end-start)/10) - 1
		}
	}

	printf("2 %d\n", start)
}
