package main

// TODO 我只算出了最多赚多少钱，但还没算Job的搭配

// 从早到晚排序的job
var sortedJob []int = []int{3, 1, 2, 5, 4, 6, 7, 8}

// 任务最大数量
var length int = 8
var job2Index map[int]int = map[int]int{
	3: 0,
	1: 1,
	2: 2,
	5: 3,
	4: 4,
	6: 5,
	7: 6,
	8: 7,
}
// 组合搭配
var choice string = ""
// 最优解集合
var job2Result map[int]int = make(map[int]int)
func main() {
	print(value(sortedJob[0]))
}

// 其实应该用更好的方式（根据job的时间占用来排除），不过为了方便，这里直接用代码写死
func should(jobId int) int {
	nextId := -1
	switch jobId {
	case 3:
		nextId = 7
	case 1:
		nextId = 4
	case 2:
		nextId = 6
	case 5:
		nextId = 8
	case 4:
		nextId = 8
	case 6:
		nextId = -1
	case 7:
		nextId = -1
	case 8:
		nextId = -1
	}
	return nextId
}

// 一样，为了方便直接代码写死，并且金额为整型
func amount(jobId int) int {
	reward := 0
	switch jobId {
	case 3:
		reward = 8
	case 1:
		reward = 5
	case 2:
		reward = 1
	case 5:
		reward = 6
	case 4:
		reward = 4
	case 6:
		reward = 3
	case 7:
		reward = 2
	case 8:
		reward = 4
	}
	return reward
}

func next(jobId int) int {
	nextId := -1
	if index, ok := job2Index[jobId]; ok && (index+1) < 8 {
		nextId = sortedJob[index+1]
	}
	return nextId
}

func value(jobId int) int {
	// 出口
	if jobId == -1 {
		return 0
	}
	// 动态规划关键点1，是否已有最优解？
	dpResult,exist:= job2Result[jobId]
	if exist{
		return dpResult
	}
	// 派生子问题
	result := Max(amount(jobId)+value(should(jobId)), value(next(jobId)))
	// 动态规划关键点2，存入最优解
	job2Result[jobId] = result
	return result
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
