package main

// 最高有几阶
var top = 50

// 能选择的步数
var steps []int = []int{1, 2}

// 动态规划的最优解结果
var stair2Result map[int]int = make(map[int]int, top)

func main() {
	print(climb(1))
}

/**
now int: 当前阶数
return: 当前阶数到终点的走法
*/
func climb(now int) int {
	if result, exist := stair2Result[now]; exist {
		return result
	}

	if now == top {
		// 继续走下去还是返回0，不如直接返回
		return 0
	}
	// 当阶选择数
	choice := 0
	// 下一阶走法数
	nextChoice := 0
	for _, step := range steps {
		end := now + step
		// 出口，当选择＞终点时，不会递归下去
		if end <= top {
			choice++
			// 下一阶的选择数
			nextChoice = nextChoice + climb(end)
		}
	}
	stair2Result[now] = choice + nextChoice
	return choice + nextChoice
}
