package main

import (
	"fmt"
	"math"
)

var res int

// 分割回文串的变体：求的是将s分割为回文子串数组的最少分割次数（其实就是len最小的那个回文子串数组）
func main() {
	var s string
	fmt.Scan(&s)
	res = math.MaxInt32 // 注意res一开始应该取int最大值
	path := make([]string, 0)
	backtracking(s, 0, path)
	fmt.Print(res)
}

func backtracking(s string, startInd int, path []string) {
	// 终止条件
	if startInd == len(s) {
		if len(path)-1 < res { // len(path)-1表示：分割次数应该为回文子串个数-1
			res = len(path) - 1
		}
		return
	}

	// 单层递归-回溯逻辑
	for i := startInd; i < len(s); i++ {
		str := s[startInd : i+1] // 截取当前字符串
		if !isPalindrom(str) {
			continue
		}
		path = append(path, str)   // 处理当前节点
		backtracking(s, i+1, path) // 递归进入下一层搜索
		path = path[:len(path)-1]  // 回溯撤销当前处理
	}
}

func isPalindrom(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
