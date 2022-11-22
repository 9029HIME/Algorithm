package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2}, 3))
}

func combinationSum4(nums []int, target int) int {
	// dp[i] = 在nums里dp[i]个解集，解集里面每个数的和都=i
	dp := make(map[int]int, target+1)
	return stateTransition(nums, target, dp)
}

// 状态转移方程
func stateTransition(nums []int, i int, dp map[int]int) int {
	// dp出口： dp[0] = 1
	if i == 0 {
		return 1
	}

	result := 0

	for _, v := range nums {
		// 选择带来的影响：算出两者的差
		previous := i - v
		if previous < 0 {
			continue
		}
		// 记忆化剪枝
		previousValue, ok := dp[previous]
		if !ok {
			previousValue = stateTransition(nums, previous, dp)
		}
		// dp[i] = SUM(dp[i的previous])
		result = result + previousValue
	}
	dp[i] = result
	return result
}
