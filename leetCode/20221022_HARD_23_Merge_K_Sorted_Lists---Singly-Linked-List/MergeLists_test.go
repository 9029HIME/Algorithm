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
