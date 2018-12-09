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
	defer writer.Flush()

	var n, c int
	scanf("%d %d\n", &n, &c)

	printf("%d\n", n+c)
}
