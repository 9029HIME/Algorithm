package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{0}))
}

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	var result int
	dp := make([]int, length)

	for i, v := range nums {
		if i == 0 {
			// dp出口
			dp[i] = 1
		} else {
			preMaxLength := 0
			// 截止到num[i]
			for j := 0; j < i; j++ {
				// 还得＜num[i]
				if nums[j] < v {
					// 的最大MaxLength
					preMaxLength = Max(dp[j], preMaxLength)
				}
			}
			if preMaxLength == 0 {
				// 前面根本没有比自己小的
				dp[i] = 1
			} else {
				dp[i] = preMaxLength + 1
			}
		}
		// 取MaxLength的最大值作为返回结果
		result = Max(result, dp[i])
	}
	return result
}

func Max(a int, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
