package main

import "math"

// 能选择的步数
var steps []int = []int{1, 2}

// 从index阶要上阶的花费，length代表总阶数
var cost []int = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
var top int = len(cost)

// 可以选择的起点（这是题目的附加条件，可以选择下标=0或下标=1的台阶开始爬楼梯
var origins []int = []int{0, 1}

// 动态规划的最优解结果
var stair2Result map[int]int = make(map[int]int, top)

func main() {
	print(startClimbing())
}

/**
对递归方法的再一层封装，主要是完成可选起点的功能
*/
func startClimbing() int {
	result := math.MaxInt64
	for origin, _ := range origins {
		temp := climb(origin)
		if result == math.MaxInt64 {
			result = temp
		} else {
			result = Min(result, temp)
		}
	}
	return result
}

/**
now int: 当前阶数
return: 当前阶数到终点的最小花费
*/
func climb(now int) int {
	if now == top {
		// 继续走下去还是返回0，不如直接返回
		return 0
	}
	if dpResult, ok := stair2Result[now]; ok {
		return dpResult
	}
	thisStairCost := cost[now]
	minNextStepCost := math.MaxInt64
	nextStepCost := 0

	for _, step := range steps {
		end := now + step
		// 出口，当选择＞终点时，不会递归下去
		if end <= top {
			// 当前选择下一阶的最低花费
			nextStepCost = climb(end)
			if minNextStepCost == math.MaxInt64 {
				minNextStepCost = nextStepCost
			} else {
				minNextStepCost = Min(minNextStepCost, nextStepCost)
			}
		}
	}
	normalResult := thisStairCost + minNextStepCost
	stair2Result[now] = normalResult
	return normalResult
}

func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
