package main

import "fmt"

func main() {
	str := "ABCACBG"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	result := ""
	for i, _ := range s {
		p1 := getPalindromic(s, i, i)
		p2 := getPalindromic(s, i, i+1)
		if len(p1) > len(result) {
			result = p1
		}
		if len(p2) > len(result) {
			result = p2
		}
	}
	return result
}

/**
根据left和right定位中间字符，返回这个中间字符的最大回文
*/
func getPalindromic(str string, left int, right int) string {
	length := len(str)

	for left >= 0 && right <= length-1 && str[left] == str[right] {
		left--
		right++
	}

	// 能走到这一步，说明left和right的值 或者其代表的字符 已经不符合要求了

	// 在golang里，也是包左不包右的
	return str[left+1 : right]
}
