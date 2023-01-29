package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	if n > 2 {
		result := &strings.Builder{}
		previous := countAndSay(n - 1)
		// n个相同字符为一组group，分割previous
		element := ' '
		group := &strings.Builder{}
		for _, c := range previous {
			if element != ' ' && c != element {
				// 计算group结果，添加到result
				result.WriteString(strconv.Itoa(group.Len()))
				result.WriteRune(element)
				// 重置group
				group.Reset()
				element = c
				group.WriteRune(c)
			} else {
				element = c
				group.WriteRune(c)
			}
		}
		// 循环结束后，会有一组group未添加到result，这里补充上
		result.WriteString(strconv.Itoa(group.Len()))
		result.WriteRune(element)
		return result.String()
	}
	return "error"
}
