package main

import "fmt"

type BinaryHeap struct {
	data     []*Node
	capacity int
	length   int
}

type RiseSink interface {
	rise(i int)
	sink(i int)
}

func (this *BinaryHeap) leftSon(i int) int {
	return i * 2
}

func (this *BinaryHeap) rightSon(i int) int {
	return (i * 2) + 1
}

func (this *BinaryHeap) father(i int) int {
	return i / 2
}

/**
实际交换节点
*/
func (this *BinaryHeap) swap(from int, to int) {
	// 在只有一个堆顶节点，进行删除的时候，有可能会自己与自己换
	if from == to {
		return
	}

	fromValue := this.data[from]
	toValue := this.data[to]
	this.data[to] = fromValue
	this.data[from] = toValue
}

func (this *BinaryHeap) introduce() {
	result := "0->"
	for i, datum := range this.data {
		if i == 0 {
			continue
		}
		if datum == nil {
			break
		}
		val := fmt.Sprintf("%v->", datum.val)
		result = result + val
	}
	fmt.Println(result)
}

func (this *BinaryHeap) put(val *Node, rs RiseSink) {
	if this.length == this.capacity {
		panic("优先级队列已满")
	} else {
		this.data[this.length+1] = val
		rs.rise(this.length + 1)
		this.length++
	}
}

/**
删除下标=i的节点
*/
func (this *BinaryHeap) del(i int, rs RiseSink) {

	if i > this.length {
		return
	}

	to := this.length
	this.swap(i, to)

	this.data[to] = nil
	this.length--

	rs.sink(i)
}

/**
删除堆顶数据，并返回它
*/
func (this *BinaryHeap) pop(rs RiseSink) *Node {
	node := this.data[1]
	this.del(1, rs)
	return node
}
