package main

import (
	"fmt"
	"sort"
)

// 扑克牌顺子：力扣、牛客原题改版：A(即1)可以接在13后面，即10、11、12、13、1也是顺子
func main() {
	var a = make([]int, 5)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a) // 先排序
	var super int
	var gap int
	for i := 0; i < len(a)-1; i++ { // 注意这里是小于len-1
		// 大小王做癞子
		if a[i] == 0 {
			super++
			continue
		}
		if a[i] == a[i+1] { // 有对子，直接返回
			fmt.Print(false)
			return
		}
		gap += a[i+1] - a[i] - 1 // 计算相邻数字间的间隔总和
	}
	if super >= gap {
		fmt.Print(true)
		return
	}
	fmt.Print(false)
}

/*
0 0 1 10 11
true

1 2 3 4 5
true
*/
