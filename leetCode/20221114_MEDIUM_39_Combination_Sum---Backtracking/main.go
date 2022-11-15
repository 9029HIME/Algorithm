package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	results := combinationSum(candidates, target)
	fmt.Println(results)
}

func combinationSum(candidates []int, target int) [][]int {
	length := len(candidates)
	results := make([][]int, 0, length)
	path := make([]int, 0, length)
	backTracking(candidates, target, 0, path, 0, &results)
	return results
}

func backTracking(candidates []int, target int, existSum int, path []int, start int, results *[][]int) {
	// 先定义回溯出口
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
		if existSum+v > target {
			continue
		}
		path = append(path, v)
		existSum = existSum + v
		backTracking(candidates, target, existSum, path, i, results)
		// 回溯策略
		path = path[0:length]
		existSum = existSum - v
	}
}
