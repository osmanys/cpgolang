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

	var n, m, N, t, x, y, z, tc int
	scanf("%d\n", &tc)
	for ; tc > 0; tc-- {
		scanf("%d %d %d %d %d\n", &N, &t, &x, &y, &z)
		n, m = Solve(N, t, x, y, z)
		printf("%d %d\n", n, m)
	}
}

// Solve func
func Solve(N, t, x, y, z int) (int, int) {
	var _type, n int
	switch t {
	// AxAyBz
	case 1:
		if x < y {
			if y < z {
				_type = 4
			} else {
				_type = 1
			}
		} else {
			if y < z {
				_type = 3
			} else {
				_type = 5
			}
		}
	// AxByAz
	case 2:
		_type = 2
	// AxByBz
	case 3:
		if x < y {
			if y < z {
				_type = 5
			} else {
				_type = 1
			}
		} else {
			if y < z {
				_type = 3
			} else {
				_type = 4
			}
		}
	// BxAyBz
	case 4:
		_type = 2
	}
	switch _type {
	case 1:
		n = (y - 1)
	case 2:
		n = (2*N + 1) - 2*y
	case 3:
		n = (y + 1)
	case 4:
		n = (2*N + 1) - (y + 1)
	case 5:
		n = (2*N + 1) - (y - 1)
	}
	_gcd := gcd(n, 2*N+1)
	return n / _gcd, (2*N + 1) / _gcd
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
