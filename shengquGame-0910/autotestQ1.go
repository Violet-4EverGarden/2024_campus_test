package main

import "fmt"

var (
	path []int
	res  int
)

func main() {
	path = make([]int, 0)
	res = 0
	var n int
	fmt.Scan(&n)
	length := 4
	backtracking(n, length, 0)
	fmt.Print(res)
}

func backtracking(n, length, startInd int) {
	if len(path) == length {
		res += 1
		fmt.Println(path)
		return
	}
	for i := startInd; i < n; i++ {
		path = append(path, i)
		backtracking(n, length, i)
		path = path[:len(path)-1]
	}
}
