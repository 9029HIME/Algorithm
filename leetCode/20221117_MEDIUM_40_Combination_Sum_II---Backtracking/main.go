package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	results := combinationSum2(candidates, target)
	fmt.Println(results)
}

func combinationSum2(candidates []int, target int) [][]int {
	length := len(candidates)
	results := make([][]int, 0, length)
	path := make([]int, 0, length)
	// 本题核心1：进行排序，保证相同的数字紧挨在一起。
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	fmt.Println(candidates)
	backTracking(candidates, target, 0, path, 0, &results)
	return results
}

func backTracking(candidates []int, target int, existSum int, path []int, start int, results *[][]int) {
	// 回溯出口
	length := len(path)
	if existSum == target {
		result := make([]int, 0, length)
		for _, v := range path {
			result = append(result, v)
		}
		*results = append(*results, result)
		return
	}

	for i := start; i < len(candidates); i++ {
		v := candidates[i]
		// 本题核心3：如果本次数字和上一次数字相等，剪枝。
		if existSum+v > target || (i > 0 && candidates[i] == candidates[i-1]) {
			continue
		}
		path = append(path, v)
		existSum = existSum + v
		// 本题核心2：以下一个数字为起点，进行遍历
		backTracking(candidates, target, existSum, path, i+1, results)
		// 回溯策略
		path = path[0:length]
		existSum = existSum - v
	}
}
