package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	// fmt.Print(t)
	res := make([]string, 0)
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)
		if len(s) < 2 || len(s)%2 != 0 { // s的长度小于2或含有奇数个字符
			res = append(res, "NO")
			continue
		}
		for len(s) != 0 {
			index := strings.Index(s, "ab") // 在s中招"ab"第一次出现的位置
			if index == -1 {
				// fmt.Print("no")
				res = append(res, "NO")
				break
			}
			s = s[:index] + s[index+2:] // 去掉"ab"
			// fmt.Print(s)
		}
		if len(s) == 0 {
			res = append(res, "YES")
		}
	}

	fmt.Print(res[0])
	for i := 1; i < len(res); i++ {
		fmt.Print("\n")
		fmt.Print(res[i])
	}
	return
}
