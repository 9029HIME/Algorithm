package main

import "fmt"

func main() {
	queens := solveNQueens(8)
	fmt.Println(len(queens))
}

func solveNQueens(n int) [][]string {
	defaultString := ""
	chessboard := make([][]string, 0, n)
	for i := 0; i < n; i++ {
		row := make([]string, n)
		chessboard = append(chessboard, row)
		defaultString = defaultString + "."
	}
	path := make([]string, 0, n)
	result := make([][]string, 0, n)
	backTracking(n, chessboard, path, &result, defaultString)
	return result
}

func backTracking(n int, chessboard [][]string, path []string, result *[][]string, defaultString string) {
	length := len(path)
	if length == n {
		// 已经有解了
		solution := make([]string, 0, n)
		for _, v := range path {
			solution = append(solution, v)
		}
		*result = append(*result, solution)
		return
	}
	for i := 0; i < n; i++ {
		// length代表当前准备放棋子的行
		if !pass(n, chessboard, length, i) {
			// 这个位置不能放棋子
			continue
		}
		// 走到这里，说明当前i这个位置是可以放棋子的
		thisLineValue := replaceStrByIndex(defaultString, i, 'Q')
		path = append(path, thisLineValue)
		chessboard[length][i] = "Q"
		backTracking(n, chessboard, path, result, defaultString)
		// 走到这里，说明该回退上一行继续遍历了，执行回溯策略
		// 将path回退为上一级
		path = path[0:length]
		// 撤回棋盘的Q
		chessboard[length][i] = ""
	}
}

// 根据行和列，判断是否能在棋盘上放棋子
func pass(n int, chessboard [][]string, row int, column int) bool {
	// 检查下方有没有Q
	for i := row; i >= 0; i-- {
		if chessboard[i][column] == "Q" {
			return false
		}
	}

	// 检查左下方有没有Q，这里需要使用对角线遍历
	i := row
	j := column
	for i-1 >= 0 && j-1 >= 0 {
		if chessboard[i-1][j-1] == "Q" {
			return false
		}
		i--
		j--
	}

	// 检查右下方有没有Q，这里需要使用对角线遍历
	i = row
	j = column
	for i-1 >= 0 && j+1 <= n-1 {
		if chessboard[i-1][j+1] == "Q" {
			return false
		}
		i--
		j++
	}

	return true
}

func replaceStrByIndex(str string, index int, newValue rune) string {
	runes := []rune(str)
	runes[index] = newValue
	newString := string(runes)
	return newString
}
