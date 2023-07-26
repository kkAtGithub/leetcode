package Q0086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var dummy1 = &ListNode{}
	var dummy2 = &ListNode{}
	p1, p2 := dummy1, dummy2
	for ; head != nil; head = head.Next {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}
