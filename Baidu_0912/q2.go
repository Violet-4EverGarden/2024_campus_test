package main

import "fmt"

// 讨厌鬼的帖子组合
// 思路：贪心
func main() {
	var n int
	var a [100010]int
	var b [100010]int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&b[i])
	}
	var ans1, ans2 int64
	for i := 1; i <= n; i++ {
		ans1 += int64(max(0, a[i]-b[i]))
		ans2 += int64(max(0, b[i]-a[i]))
	}
	fmt.Print(maxInt64(ans1, ans2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
