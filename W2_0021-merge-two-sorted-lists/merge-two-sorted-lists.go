/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	next := getSmallerNode(list1, list2)
	for next != nil {
		cur.Next = next
		cur = cur.Next

		list1, list2 = advanceLists(list1, list2, next)
		next = getSmallerNode(list1, list2)
	}

	return dummy.Next
}

func getSmallerNode(node1 *ListNode, node2 *ListNode) *ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.Val < node2.Val {
		return node1
	}
	return node2
}

func advanceLists(node1 *ListNode, node2 *ListNode, chosen *ListNode) (*ListNode, *ListNode) {
	if chosen == node1 {
		return node1.Next, node2
	}
	return node1, node2.Next
}