package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) int {
	// Implement your solution here
	var ans = 0
	var dfs func(node *Tree) int
	dfs = func(node *Tree) int {
		a := 0
		if node.L != nil {
			a = dfs(node.L)
		}
		b := 0
		if node.R != nil {
			b = dfs(node.R)
		}
		cur := min(a, b)*2 + 1
		ans = max(ans, cur)
		return cur
	}
	dfs(T) // 虽然dfs带返回值，但可以不用管，因为它只是用来更新a、b，而最终结果ans是全局变量，在dfs过程中被更新。
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
(1, (2, None, (4, None, None)), (3, (5, (7, None, None), (8, None, None)), (6, (9, None, None), (10, (11, None, None), None))))

7
*/
