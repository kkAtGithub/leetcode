package Q0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}
	p, p1, p2 := head, list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return head.Next
}
