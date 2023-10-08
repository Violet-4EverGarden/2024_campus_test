package main

import "fmt"

// 雪球：其实就是一个按高度乘x的公式，可以通过逆序遍历来简写成以下格式
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = (ans*x + a[i]) % (1e9 + 7)
	}
	ans = (ans * x) % (1e9 + 7)
	fmt.Print(ans)
	return
}
