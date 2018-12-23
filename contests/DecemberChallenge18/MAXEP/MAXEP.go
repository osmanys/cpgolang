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
	Solve(func() int {
		return nextInt()
	}, func(s string) {
		defer writer.Flush()
		print(s)
	})
}

// Solve function
func Solve(scan func() int, print func(string)) {
	n := scan()
	scan()

	start := 1
	end := n
	var step, r, i int
	for start < end {
		if (end-start)/10 == 0 {
			step = 1
		} else {
			step = (end - start) / 10
		}
		for i = start; i <= end; i += step {
			print(fmt.Sprintf("1 %d\n", i))
			r = scan()
			if r == 1 {
				print("2\n")
				break
			} else if r == -1 {
				return
			}
		}
		if i < end {
			end = i
		}
		if i > start {
			start = i - step + 1
		}
	}

	print(fmt.Sprintf("3 %d\n", end))
}
