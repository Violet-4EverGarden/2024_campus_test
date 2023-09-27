package main

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param task char字符型一维数组 任务列表
 * @param n int整型 执行次数
 * @return int整型
 */

type aa struct {
	task  byte
	count int
}

// 任务调度器：max((a-1)*(n+1)+x,t)，a是最多的同种任务的个数，n是间隔，t是总任务数，x是数量最多的任务的种数

func leastInterval(task []byte, n int) int {
	// write code here
	if n == 0 {
		return len(task)
	}

	hashMap := make(map[byte]int)
	for i := 0; i < len(task); i++ {
		if _, ok := hashMap[task[i]]; !ok {
			hashMap[task[i]] = 1
		} else {
			hashMap[task[i]]++
		}
	}
	res := make([]aa, 0)
	for k, v := range hashMap {
		var tmp aa
		tmp.task = k
		tmp.count = v
		res = append(res, tmp)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].count > res[j].count
	})

	x := 1
	maxTaskNum := res[0].count
	for i := 1; i < len(res); i++ {
		if res[i].count == maxTaskNum {
			x++
		} else {
			break
		}
	}
	result := max((maxTaskNum-1)*(n+1)+x, len(task))
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
