package main

import "fmt"

func main() {
	fmt.Println(strStr("buasadsct", "sad"))
}

func strStr(haystack string, needle string) int {
	return Slide(haystack, needle)
}

func Slide(haystack string, needle string) int {
	left := 0
	right := 0
	length := len(haystack)
	needleLen := len(needle)
	window := ""
	for left < length && right < length {
		// right资源
		rightC := string(haystack[right])
		// 增加right
		right++
		// 更新窗口
		window = window + rightC

		for LeftNeedShrink(left, right, needleLen) {
			// 比较window 和 needle
			if window == needle {
				return left
			}
			// 不需要获取left资源，直接++
			left++
			// 更新窗口
			window = window[1:]
		}

	}
	return -1
}

func LeftNeedShrink(left int, right int, needleLen int) bool {
	return right-left+1 > needleLen
}
