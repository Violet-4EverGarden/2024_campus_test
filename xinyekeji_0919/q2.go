package main

import (
	"container/list"
	"fmt"
)

// 有效的括号
func main() {
	var s string
	fmt.Scan(&s)
	if isValid(s) == true {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}

func isValid(s string) bool {
	stack := list.New()
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.PushBack(s[i])
		} else {
			if stack.Len() == 0 {
				return false
			}
			top := stack.Remove(stack.Back()).(byte)
			if top != mapping[s[i]] {
				return false
			}
		}
	}

	if stack.Len() != 0 { // 避免s只有一个正括号的情况，如"["
		return false
	}
	return true
}
