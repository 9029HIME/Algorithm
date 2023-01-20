package main

import "fmt"

func main() {
	fmt.Println(SearchFirst([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(SearchLast([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 1))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	return []int{SearchFirst(nums, target), SearchLast(nums, target)}
}

// 普通的二分代码
//func SearchFirst(nums []int, target int) int {
//	length := len(nums)
//	left := 0
//	right := length - 1
//	for left <= right {
//		mid := (left + right) / 2
//		v := nums[mid]
//		if v == target {
//			return mid
//		} else if v < target {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return -1
//}

func SearchFirst(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		v := nums[mid]

		if mid != 0 && nums[mid-1] == target {
			right = mid - 1
			continue
		}

		if v == target {
			return mid
		} else if v < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SearchLast(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		v := nums[mid]

		if mid != length-1 && nums[mid+1] == target {
			left = mid + 1
			continue
		}

		if v == target {
			return mid
		} else if v < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
