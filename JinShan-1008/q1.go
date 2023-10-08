package main

import "fmt"

// 构造合法单词：一个合法单词不能有连续的两个辅音字母相连，元音字母包括：a、e、i、o、u
func main() {
	var s string
	fmt.Scan(&s)
	t := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			t++
		}
	}
	fmt.Printf("%d", min(t*2+1, len(s))) // 一个元音可以带一个辅音，最后再+1，因此最多为2*t+1；取其与s长度的较小值
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
