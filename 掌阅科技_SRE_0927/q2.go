package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
输入:
1, 3, 15, 11, 2
23, 127, 235, 19, 8
*/
// 此题关键在于对输入的处理优点恶心，其它就暴力可解
// 注意fmt.Scanln在遇到空格后也会结束读取！！！所以当一行字符串中有空格时最好使用bufio来读取整行！
func main() {
	var s1, s2 string
	reader := bufio.NewReader(os.Stdin)
	s1, _ = reader.ReadString('\n')
	s1 = strings.TrimSpace(s1)
	s2, _ = reader.ReadString('\n')
	s2 = strings.TrimSpace(s2)

	s1_slice := strings.Split(s1, ", ")
	s2_slice := strings.Split(s2, ", ")
	nums1 := trans(s1_slice)
	nums2 := trans(s2_slice)

	min := math.MaxInt32
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			tmp := Abs(nums1[i] - nums2[j])
			if tmp < min {
				min = tmp
			}
		}
	}
	fmt.Print(min)
}

func trans(s []string) []int {
	var nums = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		nums[i], _ = strconv.Atoi(s[i])
	}
	return nums
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
