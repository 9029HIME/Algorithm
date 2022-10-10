package main

import "fmt"

/**
请你设计并实现一个满足LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	cache := InitLRUCache(4)
	node1 := &Node{
		key:   1,
		value: 100,
	}

	node2 := &Node{
		key:   2,
		value: 100,
	}

	node3 := &Node{
		key:   3,
		value: 100,
	}

	cache.push(node1)
	cache.push(node2)
	cache.push(node3)
	cache.introduce()

	cache.refresh(node1)
	fmt.Println()
	cache.introduce()

	cache.pop(node2)
	fmt.Println()
	cache.introduce()
}

type LRUCache struct {
	capacity      int
	linkedHashMap map[int]*Node
	head          *Node
	tail          *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	pred  *Node
}

func InitLRUCache(capacity int) *LRUCache {
	cache := &LRUCache{
		capacity:      capacity,
		linkedHashMap: make(map[int]*Node, capacity),
		head:          new(Node),
		tail:          new(Node),
	}
	cache.head.next = cache.tail
	cache.tail.pred = cache.head
	return cache
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {

	return 0
}

func (this *LRUCache) Put(key int, value int) {

}

/**
刷新节点的新鲜程度
*/
func (this *LRUCache) refresh(node *Node) {
	oldPred := node.pred
	oldNext := node.next

	oldPred.next = oldNext
	oldNext.pred = oldPred

	newPred := this.tail.pred
	newPred.next = node
	node.pred = newPred

	node.next = this.tail
	this.tail.pred = node
}

/**
往队列插入数据
*/
func (this *LRUCache) push(node *Node) {
	pred := this.tail.pred
	pred.next = node
	node.pred = pred

	node.next = this.tail
	this.tail.pred = node
}

/**
弹出这个节点，并返回key
*/
func (this *LRUCache) pop(node *Node) int {
	pred := node.pred
	next := node.next

	pred.next = next
	next.pred = pred

	node.pred = nil
	node.next = nil
	return node.key
}

// 实验用代码，与leetcode无关
func (this *LRUCache) introduce() {
	node := this.head
	for node != this.tail && node.next != nil {
		fmt.Print(node.key)
		fmt.Print("->")
		node = node.next
	}
	fmt.Print(this.tail.key)
}
