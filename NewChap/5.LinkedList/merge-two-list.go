package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next
}

func main() {
	node1_6 := &ListNode{
		Val:  6,
		Next: nil,
	}
	node1_5 := &ListNode{
		Val:  5,
		Next: node1_6,
	}
	node1_2 := &ListNode{
		Val:  2,
		Next: node1_5,
	}
	node1_1 := &ListNode{
		Val:  1,
		Next: node1_2,
	}

	node2_8 := &ListNode{
		Val:  8,
		Next: nil,
	}
	node2_6 := &ListNode{
		Val:  6,
		Next: node2_8,
	}
	node2_4 := &ListNode{
		Val:  2,
		Next: node2_6,
	}
	node2_1 := &ListNode{
		Val:  1,
		Next: node2_4,
	}

	res := mergeTwoLists(node1_1, node2_1)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}
