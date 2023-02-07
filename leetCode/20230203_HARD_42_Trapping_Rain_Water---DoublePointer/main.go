package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])
		if leftMax < rightMax {
			// 对于left来说，leftLeftMax 必定 ＜ leftRightMax
			result = result + (leftMax - height[left])
			left++
		} else {
			//对于right来说,rightRightMax 必定 ≤ rightLeftMax
			result = result + (rightMax - height[right])
			right--
		}
	}
	return result
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
