package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var res = make([]int, 0)
	res = append(res, a[0])
	a = a[1:]
	for len(a) > 0 {
		a = append(a[1:], a[0])
		res = append(res, a[0])
		a = a[1:]
	}
	fmt.Print(res[0])
	for i := 1; i < len(res); i++ {
		fmt.Printf(" %d", res[i])
	}
}

/*
6
2 3 2 1 4 5
*/
