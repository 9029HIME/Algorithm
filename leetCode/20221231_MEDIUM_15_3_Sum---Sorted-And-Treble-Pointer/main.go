package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, 2)
	tail := len(nums) - 1
	for i, num := range nums {
		if num > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := tail
		for j < k {
			sum := num + nums[j] + nums[k]
			if sum > 0 {
				// 偏大了，让最大值减少一点
				k--
			} else if sum < 0 {
				// 偏小了，让最小值加大一点
				j++
			} else {
				// 能走到这，说明 sum == 0，添加结果
				r := []int{num, nums[j], nums[k]}
				result = append(result, r)
				// 找到解后，缩小j和k的范围，寻找同等i的下一个解
				for j < k && nums[j+1] == nums[j] {
					// 避免重复解，对重复j进行跳过
					j++
				}
				for j < k && nums[k-1] == nums[k] {
					// 避免重复解，对重复k进行跳过
					k--
				}
				// 能走到这，说明下一个j和k都不是重复的，可以自增
				j++
				k--
			}
		}
	}
	return result
}
