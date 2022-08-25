package main

/*
1.题目
寻找数组的中心下标

2.描述
给你一个整数数组 nums ，请计算数组的 中心下标 。
数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。

3.示例
输入：nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

4.思路
这道题的核心数学思想是：假设中心下标是i，那么 [(i前面的值的和) * 2 ] + i代表的值 = 数组值总和

5.结果：
解法的时间复杂度：O(2n) = O(n)

*/
func main() {
	var nums []int = []int{-1, -1, -1, -1, -1, 0}
	print(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	result := -1
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for k, v := range nums {
		a := (leftSum * 2) + v
		if a == sum {
			result = k
			break
		} else {
			leftSum += v
		}
	}
	return result
}
