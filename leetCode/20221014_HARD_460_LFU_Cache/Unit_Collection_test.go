package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Data struct {
	payload int
}

func (this *Data) hashCode() int {
	return this.payload
}

func (this *Data) String() string {
	return strconv.Itoa(this.payload)
}

/**
测试打印LinkedHashSet的map、链表数据
*/
func TestLinkedHashSet(t *testing.T) {
	set := InitLinkedHashSet(16)
	one := &Data{
		payload: 1,
	}
	two := &Data{
		payload: 2,
	}
	three := &Data{
		payload: 3,
	}

	// 新增后看看集合的数据
	set.put(one)
	set.put(two)
	set.put(three)
	set.introduce()
	fmt.Println()

	// 测试contains-api
	fmt.Println(set.contains(one))
	fmt.Println()

	// 测试删除2节点后，集合的数据
	set.remove(two)
	set.introduce()
	fmt.Println()

	// 测试删除3节点后，集合的数据
	set.remove(three)
	set.introduce()
	fmt.Println()

	// 测试删除1节点后，集合的数据
	set.remove(one)
	set.introduce()
	fmt.Println()

	// 测试假删除后，集合的数据
	set.remove(one)
	set.introduce()
	fmt.Println()

	// 测试isEmpty-api
	fmt.Println(set.isEmpty())
	fmt.Println()

	set.delNode(set.head)
}

/**
测试打印LFU的Kv数据
*/
func TestLFUKvData(t *testing.T) {
	lfu := InitLFU(3)
	one := &Data{
		payload: 1,
	}
	two := &Data{
		payload: 2,
	}
	three := &Data{
		payload: 3,
	}
	four := &Data{
		payload: 4,
	}

	lfu.introduceKv()
	lfu.put(one, 101)
	lfu.introduceKv()
	lfu.put(two, 102)
	lfu.introduceKv()
	lfu.put(three, 103)
	lfu.introduceKv()
	lfu.put(four, 104)
	lfu.introduceKv()
}

/**
测试打印LFU的Kv数据
*/
func TestLFUKfData(t *testing.T) {
	lfu := InitLFU(3)
	one := &Data{
		payload: 1,
	}
	two := &Data{
		payload: 2,
	}
	three := &Data{
		payload: 3,
	}
	four := &Data{
		payload: 4,
	}

	lfu.introduceKf()
	lfu.put(one, 101)
	lfu.introduceKf()
	lfu.put(two, 102)
	lfu.introduceKf()
	lfu.put(three, 103)
	lfu.introduceKf()
	lfu.put(four, 104)
	lfu.introduceKf()
}

/**
测试打印LFU的Fvs数据
*/
func TestLFUFvsData(t *testing.T) {
	lfu := InitLFU(3)
	one := &Data{
		payload: 1,
	}
	two := &Data{
		payload: 2,
	}
	three := &Data{
		payload: 3,
	}
	four := &Data{
		payload: 4,
	}

	lfu.introduceFvs()
	lfu.put(one, 101)
	lfu.introduceFvs()
	lfu.put(two, 102)
	lfu.introduceFvs()
	lfu.put(three, 103)
	lfu.introduceFvs()
	lfu.put(four, 104)
	lfu.introduceFvs()
}

/**
测试打印LFU的所有数据
*/
func TestLFUIntroduce(t *testing.T) {
	lfu := InitLFU(3)
	one := &Data{
		payload: 1,
	}
	two := &Data{
		payload: 2,
	}
	three := &Data{
		payload: 3,
	}
	four := &Data{
		payload: 4,
	}
	var getData int

	fmt.Println("初始化情况下↓")
	lfu.introduce()

	fmt.Println("存放1-101数据后↓")
	lfu.put(one, 101)
	lfu.introduce()

	fmt.Println("存放2-102数据后↓")
	lfu.put(two, 102)
	lfu.introduce()

	getData, _ = lfu.get(one)
	fmt.Println(fmt.Printf("get一次1的数据：%v，LFU的结果是↓", getData))
	lfu.introduce()

	getData, _ = lfu.get(one)
	fmt.Println(fmt.Printf("再get一次1的数据：%v，LFU的结果是↓", getData))
	lfu.introduce()

	getData, _ = lfu.get(two)
	fmt.Println(fmt.Printf("get一次2的数据：%v，LFU的结果是↓", getData))
	lfu.introduce()

	fmt.Println("存放3-103数据后↓")
	lfu.put(three, 103)
	lfu.introduce()

	getData, _ = lfu.get(three)
	fmt.Println(fmt.Printf("get一次3的数据：%v，LFU的结果是↓", getData))
	lfu.introduce()

	fmt.Println("存放4-104数据后↓")
	lfu.put(four, 104)
	lfu.introduce()
}
