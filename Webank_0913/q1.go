package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	var mapInfo = make(map[int]int)
	for i := 1; i <= n; i++ {
		ind, ok := mapInfo[a[i]]
		if !ok {
			mapInfo[a[i]] = i
		} else {
			if ind < i {
				mapInfo[a[i]] = i
			}
		}
	}
	// fmt.Println(mapInfo)
	type thing struct {
		color int
		index int
	}
	var things []thing
	for k, v := range mapInfo {
		things = append(things, thing{color: k, index: v})
	}
	sort.Slice(things, func(i, j int) bool {
		return things[i].index < things[j].index // 按index升序
	})
	fmt.Print(things[0].color)
	for i := 1; i < len(things); i++ {
		fmt.Print(" ")
		fmt.Print(things[i].color)
	}
}

/*
8
1 2 1 3 4 2 4 4
*/
