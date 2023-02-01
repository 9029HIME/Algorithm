package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{1, 1}))
}

func firstMissingPositive(nums []int) int {
	length := len(nums)
	// 将负数统一设置为length + 1
	for i, v := range nums {
		if v <= 0 {
			nums[i] = length + 1
		}
	}
	// 取[1,length]的绝对值n，将n-1作为下标，把nums[n-1]统一设为负数
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if v <= length {
			if nums[v-1] > 0 {
				nums[v-1] = -nums[v-1]
			}
		}
	}
	// 将第一个正整数的下标+1作为结果
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}

	return length + 1
}
