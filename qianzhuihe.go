package main

import "fmt"

// 通过前缀和求任意区间的区间和问题：O(n)
func main() {
	var n, m int
	fmt.Scan(&n, &m) //输入元素个数n和查询个数m
	var a = make([]int, n+1)
	// 读取元素
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	// 初始化前缀和
	var s = make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}

	for m > 0 {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Printf("%d\n", s[r]-s[l-1])
		m--
	}
}

// 二维前缀和
//func main() {
//	var n, m, q int
//	fmt.Scan(&n, &m, &q) // 行数n，列数m，查询个数q
//	// 读取元素
//	var a = make([][]int, n+1)
//	for i := 1; i <= n; i++ {
//		a[i] = make([]int, m+1)
//		for j := 1; j <= m; j++ {
//			fmt.Scan(&a[i][j])
//		}
//	}
//	// 初始化前缀和
//	var s = make([][]int, n+1)
//	for i := 1; i <= n; i++ {
//		s[i] = make([]int, m+1)
//		for j := 1; j <= m; j++ {
//			s[i][j] = s[i][j-1] + s[i-1][j] - s[i-1][j-1]
//		}
//	}
//	// 开始查询
//	for q > 0 {
//		var x1, y1, x2, y2 int
//		fmt.Scan(&x1, &y1, &x2, &y2)
//		fmt.Printf("%d\n", s[x2][y2]-s[x1-1][y2]-s[x2][y1-1]+s[x1-1][y1-1])
//		q--
//	}
//}
