package main

import "testing"

func TestListIntroduce(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	list := InitSinglyLinkedList(slice)
	list.introduce()
}

func Test2Lists(t *testing.T) {
	values1 := []int{1, 3}
	values2 := []int{2, 4, 6, 8}

	l1 := InitSinglyLinkedList(values1)
	l2 := InitSinglyLinkedList(values2)

	result := Merge2SortedLists(l1, l2)
	result.introduce()
}

/**
测试节点数据在最小堆的作用
*/
func TestSmallNodeHeap(t *testing.T) {
	values2 := []int{9, 10, 11}
	queue := InitSmallHeapQueue(len(values2))
	l2 := InitSinglyLinkedList(values2)
	node := l2.head.next
	for node != nil {
		queue.put(node)
		node = node.next
	}

	queue.introduce()

	queue.pop()
	queue.introduce()

}

func TestKLists(t *testing.T) {
	values1 := []int{1, 3, 5, 7, 9}
	values2 := []int{2, 4, 6, 8}
	values3 := []int{11, 13, 15, 17, 19}
	values4 := []int{10, 12, 14, 16, 18}

	l1 := InitSinglyLinkedList(values1)
	l2 := InitSinglyLinkedList(values2)
	l3 := InitSinglyLinkedList(values3)
	l4 := InitSinglyLinkedList(values4)

	lists := []*SinglyLinkedList{l1, l2, l3, l4}

	result := MergeKSortedLists(lists)
	result.introduce()
}
