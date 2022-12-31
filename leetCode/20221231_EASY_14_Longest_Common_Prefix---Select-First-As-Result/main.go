package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	// 将第一个字符串作为最长公共前缀
	result := strs[0]
	// 从第二个字符串开始，比较最长公共前缀
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		// 临界出口：最小的长度
		min := Min(len(result), len(str))
		// 公共前缀长度
		count := 0
		for j := 0; j < min; j++ {
			s1 := result[j]
			s2 := str[j]
			if s1 == s2 {
				count++
			} else {
				break
			}
		}
		// 没有公共前缀，直接返回
		if count == 0 {
			return ""
		}
		// 更新结果
		result = result[:count]
	}
	return result
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
