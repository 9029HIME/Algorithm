package main

import "fmt"

type LinkedHashSet struct {
	hashMap map[int]*Node
	head    *Node
	tail    *Node
}

type Node struct {
	pred *Node
	next *Node
	// 接口是不需要指针的，指定具体值时，必须是implement结构体的指针
	data HashEntity
}

/**
类似Java的hashCode
*/
type HashEntity interface {
	hashCode() int
}

func InitLinkedHashSet(capacity int) *LinkedHashSet {
	set := &LinkedHashSet{
		hashMap: make(map[int]*Node, capacity),
		head:    new(Node),
		tail:    new(Node),
	}
	set.head.next = set.tail
	set.tail.pred = set.head
	return set
}

func (this *LinkedHashSet) put(data HashEntity) {
	hashCode := data.hashCode()
	if eN, ok := this.hashMap[hashCode]; ok {
		eN.data = data
	} else {
		n := &Node{
			data: data,
		}
		this.hashMap[hashCode] = n

		this.addNode(n)
	}
}

func (this *LinkedHashSet) contains(data HashEntity) bool {
	hashCode := data.hashCode()
	_, ok := this.hashMap[hashCode]
	return ok
}

func (this *LinkedHashSet) remove(data HashEntity) {
	hashCode := data.hashCode()
	// 可别忘了，链表上的节点也要删除
	node, ok := this.hashMap[hashCode]
	if ok {
		this.delNode(node)
		delete(this.hashMap, hashCode)
	}
}

func (this *LinkedHashSet) isEmpty() bool {
	return len(this.hashMap) == 0
}

/**
将节点从链表中删除
*/
func (this *LinkedHashSet) delNode(node *Node) {
	if node == this.head || node == this.tail {
		panic("无法删除头节点或尾节点")
	}

	pred := node.pred
	next := node.next
	pred.next = next
	next.pred = pred
	node.pred = nil
	node.next = nil
}

/**
将节点添加到链尾
*/
func (this *LinkedHashSet) addNode(node *Node) {
	pred := this.tail.pred
	pred.next = node
	node.pred = pred
	node.next = this.tail
	this.tail.pred = node
}

/**
将数据都打印出来
*/
func (this *LinkedHashSet) introduce() {
	hashMap := this.hashMap
	for k, v := range hashMap {
		entry := fmt.Sprintf("[%s:%s]---", k, v.data)
		fmt.Print(entry)
	}
	fmt.Println()
	node := this.head
	fmt.Print("head→")
	for node.next != nil && node.next != this.tail {
		node = node.next
		msg := fmt.Sprintf("<%s>→", node.data)
		fmt.Print(msg)
	}
	fmt.Print("tail")
	fmt.Println()
}
