package main

import "fmt"

func main() {
	// 二分查找的单元测试
	fmt.Println(Dichotomy([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 6, 5))
	//  本题的测试用例
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1, 3, 5}, 2))
}

func search(nums []int, target int) int {
	return DoSearch(nums, target, 0, len(nums)-1)
}

/**
nums = 数组
target = 期望值
left = 二分的范围最小值
right = 二分的范围最大值
在这个函数中，会将nums以left和right为基础进行二分，其中一个区间必定有序，另外一个区间可能无序，也可能有序。
*/
func DoSearch(nums []int, target int, left int, right int) int {
	// 找中点的临界case（具体看思路）：
	if right == left+1 || left == right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		return -1
	}

	mid := (left + right) / 2
	lV := nums[left]
	mV := nums[mid]
	nextLeft := 0
	nextRight := 0
	if lV <= mV {
		// 在左区间有序
		nextRight = right
		right = mid
		nextLeft = mid
	} else {
		// 在右区间有序
		nextLeft = left
		left = mid
		nextRight = mid
	}
	dichotomy := Dichotomy(nums, left, right, target)
	if dichotomy != -1 {
		return dichotomy
	}

	// 有序区间没有，继续doSearch无序区间
	return DoSearch(nums, target, nextLeft, nextRight)

}

/**
nums = 数组
left = 二分的范围最小值
right = 二分的范围最大值
target = 二分查找的值
return 值的下标
*/
func Dichotomy(nums []int, left int, right int, target int) int {
	// 二分查找的临界case（具体看思路）：
	if right == left+1 || left == right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[mid] < target && target < nums[right] {
		return Dichotomy(nums, mid, right, target)
	} else if nums[left] < target && target < nums[mid] {
		return Dichotomy(nums, left, mid, target)
	} else {
		return -1
	}
}
