package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var m []int64
	var t2 = []int{}
	var n, q, tmp int64
	var xi string = "3 10 8 6 11"
	//no of shop in the city
	fmt.Scanf("%d", &n)
	//prices of the bottles in i-th shop
	//	fmt.Scanf("%s", &xi)

	//the number of days Vasiliy plans to buy the drink
	fmt.Scanf("%d", &q)
	// number of coins can be spent on the i-th day
	for i := 0; int64(i) < q; i++ {
		fmt.Scanf("%d", &tmp)
		m = append(m, tmp)
	}

	//sort the drink price and convert to int
	x := strings.Split(xi, " ")
	for _, i := range x {
		j, err := strconv.ParseInt(i, 10, 8)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, int(j))
	}
	sort.Ints(t2)

	//fmt.Println("m=", m, " t2=", t2)
	for _, i := range m {
		for j, k := range t2 {
			//fmt.Printf("i=%d -- k=%d\n", i, k)
			if i < int64(k) {
				fmt.Println(j)
				break
			} else if i == int64(k) {
				fmt.Println(j + 1)
				break
			}
		}
	}
}
