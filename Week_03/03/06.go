/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/
package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint1(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append([]int{head.Val}, res...)
		head = head.Next
	}
	return res
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint(head.Next), head.Val)
}
