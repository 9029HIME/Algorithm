package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2}))
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
			statusSlice := make([]int, 2)
			kSlice = append(kSlice, statusSlice)
		}
		dp = append(dp, kSlice)

	}

	// 在最后一天不持有的最大收益
	notHold := stateTransition(prices, dp, maxDay, k, 0)

	return notHold
}

func stateTransition(prices []int, dp [][][]int, maxDay int, k int, status int) int {
	// 初始化k的出口
	for i := 0; i < len(prices); i++ {
		dp[i][0][1] = 0
		dp[i][0][0] = 0
	}

	// 从第1天开始，算出最大收益
	for i := 0; i <= maxDay; i++ {
		for j := k; j >= 0; j-- {
			// 初始化i的出口
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			// 今天k情况下，没持有的最大收益 = Max(昨天没持有，今天继续每持有 , 昨天持有，今天卖出)
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 今天k的情况下，持有的最大收益 = Max(昨天k-1没持有，今天买入 , 昨天持有，今天继续持有)
			dp[i][k][1] = max(dp[i-1][k-1][0]-prices[i], dp[i-1][k][1])
		}
	}
	return dp[maxDay][k][0]
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
