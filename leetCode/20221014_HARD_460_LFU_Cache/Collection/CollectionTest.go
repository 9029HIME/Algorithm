package main

import (
	"fmt"
	"strconv"
)

func main() {
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

}

type Data struct {
	payload int
}

func (this *Data) hashCode() int {
	return this.payload
}

func (this *Data) String() string {
	return strconv.Itoa(this.payload)
}
