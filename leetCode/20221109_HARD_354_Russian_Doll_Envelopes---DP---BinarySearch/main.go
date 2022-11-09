package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	i := maxEnvelopes(envelopes)
	fmt.Println(i)
}

func maxEnvelopes(envelopes [][]int) int {
	var result int
	length := len(envelopes)
	if length == 0 {
		return -1
	}
	// 宽度递增，同宽度下高度递减排序
	weightHeightSort(envelopes)

	dp := make([]int, length)
	// 通过LIS获取高度的结果集
	for i, h := range envelopes {
		if i == 0 {
			// dp出口
			dp[i] = 1
		} else {
			preMaxLength := 0
			// 截止到h[i]
			for j := 0; j < i; j++ {
				// 还得＜h[i]
				if envelopes[j][1] < h[1] {
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

func weightHeightSort(envelopes [][]int) {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[j][1] < envelopes[i][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
}

func Max(a int, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}
