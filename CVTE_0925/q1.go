package main

import "fmt"

type ListNode struct {
	Val  byte
	Next *ListNode
}

// 删除无序链表的重复元素，仅保留第一个
// 输入：'a' -> 'b' -> 'c' -> 'a' -> 'b'
// 输出：'a' -> 'b' -> 'c'
func main() {
	node5 := &ListNode{Val: 'a', Next: nil}
	node4 := &ListNode{Val: 'b', Next: node5}
	node3 := &ListNode{Val: 'c', Next: node4}
	node2 := &ListNode{Val: 'b', Next: node3}
	node1 := &ListNode{Val: 'a', Next: node2}
	deleteDuplicates(node1)
	cur := node1
	for cur != nil {
		fmt.Printf("%c", cur.Val)
		fmt.Print(" ")
		cur = cur.Next
	}
	return
}

// 思路：使用map记录已经出现过的节点值，如果当前值已出现就直接跳过（断开前序节点到当前节点的连接，pre直接链接到cur的Next）
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hashMap := make(map[byte]int)
	hashMap[head.Val] = 1
	pre, cur := head, head.Next
	for cur != nil {
		count, ok := hashMap[cur.Val]
		if !ok || count == 0 {
			// 当前cur不是重复节点，则添加到map
			hashMap[cur.Val]++
			pre, cur = cur, cur.Next
		} else {
			// 当前cur为重复节点，断开连接
			cur = cur.Next
			pre.Next = cur
		}
	}
	return head
}
