package main

import "fmt"

var candidates = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	k := 3
	n := 7
	results := combinationSum3(k, n)
	fmt.Println(results)
}

func combinationSum3(k int, n int) [][]int {
	length := len(candidates)
	results := make([][]int, 0, length)
	path := make([]int, 0, length)
	backTracking(n, 0, path, 0, k, &results)
	return results
}

func backTracking(target int, existSum int, path []int, start int, k int, results *[][]int) {
	// 先定义回溯出口
	length := len(path)
	// 长度到达k，再判断是否要添加结果
	if length == k {
		if existSum == target {
			result := make([]int, 0, length)
			for _, v := range path {
				result = append(result, v)
			}
			*results = append(*results, result)
		}
		return
	}

	for i := start; i < len(candidates); i++ {

		v := candidates[i]
		// 手动剪枝
		if existSum+v > target {
			continue
		}
		path = append(path, v)
		existSum = existSum + v
		backTracking(target, existSum, path, i+1, k, results)
		// 回溯策略
		path = path[0:length]
		existSum = existSum - v
	}
}
