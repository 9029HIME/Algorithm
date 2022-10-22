package main

/**
合并两个有序链表的方法
*/
func Merge2SortedLists(l1 *SinglyLinkedList, l2 *SinglyLinkedList) *SinglyLinkedList {
	result := NewSinglyLinkedList()
	tail := result.head

	l := l1.head.next
	r := l2.head.next

	for l != nil && r != nil {
		lVal := l.val
		rVal := r.val
		var minNode *Node

		// 毕竟是单链表，出队操作比较复杂，逻辑出队就好了
		if lVal <= rVal {
			minNode = l
			l = l.next
		} else {
			minNode = r
			r = r.next
		}

		rN := new(Node)
		rN.val = minNode.val
		// 维护tail
		tail.next = rN
		tail = tail.next
	}

	// 走到这 说明已经有一个队列遍历完毕了
	if l == nil {
		tail.next = r
	} else {
		tail.next = l
	}

	return result
}
