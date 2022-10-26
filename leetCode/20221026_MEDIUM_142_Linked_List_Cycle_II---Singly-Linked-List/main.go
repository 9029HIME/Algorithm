package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := new(ListNode)
	one := &ListNode{
		Val: 1,
	}
	two := &ListNode{
		Val: 2,
	}
	three := &ListNode{
		Val: 3,
	}
	four := &ListNode{
		Val: 4,
	}
	five := &ListNode{
		Val: 5,
	}

	// 连接链表
	head.Next = one
	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five

	// 环的起点
	five.Next = two

	fmt.Println(detectCycle(head))
}

// 还是那句话：第一个节点是哑节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) String() string {
	return strconv.Itoa(this.Val)
}

func detectCycle(head *ListNode) *ListNode {
	var result *ListNode = nil
	var intersection *ListNode = nil

	slow := head
	fast := head

	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		// 说明没环，快指针已经到尾了
		if fast == nil {
			return nil
		}
		fast = fast.Next
		// 找到相交点了，别再循环了
		if slow == fast {
			intersection = slow
			break
		}
	}

	fmt.Println(intersection)

	// 既然有相交点，那就重置慢指针，同速执行
	if intersection != nil {
		slow = head
		// 能走到这一步，说明必定有环，就不用考虑nil情况了
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		result = slow
	}

	return result
}
