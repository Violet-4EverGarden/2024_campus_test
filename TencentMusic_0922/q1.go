package main

//import "fmt"

// 小红的数位删除：逆天题
// 一个正整数n，每次操作可以删除一位数，最少操作多少次可以使其变为5的倍数？注意所有数删除完后为0，0为5的倍数，所以必定有解
func fun(n int) int {
	// write code here
	res := 0
	for n%5 != 0 {
		n = n / 10
		res++
	}
	return res
}
