package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type aa struct {
	num   int
	count int
}

// 前k个高频元素：先用map统计出现次数，再转为struct切片排序，再输出前k个
func main() {
	var s string
	fmt.Scan(&s)
	s_slice := strings.Split(s, ",")
	var k int
	fmt.Scan(&k)

	hashMap := make(map[int]int)
	for i := 0; i < len(s_slice); i++ {
		num, _ := strconv.Atoi(s_slice[i])
		if _, ok := hashMap[num]; !ok {
			hashMap[num] = 1
		} else {
			hashMap[num]++
		}
	}

	var res = make([]aa, 0)
	for k, v := range hashMap {
		var tmp aa
		tmp.num = k
		tmp.count = v
		res = append(res, tmp)
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].count > res[i].count {
			return true
		}
		return false
	})

	fmt.Print(res[0].num)
	for i := 1; i < k; i++ {
		fmt.Print(",")
		fmt.Print(res[i].num)
	}
	return
}
