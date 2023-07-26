package Q0023

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	_lists := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		if list == nil {
			continue
		}
		_lists = append(_lists, list)
	}
	length := len(_lists)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(_lists, i, length)
	}
	dummy := &ListNode{}
	p := dummy

	for nilCount := 0; nilCount < length; {
		p.Next = _lists[0]
		p = p.Next
		_lists[0] = _lists[0].Next
		if _lists[0] == nil {
			nilCount++
			_lists[0], _lists[length-nilCount] = _lists[length-nilCount], _lists[0]
		}
		adjustHeap(_lists, 0, length-nilCount)
	}

	return dummy.Next
}

func adjustHeap(heap []*ListNode, i, length int) {
	temp := heap[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && heap[k].Val > heap[k+1].Val {
			k++
		}
		if temp.Val > heap[k].Val {
			heap[i] = heap[k]
			i = k
		}
	}
	heap[i] = temp
}
