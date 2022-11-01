package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow := head
	fast := head

	if slow == nil {
		goto exit
	}

	for fast != nil {
		sV := slow.Val
		fV := fast.Val
		if sV != fV {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	// 避免最后N个数据重复的情况，比如用例1→1→2→3→3，如果不加这个，链表结果是1→2→3→3
	slow.Next = nil

exit:
	return head
}
