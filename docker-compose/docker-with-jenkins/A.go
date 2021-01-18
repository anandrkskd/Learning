package main

import "fmt"

func main() {
	a := [5]int64{100, 20, 10, 5, 1}
	var input, noOfBill int64
	fmt.Scanf("%d", &input)
	for _, j := range a {
		if input/j > 0 && input%j == 0 {
			noOfBill = noOfBill + input/int64(j)
			break
		} else if input/j > 0 {
			noOfBill = noOfBill + input/int64(j)
		}
		input = input % j
	}
	fmt.Println(noOfBill)
}
