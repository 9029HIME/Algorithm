package main

import "fmt"

func main() {
	head := new(ListNode)

	one := &ListNode{
		Val: 1,
	}
	//
	//two := &ListNode{
	//	Val: 2,
	//}
	//
	//three := &ListNode{
	//	Val: 3,
	//}
	//
	//four := &ListNode{
	//	Val: 4,
	//}
	//
	//five := &ListNode{
	//	Val: 5,
	//}
	head.Next = one
	//one.Next = two
	//two.Next = three
	//three.Next = four
	//four.Next = five

	head.introduce()
	removeNthFromEnd(head, 1).introduce()
}

// 第一个节点是哑巴节点 处理情况1：倒数第k个节点就是头节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) introduce() {
	node := this.Next
	result := "head->"
	for node != nil {
		val := fmt.Sprintf("%v->", node.Val)
		result = result + val
		node = node.Next
	}
	result = result[0 : len(result)-2]
	fmt.Println(result)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 处理情况3：链表为空
	if head.Next == nil {
		return head
	}
	front := head
	back := head
	// n-1：倒数第n个，+1是因为要找到第n个的前置节点
	for i := 0; i < (n-1)+1; i++ {
		front = front.Next
		// 处理情况2：链表没有倒数第k个节点
		if front == nil {
			return head
		}
	}

	for front.Next != nil {
		back = back.Next
		front = front.Next
	}

	// 此时back就是链表的倒数第n个节点的前置节点
	nNode := back.Next
	next := nNode.Next
	nNode.Next = nil
	back.Next = next

	return head
}
