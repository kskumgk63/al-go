package merged_two_sorted_lists

/*
Feedback
- "container/heap" is not necessary
- should have implemented m]methods like heap.Push()
*/
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var b, l *ListNode
	if l1.Val <= l2.Val {
		l = l1
		l1 = l1.Next
	} else {
		l = l2
		l2 = l2.Next
	}
	b = l
	for ; l1 != nil; l1 = l1.Next {
		for ; l2 != nil && l2.Val < l1.Val; l2 = l2.Next {
			l.Next = l2
			l = l.Next
		}
		l.Next = l1
		l = l.Next
	}
	if l2 != nil {
		l.Next = l2
	} else {
		l.Next = nil
	}
	return b
}
