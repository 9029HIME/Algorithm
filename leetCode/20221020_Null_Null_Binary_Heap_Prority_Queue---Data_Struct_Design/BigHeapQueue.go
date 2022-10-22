package _0221020_Null_Null_Binary_Heap_Prority_Queue___Data_Struct_Design

/**
基于最大堆的优先级队列
*/
type BigHeapQueue struct {
	BinaryHeap
}

func InitBigHeapQueue(capacity int) *BigHeapQueue {
	p := new(BigHeapQueue)
	// 记得冗余index=0
	p.data = make([]int, capacity+1)
	p.capacity = capacity
	p.length = 0
	return p
}

/**
上升节点
因为i下标节点所在的位置太小了，需要上升
*/
func (this *BigHeapQueue) rise(i int) {
	fI := this.father(i)
	if fI == 0 {
		// 说明已经在堆顶了
		return
	}
	val := this.data[i]
	father := this.data[fI]
	if father < val {
		this.swap(i, fI)
	} else {
		// 递归出口
		return
	}
	// 此时节点就在father所在的下标了
	this.rise(fI)
}

/**
下沉节点
因为i下标的节点所在的位置太大了，需要下沉
*/
func (this *BigHeapQueue) sink(i int) {
	// 既然要下沉，那就找最大
	lI := this.leftSon(i)
	rI := this.rightSon(i)

	if lI > this.length || rI > this.length {
		// 已经是叶子节点了
		return
	}
	left := this.data[lI]
	right := this.data[rI]

	var max int
	var maxI int
	if left >= right {
		max = left
		maxI = lI
	} else {
		max = right
		maxI = rI
	}
	// 得和最大的子节点进行比较，如果小于它，就和它swap
	val := this.data[i]
	if val < max {
		this.swap(i, maxI)
	} else {
		// 出口
		return
	}
	this.sink(maxI)
}

func (this *BigHeapQueue) put(val int) {
	this.BinaryHeap.put(val, this)
}

func (this *BigHeapQueue) del(i int) {
	this.BinaryHeap.del(i, this)
}
