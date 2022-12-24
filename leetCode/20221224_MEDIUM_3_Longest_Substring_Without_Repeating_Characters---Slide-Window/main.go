package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {
	return Slide(s)
}

func Slide(s string) int {
	result := 0
	length := len(s)
	if length == 0 {
		return 0
	}
	// k:v = char : frequency
	window := make(map[string]int)
	left := 0
	right := 0
	for left < length && right < length {
		// 获取right资源
		rightC := string(s[right])

		right++

		// 更新窗口
		window[rightC] = window[rightC] + 1

		// 是否需要移动left
		for LeftNeedShrink(rightC, window) {
			// 获取left资源
			leftC := string(s[left])

			left++

			// 更新窗口
			window[leftC] = window[leftC] - 1
		}
		tmpResult := right - left
		if result < tmpResult {
			result = tmpResult
		}
	}

	return result
}

func LeftNeedShrink(rightC string, window map[string]int) bool {
	// 右指针移动后，发现了重复字符
	return window[rightC] > 1
}
