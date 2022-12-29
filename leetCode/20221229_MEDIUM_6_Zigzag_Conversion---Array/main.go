package main

import "fmt"

func main() {
	fmt.Println(convert("AB", 1))
}

func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	if numRows < 2 {
		return s
	}
	row := 0
	step := -1
	for _, c := range s {
		rows[row] = rows[row] + string(c)
		if row == 0 || row == numRows-1 {
			step = -step
		}
		row = row + step
	}
	result := ""
	for _, s := range rows {
		result = result + s
	}
	return result
}
