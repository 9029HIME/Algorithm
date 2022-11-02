package main

import "fmt"

func main() {
	indices := twoSum([]int{2, 7, 11, 15}, 9)
	for _, index := range indices {
		fmt.Println(index)
	}
}

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	if length < 2 {
		return nil
	}

	left := 0
	right := length - 1

	for left != right {
		vL := numbers[left]
		vR := numbers[right]
		tmp := vL + vR
		if tmp == target {
			// 注意题目要的是index + 1
			return []int{left + 1, right + 1}
		} else if tmp > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
