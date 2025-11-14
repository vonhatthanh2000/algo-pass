package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	fmt.Println(newHead)
	if head.Next != nil {
		newHead = reverseListRecursive(head.Next)
		fmt.Println(1)
		fmt.Println(head.Next.Val)
		fmt.Println(head.Next.Next)
		fmt.Println(head)
		head.Next.Next = head
	}
	head.Next = nil

	return newHead
}

// func main() {

// 	node6 := &ListNode{
// 		Val:  6,
// 		Next: nil,
// 	}
// 	node5 := &ListNode{
// 		Val:  5,
// 		Next: node6,
// 	}
// 	node4 := &ListNode{
// 		Val:  4,
// 		Next: node5,
// 	}
// 	node3 := &ListNode{
// 		Val:  3,
// 		Next: node4,
// 	}
// 	node2 := &ListNode{
// 		Val:  2,
// 		Next: node3,
// 	}
// 	node1 := &ListNode{
// 		Val:  1,
// 		Next: node2,
// 	}

// 	reverseListRecursive(node1)

// }
