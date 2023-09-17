package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node7 := &TreeNode{Val: 0, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 0, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 1, Left: node6, Right: node7}
	node2 := &TreeNode{Val: 1, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 0, Left: node2, Right: node3}
	fmt.Print(pathNumber(node1))
}

func pathNumber(root *TreeNode) int {
	// write code here
	count := 0
	res := 0
	// path := make([]int, 0)
	var Order func(node *TreeNode)
	Order = func(node *TreeNode) {
		// 递归终止条件
		if node == nil {
			return
		}
		// 中
		if node.Val == 1 {
			count++
		} else {
			count--
		}
		// path = append(path, node.Val)
		if node.Left != nil || node.Right != nil {
			// 左
			Order(node.Left)
			// 右
			Order(node.Right)
		} else {
			if count == 1 {
				res += 1
			}
			// fmt.Println(path)
		}
		// 回溯撤销处理
		if node.Val == 1 {
			count--
		} else {
			count++
		}
		// path = path[:len(path)-1]
	}
	Order(root)
	return res
}
