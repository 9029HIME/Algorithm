package main

import (
	"fmt"
	"strconv"
)

type LFU struct {
	kv       map[int]int
	kf       map[int]int
	fvs      map[int]*LinkedHashSet
	capacity int
	minFeq   int
}

/**
构造方法
*/
func InitLFU(capacity int) *LFU {
	return &LFU{
		kv:       make(map[int]int),
		kf:       make(map[int]int),
		fvs:      make(map[int]*LinkedHashSet),
		capacity: capacity,
		minFeq:   0,
	}
}

/**
LFU-对外暴露-存放数据的方法
*/
func (this *LFU) put(key HashEntity, value int) {
	hashCode := key.hashCode()
	// 先看kv里是否存在？
	if _, ok := this.kv[hashCode]; ok {
		this.kv[hashCode] = value
		// 提鲜操作
		this.refresh(key)
		return
	}

	if len(this.kv) == this.capacity {
		this.purge()
	}

	this.kv[hashCode] = value
	this.kf[hashCode] = 1
	// 新加入的元素肯定是1
	this.minFeq = 1
	newSet, ok := this.fvs[1]
	if !ok {
		newSet = InitLinkedHashSet(8)
		this.fvs[1] = newSet
	}
	newSet.put(key)
}

/**
LFU-对外暴露-获取数据的方法
*/
func (this *LFU) get(key HashEntity) (int, bool) {
	hashCode := key.hashCode()
	value, ok := this.kv[hashCode]
	if !ok {
		return 0, false
	}
	this.refresh(key)
	return value, true
}

/**
LFU-自身使用-提鲜数据的方法
*/
func (this *LFU) refresh(key HashEntity) int {
	hashCode := key.hashCode()
	// 既然已经存在于kv，那么一定存在于kf和fvs
	feq := this.kf[hashCode]
	set := this.fvs[feq]
	set.remove(key)

	newFeq := feq + 1
	newSet, ok := this.fvs[newFeq]
	if !ok {
		newSet = InitLinkedHashSet(8)
		this.fvs[newFeq] = newSet
	}
	newSet.put(key)
	this.kf[hashCode] = newFeq

	// 记得检查旧set是否为空，为空则清除，否则会内存泄漏
	if set.isEmpty() {
		delete(this.fvs, feq)
	}

	/*
		判断最小频率
		既然能走到提鲜这一步，说明minFeq已经有确切的值了（不是0）
		1. 如果旧feq == minFeq，旧set不存在了，那么就是新set的freq
		2. 如果旧feq == minFeq，旧set还在，那么minFeq不变
		3. 如果旧feq > minFeq，旧set不存在了，minFeq不变
		4. 如果旧feq > minFeq，旧set还在，minFeq不变
	*/
	if feq == this.minFeq {
		if set.isEmpty() {
			this.minFeq = newFeq
		}
	}

	return newFeq
}

/**
LFU-自身使用-清除频率最小、最旧数据的方法
*/
func (this *LFU) purge() {
	minFeq := this.minFeq
	set := this.fvs[minFeq]
	pendingDel := set.head.next
	data := pendingDel.data
	hashCode := data.hashCode()

	delete(this.kv, hashCode)
	delete(this.kf, hashCode)
	set.remove(data)

	if set.isEmpty() {
		delete(this.fvs, minFeq)
	}

	/*
		假如 fvs 是 1-[x] 3-[y,z]的结构
		此时清掉了x在kv、kf的存在，也清除了1在fvs的存在（包括set），那么接下来的minFeq就应该等于3
		那么问题来了，怎么从1找到3呢？
		能找其实有一种方案： 从1开始循环递增，直到在fvs中到有效的set

		其实我陷入了一个误区，回头看一下什么时候会用到minFeq？新增数据（直接设为1)、提鲜数据（视情况+1或者不变）、清除数据之前（拿到正确的minFeq）
		也就是说，在清除节点后，我是不需要修改minFeq的值的，保持不变即可

		可能我之后看这篇笔记会不太理解，为什么purge不需要修改minFeq？
		以上面的1-[x] 3-[y,z]举例，purge之后就是 3-[y,z]、minFeq=1。我提鲜y、z节点后完全不用管minFeq，因为毕竟3本身不是minFeq，而且提鲜不会导致数据容量发生变化。
			当我插入新节点后，我会给minFeq直接赋值为1，容量量满后，准备再次purge，才要拿到正确的minFeq。
			但容量满是需要经历 purge → 插入 → 提鲜（不一定有） → 容量满，这几个步骤，在这中间就已经会将minFeq重置为1了，所以没必要在purge阶段修改minFeq。

		简单来说就是：既然你已经purge了，那么下次purge肯定发生了put或者refresh，minFeq就会在这两个阶段动态维护好了，purge你就别操这个心了。
	*/
}

/**
LFU-测试用例方法-打印Kv的值
*/
func (this *LFU) introduceKv() {
	kv := this.kv

	kvResult := "kv: "
	for k, v := range kv {
		kvValue := fmt.Sprintf("[%s-%v]---", strconv.Itoa(k), strconv.Itoa(v))
		kvResult = kvResult + kvValue
	}
	fmt.Println(kvResult)
}

/**
LFU-测试用例方法-打印Kf的值
*/
func (this *LFU) introduceKf() {
	kf := this.kf

	kvResult := "kf: "
	for k, v := range kf {
		kvValue := fmt.Sprintf("[%s-%v]---", strconv.Itoa(k), strconv.Itoa(v))
		kvResult = kvResult + kvValue
	}
	fmt.Println(kvResult)
}

/**
LFU-测试用例方法-打印Fvs的值
*/
func (this *LFU) introduceFvs() {
	fvs := this.fvs

	fvsResult := "fvs: "
	for k, v := range fvs {
		node := v.head
		setValue := ""
		for node.next != nil && node.next != v.tail {
			node = node.next
			msg := fmt.Sprintf("<%s>→", node.data)
			setValue = setValue + msg
		}
		kvValue := fmt.Sprintf("[%s-%s]---", strconv.Itoa(k), setValue)
		fvsResult = fvsResult + kvValue
	}
	fmt.Println(fvsResult)
}

/**
LFU-测试用例方法-打印LFU所有值
*/
func (this *LFU) introduce() {
	this.introduceKv()
	this.introduceKf()
	this.introduceFvs()
	fmt.Println("minFeq：", this.minFeq)
	fmt.Println()
}
