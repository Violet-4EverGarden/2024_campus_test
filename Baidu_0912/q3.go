package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	// 读取元素
	var a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	// 初始化差分数组，下标从1开始，由于需要r+1因此长度为n+2
	var A = make([]int, n+2) // 由于初始愤怒值为0，因此差分数组各元素初始值也为0

	// 开始修改
	var res int
	for k := 1; k <= m; k++ {
		var l, r int
		fmt.Scan(&l, &r)
		A[l] += 1
		A[r+1] -= 1

		var ok bool
		var B = make([]int, n+1) // 计算差分数组前缀和，还原序列
		for i := 1; i <= n; i++ {
			B[i] = B[i-1] + A[i]
			if B[i] > a[i] {
				ok = true
			}
		}
		if ok {
			res = k - 1
			break
		} else {
			res += 1
		}
	}
	fmt.Print(res)
}
