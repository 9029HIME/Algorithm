package main

import "fmt"

/**
链表与节点的定义
*/
type Node struct {
	next *Node
	val  int
}

type SinglyLinkedList struct {
	head *Node
}

/**
初始化有序链表（切片得是有序的）
*/
func InitSinglyLinkedList(values []int) *SinglyLinkedList {
	list := new(SinglyLinkedList)
	list.head = new(Node)

	node := list.head
	for _, v := range values {
		n := new(Node)
		n.val = v
		node.next = n
		node = node.next
	}

	return list
}

/**
获取Result链表
*/
func NewSinglyLinkedList() *SinglyLinkedList {
	list := new(SinglyLinkedList)
	list.head = new(Node)
	return list
}

/**
打印链表结构
*/
func (this *SinglyLinkedList) introduce() {
	result := "head-->"
	node := this.head.next
	for node != nil {
		val := fmt.Sprintf("%v-->", node.val)
		result = result + val
		node = node.next
	}
	result = result[0 : len(result)-3]
	fmt.Println(result)
	fmt.Println()

}
