package main

/**
1.题目
合并两个有序链表

2.描述
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

3.示例
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
输入：l1 = [], l2 = []
输出：[]
输入：l1 = [], l2 = [0]
输出：[0]

4.核心思想
既然两个都是有序的，那就充分发挥它们有序的特性。使用双指针比较，每次都取两个队列最小的值进行比较，最小的那个值移出来，放到结果链表里

↓
1-2-3-4

7-8-9-10
↑

1最小

↓
2-3-4
1-
7-8-9-10
↑

2最小

↓
3-4
1-2
7-8-9-10
↑

*/
func main() {
	var first []int = nil
	var second []int = []int{0}

	firstHead := new(ListNode)
	secondHead := new(ListNode)

	for _, v := range first {
		if firstHead.Val == 0 {
			firstHead.Val = v
		} else {
			var latestNode *ListNode = firstHead
			next := new(ListNode)
			next.Val = v
			for latestNode.Next != nil {
				latestNode = latestNode.Next
			}
			latestNode.Next = next
		}
	}

	for _, v := range second {
		if secondHead.Val == 0 {
			secondHead.Val = v
		} else {
			var latestNode *ListNode = secondHead
			next := new(ListNode)
			next.Val = v
			for latestNode.Next != nil {
				latestNode = latestNode.Next
			}
			latestNode.Next = next
		}
	}

	lists := mergeTwoLists(nil, secondHead)
	print(lists)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pointer1 := list1
	pointer2 := list2
	var result *ListNode = &ListNode{}
	pointer3 := result

	for pointer1 != nil && pointer2 != nil {
		var smaller int
		if pointer1.Val < pointer2.Val {
			smaller = pointer1.Val
			pointer1 = pointer1.Next
		} else {
			smaller = pointer2.Val
			pointer2 = pointer2.Next
		}
		var smallerNode *ListNode = &ListNode{
			Val: smaller,
		}
		pointer3.Next = smallerNode
		pointer3 = pointer3.Next
	}

	if pointer1 == nil {
		pointer3.Next = pointer2
	}

	if pointer2 == nil {
		pointer3.Next = pointer1
	}

	result = result.Next

	return result
}
