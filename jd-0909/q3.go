package main

import "fmt"

// 计算子数组权值之和
// 数组权值计算公式：1*a[1]+2*a[2]+
func main() {
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		tmp := a[i] // a[i] * 1
		res += tmp
		count := 2
		for j := i + 1; j < n; j++ {
			tmp += a[j] * count
			count++
			res += tmp
		}
	}
	fmt.Print(res)
}
