package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := &ListNode{
		Val:  1,
		Next: nil,
	}

	b := &ListNode{
		Val:  4,
		Next: nil,
	}

	c := &ListNode{
		Val:  3,
		Next: nil,
	}

	d := &ListNode{
		Val:  2,
		Next: nil,
	}

	e := &ListNode{
		Val:  5,
		Next: nil,
	}

	f := &ListNode{
		Val:  2,
		Next: nil,
	}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

	result := partition(a, 3)
	resultVal := ""
	for result != nil {
		val := result.Val
		resultVal = resultVal + strconv.Itoa(val) + "-->"
		result = result.Next
	}
	fmt.Println(resultVal)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left := new(ListNode)
	right := new(ListNode)
	lp := left
	rp := right

	for head != nil {
		nextLoop := head.Next
		val := head.Val
		if val < x {
			next := lp.Next
			lp.Next = head
			head.Next = next

			lp = lp.Next
		} else {
			next := rp.Next
			rp.Next = head
			head.Next = next

			rp = rp.Next
		}

		head = nextLoop
	}

	if right.Next != nil {
		lp.Next = right.Next
	}

	return left.Next
}
