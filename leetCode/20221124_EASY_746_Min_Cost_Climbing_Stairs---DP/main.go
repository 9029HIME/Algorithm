package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

// 能选择的步数
var steps []int = []int{1, 2}

// 可以选择的起点
var origins []int = []int{0, 1}

func minCostClimbingStairs(cost []int) int {
	dp := make(map[int]int)
	minResult := math.MaxInt64

	for _, origin := range origins {
		result := stateTransition(cost, origin, dp)
		if result <= minResult {
			minResult = result
		}
	}
	return minResult
}

// 状态转移方程
func stateTransition(costs []int, i int, dp map[int]int) int {
	// dp出口，已经在最高阶了
	if i >= len(costs) {
		return 0
	}

	result := 0
	thisLevelValue := costs[i]
	minNextLevelValue := math.MaxInt64
	// 从某一阶开始，可以选择跨度
	for _, step := range steps {
		nextLevel := i + step
		nextLevelValue := 0

		// 记忆化剪枝
		v, ok := dp[nextLevel]
		if ok {
			nextLevelValue = v
		} else {
			nextLevelValue = stateTransition(costs, nextLevel, dp)
		}

		// 取最小值
		if nextLevelValue < minNextLevelValue {
			minNextLevelValue = nextLevelValue
		}
	}
	result = thisLevelValue + minNextLevelValue
	dp[i] = result
	return result
}
