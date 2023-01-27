package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 7, 10}, 4))
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	// 边界处理
	if target < nums[left] {
		return left
	}
	if target > nums[right] {
		return length
	}

	for left <= right {
		mid := (right + left) / 2
		v := nums[mid]
		if target == v {
			return mid
		} else if target < v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	v := nums[left]
	if v < target {
		return left + 1
	} else {
		return left
	}
}
