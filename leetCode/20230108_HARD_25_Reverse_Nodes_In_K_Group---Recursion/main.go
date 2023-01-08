package main

import "fmt"

func main() {
	head := BuildSinglyLinedList([]int{1, 2, 3, 4, 5})
	newHead := reverseKGroup(head, 2)
	fmt.Println(newHead)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return Swap(head, k)
}

/**
A → B → C , k=3
head = A
after swap: C → B → A
return C

A → B , k=3
head = A
after swap: A → B
return A
*/
func Swap(head *ListNode, k int) *ListNode {
	pointer := head
	/*
		校验包括自己，后面一共有K个节点，如果不是，直接返回head，保证K个一组
		A → B → C → D
		↑
		A → B → C → D
			↑
		A → B → C → D
				↑
	*/
	for i := 0; i < k; i++ {
		if pointer == nil {
			return head
		}
		pointer = pointer.Next
	}
	/*
		A → B → C → D → E → F
					↑
	*/
	nextHead := pointer
	/*
		A → B → C → F → E → D
					↑
	*/
	nextHead = Swap(nextHead, k)

	/*
		下面的流程跟drawio图一致
	*/
	pointer = head
	previous := nextHead
	for i := 0; i < k; i++ {
		next := pointer.Next
		pointer.Next = previous
		previous = pointer
		pointer = next
	}
	return previous
}

/**
通过切片初始化链表
*/
func BuildSinglyLinedList(nums []int) *ListNode {
	var head *ListNode
	var pointer *ListNode
	for _, num := range nums {
		if head == nil {
			head = new(ListNode)
			head.Val = num
			pointer = head
		} else {
			next := new(ListNode)
			next.Val = num
			pointer.Next = next
			pointer = pointer.Next
		}
	}
	return head
}
