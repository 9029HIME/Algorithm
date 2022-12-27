package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addOne := false
	var result *ListNode
	var rPointer *ListNode
	aPointer := l1
	bPointer := l2
	for aPointer != nil && bPointer != nil {
		a := aPointer.Val
		b := bPointer.Val
		c := a + b

		if addOne {
			c++
			addOne = false
		}
		d := c - 10
		if d >= 0 {
			addOne = true
			c = d
		}
		if result == nil {
			result = new(ListNode)
			result.Val = c
			rPointer = result
		} else {
			next := new(ListNode)
			next.Val = c
			rPointer.Next = next
			rPointer = rPointer.Next
		}

		aPointer = aPointer.Next
		bPointer = bPointer.Next
	}

	// 其中一个链表遍历完，将另一个链表直接追加到result
	if aPointer == nil {
		rPointer.Next = bPointer
	}
	if bPointer == nil {
		rPointer.Next = aPointer
	}

	// 梳理另一个链表的值，是否需要进1
	for rPointer.Next != nil {
		e := rPointer.Next.Val
		if addOne {
			e++
			addOne = false
		}
		f := e - 10
		if f >= 0 {
			addOne = true
			e = f
		}
		rPointer.Next.Val = e
		rPointer = rPointer.Next
	}

	// 如果仍需要进1，则补充一个新节点，值为1
	if addOne {
		rPointer.Next = &ListNode{
			Val: 1,
		}
	}

	return result
}
