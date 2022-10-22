package main

func MergeKSortedLists(lists []*SinglyLinkedList) *SinglyLinkedList {
	heap := InitSmallHeapQueue(len(lists))
	result := NewSinglyLinkedList()
	tail := result.head

	for _, v := range lists {
		//直接第一个节点放到最小堆里
		first := v.head.next
		heap.put(first)
	}

	pop := heap.pop()
	var nextAdd *Node

	for pop != nil {
		tail.next = pop
		tail = tail.next
		nextAdd = pop.next
		if nextAdd != nil {
			heap.put(nextAdd)
		}
		pop = heap.pop()
	}
	return result
}
