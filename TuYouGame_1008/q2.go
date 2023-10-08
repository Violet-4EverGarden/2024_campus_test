package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Printf("%d %d\n", i, j)
		}
	}
	return
}
