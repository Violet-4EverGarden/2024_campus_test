package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	// fmt.Println(s)
	for s[0] == ' ' && len(s) > 0 {
		s = s[1:]
	}
	if s == "" {
		fmt.Printf("%d", 0)
		return
	}
	if s[0] != '+' && s[0] != '-' && s[0] < '0' && s[0] > '9' {
		fmt.Printf("%d", 0)
		return
	}

	tag := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		tag = -1
		s = s[1:]
	}
	// 判断字符串中间是否有非数字字符
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			fmt.Print(0)
			return
		}
	}

	num, _ := strconv.Atoi(s)
	num = tag * num

	// 判断是否超过32为整数的最大或最小
	if num > math.MaxInt32 || num < math.MinInt32 {
		if num > math.MaxInt32 {
			fmt.Print(math.MaxInt32)
			return
		}
		if num < math.MinInt32 {
			fmt.Print(math.MinInt32)
			return
		}
	}
	fmt.Print(num)
	return
}
