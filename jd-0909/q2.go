package main

import "fmt"

// 类似LC.746 使用最小花费爬楼梯
// 只是costs[i]表示爬到第i层需要消耗的体力值（LC.746是花费这么多才能向上爬）
/*
输入：
5
0 3 2 1 0
输出：
2
*/
func main() {
	var n int
	fmt.Scan(&n)
	var costs = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&costs[i])
	}

	if n <= 3 {
		fmt.Print(costs[n-1])
		return
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = costs[1]
	dp[2] = costs[2]
	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1]+costs[i], dp[i-2]+costs[i])
	}
	fmt.Print(dp[n-1])
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
