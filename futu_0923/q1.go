package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	var nm = make([][]int, 0)
	for i := 0; i < T; i++ {
		tmp := make([]int, 2)
		fmt.Scan(&tmp[0], &tmp[1])
		nm = append(nm, tmp)
	}
	// fmt.Println(nm)
	for i := 0; i < T; i++ {
		n, m := nm[i][0], nm[i][1]
		a := n / m
		b := n % m
		t := a + boolToInt(b != 0)
		fmt.Println(t)
		fmt.Printf("%d", n-(t-1)*m)
		for j := 1; j < t; j++ {
			fmt.Printf(" %d", m)
		}
		fmt.Print("\n")
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
