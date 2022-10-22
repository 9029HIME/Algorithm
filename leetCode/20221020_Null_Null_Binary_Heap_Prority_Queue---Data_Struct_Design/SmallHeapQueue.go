package _0221020_Null_Null_Binary_Heap_Prority_Queue___Data_Struct_Design

import "math"

/**
基于最小堆的优先级队列
*/
type SmallHeapQueue struct {
	BinaryHeap
}

func InitSmallHeapQueue(capacity int) *SmallHeapQueue {
	p := new(SmallHeapQueue)
	// 记得冗余index=0
	p.data = make([]int, capacity+1)
	p.capacity = capacity
	p.length = 0
	return p
}

/**
下沉指定下标的节点（最小堆，将大值下沉）
*/
func (this *SmallHeapQueue) sink(i int) {
	// 既然要下沉，那就找最大
	lI := this.leftSon(i)
	rI := this.rightSon(i)

	if lI > this.length && rI > this.length {
		// 已经是叶子节点了
		return
	}
	left := this.data[lI]
	right := this.data[rI]

	if left == 0 {
		left = math.MaxInt
	}
	if right == 0 {
		right = math.MaxInt
	}

	var max int
	var maxI int
	if left <= right {
		max = left
		maxI = lI
	} else {
		max = right
		maxI = rI
	}
	// 得和最大的子节点进行比较，如果小于它，就和它swap
	val := this.data[i]
	if val > max {
		this.swap(i, maxI)
	} else {
		// 出口
		return
	}
	this.sink(maxI)
}

/**
上升指定下标的节点（最小堆，将小值上升）
*/
func (this *SmallHeapQueue) rise(i int) {
	fI := this.father(i)
	if fI == 0 {
		// 说明已经在堆顶了
		return
	}
	val := this.data[i]
	father := this.data[fI]
	if father > val {
		this.swap(i, fI)
	} else {
		// 递归出口
		return
	}
	// 此时节点就在father所在的下标了
	this.rise(fI)
}

func (this *SmallHeapQueue) put(val int) {
	this.BinaryHeap.put(val, this)
}

func (this *SmallHeapQueue) del(i int) {
	this.BinaryHeap.del(i, this)
}
