package _0221020_Null_Null_Binary_Heap_Prority_Queue___Data_Struct_Design

import "fmt"

/**
二叉堆的基本属性
*/
type BinaryHeap struct {
	data     []int
	capacity int
	length   int
}

/**
上升、下沉的具体实现，以接口的形式
TODO Golang好像不支持
1. 父类A方法调用父类B方法
2. 子类重写父类B方法
3. 子类调用父类A方法，实际调用子类B方法
这样的特性，所以采用接口的方式实现
*/
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

/**
打印数组
*/
func (this *BinaryHeap) introduce() {
	result := ""
	for i, datum := range this.data {
		if i == 0 {
			continue
		}
		val := fmt.Sprintf("%v->", datum)
		result = result + val
	}
	result = result[0 : len(result)-2]
	fmt.Println(result)
}

/**
我这里就不做动态扩容了，超出容量直接panic
*/
func (this *BinaryHeap) put(val int, rs RiseSink) {
	if this.length == this.capacity {
		panic("优先级队列已满")
	} else {
		/*
			如何算出当前队列的尾巴节点？
			如果我的length = 0，代表我这个插入的数据是第一个数据，也就是堆顶，插入index=1
			如果我的length = 1，说明我插入的这个数据是堆顶的左孩子节点，插入index = 2
		*/
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
		panic("无法删除没有叶子节点的数据")
	}

	// 最后一个叶子节点下标 是 this.length
	to := this.length
	this.swap(i, to)

	// 此时待删除节点 就在 this.length所在的下标了
	this.data[to] = 0
	this.length--

	// 此时i节点 是 rise上来节点，得判断是否要sink
	rs.sink(i)
}
