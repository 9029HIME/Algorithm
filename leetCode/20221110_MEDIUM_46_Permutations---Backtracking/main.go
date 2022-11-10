package main

import "fmt"

func main() {
	nums := []int{0, 1, 2}
	result := permute(nums)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	length := len(nums)
	used := make([]bool, length)
	// 记得要指定切片的长度、再指定容量,否则会初始化[0,0,0]，而不是[]，这里踩了一个基础坑
	path := make([]int, 0, length)
	result := make([][]int, 0, length)
	backtracking(nums, used, path, &result)
	return result
}

func backtracking(nums []int, used []bool, path []int, result *[][]int) {
	lenght := len(path)
	if lenght == len(nums) {
		// 已经到达叶子节点了，代表已存在一个结果
		leaf := make([]int, 0, lenght)
		for _, v := range path {
			leaf = append(leaf, v)
		}
		*result = append(*result, leaf)
		return
	}
	for i, v := range nums {
		if used[i] == true {
			// 这个数字在选择路径上已经使用过了
			continue
		}
		path = append(path, v)
		used[i] = true
		backtracking(nums, used, path, result)
		used[i] = false
		path = path[0:lenght]
	}
}
