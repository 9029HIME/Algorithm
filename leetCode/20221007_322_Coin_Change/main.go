package main

import (
	"fmt"
)

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。

示例1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0
*/

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	rememberSet := make(map[int]int)
	return dp(amount, coins, rememberSet)
}

func dp(balance int, coins []int, rememberSet map[int]int) int {
	if balance == 0 {
		return 0
	}

	if v, ok := rememberSet[balance]; ok {
		return v
	}
	result := 0
	for _, coin := range coins {
		balance2 := balance - coin
		i := 0
		if balance2 == 0 {
			i = 1
			rememberSet[balance] = 1
		} else if balance2 < 0 {
			i = -1
			rememberSet[balance] = -1
		} else {
			j := dp(balance2, coins, rememberSet)
			if j != -1 {
				i = j + 1
			} else {
				i = j
			}
		}
		result = min(i, result)
	}
	rememberSet[balance] = result
	return result
}

func min(tmp int, result int) int {
	if result <= 0 {
		return tmp
	}

	if tmp <= 0 {
		return result
	}

	if tmp <= result {
		return tmp
	} else {
		return result
	}
}
