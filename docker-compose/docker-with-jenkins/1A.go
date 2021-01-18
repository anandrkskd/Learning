package main

import "fmt"

func main() {
	var m, n, a int64
	//var noOfSquare int64 = 0

	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)

	ma := m / a
	na := n / a
	if m%a != 0 {
		ma = ma + 1
	}
	if n%a != 0 {
		na = na + 1
	}

	fmt.Println(ma * na)
}
