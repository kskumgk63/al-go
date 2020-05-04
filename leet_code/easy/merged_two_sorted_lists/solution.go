package merged_two_sorted_lists

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type intHeap []int

func (h intHeap) Len() int { return len(h) }

func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	if v, ok := x.(int); ok {
		*h = append(*h, v)
	}
}

func (h *intHeap) Pop() interface{} {
	if h.Len() < 1 {
		return nil
	}
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func (l *ListNode) setFunc(f func() (int, bool)) {
	i, ok := f()
	if !ok {
		l = nil
		return
	}
	l.Val = i
	if l.Next == nil {
		l.Next = &ListNode{}
	}
	l.Next.setFunc(f)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &intHeap{}
	heap.Init(h)

	for list1, list2 := l1, l2; list1 != nil && list2 != nil; list1, list2 = list1.Next, list2.Next {
		h.Push(list1.Val)
		h.Push(list2.Val)
	}

	ret := &ListNode{}
	ret.setFunc(func() (int, bool) {
		i, ok := h.Pop().(int)
		return i, ok
	})

	return ret
}
