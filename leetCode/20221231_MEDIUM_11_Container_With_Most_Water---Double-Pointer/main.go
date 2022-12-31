package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	result := 0
	for i < j {
		// 最小高度
		min := Min(height[i], height[j])
		// 距离
		distance := j - i
		// 面积
		area := min * distance
		result = Max(result, area)
		// 挪动最小的线
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
