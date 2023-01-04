package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0, 2)
	length := len(nums)
	tail := length - 1
	for i, num := range nums {
		// 避免i的重复值
		if i > 0 && nums[i-1] == num {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			// 避免j的重复值，为什么要j > i + 1？有可能[2,2,2,2,2]，此时i和j都等于2，如果没有j > i + 1就直接跳过了
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := tail
			for k < l {
				sum := num + nums[j] + nums[k] + nums[l]
				if sum > target {
					// l偏大，缩小l
					l--
				} else if sum < target {
					// k变小，加大k
					k++
				} else {
					// 有解
					result := []int{num, nums[j], nums[k], nums[l]}
					results = append(results, result)
					// 挪动k和l，注意去重
					for k < l && nums[k+1] == nums[k] {
						k++
					}
					for k < l && nums[l-1] == nums[l] {
						l--
					}
					// 走到这一步，下一个k和下一个l必定不重复，放心挪动
					k++
					l--
				}
			}
		}
	}
	return results
}
