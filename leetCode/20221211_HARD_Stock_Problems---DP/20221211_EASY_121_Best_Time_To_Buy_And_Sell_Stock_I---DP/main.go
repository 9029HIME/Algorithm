package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	maxDay := len(prices) - 1
	k := 1
	dp := make([][][]int, 0, maxDay+1)
	for i := 0; i <= maxDay; i++ {
		// k的可能性是[0,k]
		kSlice := make([][]int, 0, k+1)
		for j := 0; j < k+1; j++ {
			// 只有持有 和 不持有两种状态
			statusSlice := make([]int, 0, 2)
			kSlice = append(kSlice, statusSlice)
		}
		dp = append(dp, kSlice)

	}

	// 在最后一天不持有的最大收益
	notHold := stateTransition(prices, dp, maxDay, k, 0)

	return notHold
}

func stateTransition(prices []int, dp [][][]int, maxDay int, k int, status int) int {
	for j := 0; j <= maxDay; j++ {
		// 从第1天开始，算出最大收益
		if j == 0 {
		}
	}
	return dp[maxDay][k][0]
}
