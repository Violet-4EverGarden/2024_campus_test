package main

import "fmt"

// 小红的棋子移动
// 思路：如果行、列均为偶数或均为奇数，则小红必输
func main() {
	var t int
	fmt.Scan(&t)
	res := make([]string, 0)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		if (n%2)^(m%2) != 0 {
			res = append(res, "Yes")
		} else {
			res = append(res, "No")
		}
	}
	fmt.Print(res[0])
	for i := 1; i < len(res); i++ {
		fmt.Print("\n")
		fmt.Print(res[i])
	}
}
