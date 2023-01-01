package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt64
	sort.Ints(nums)
	tail := len(nums) - 1

	for i, num := range nums {
		j := i + 1
		k := tail
		for j < k {
			sum := num + nums[j] + nums[k]
			abs := Abs(sum, target)
			if abs == 0 {
				return sum
			}
			if abs < Abs(result, target) {
				// 相差更小，替换result
				result = sum
			}
			// 根据sum和target的关系，挪动指针
			if sum > target {
				// 偏大了，减少k
				k--
			} else {
				// 偏小了，增加j
				j++
			}
		}
	}
	return result
}

func Abs(a int, b int) int {
	c := a - b
	if c < 0 {
		c = -c
	}
	return c
}
