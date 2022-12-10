package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("aa", "aa"))
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
}

func minWindow(s string, t string) string {
	window := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, v := range t {
		existV, ok := tMap[v]
		if ok {
			tMap[v] = existV + 1
		} else {
			tMap[v] = 1
		}
		window[v] = 0
	}

	result := slide(window, tMap, s)
	return result
}

func slide(window map[rune]int, tMap map[rune]int, s string) string {
	result := ""
	length := len(s)
	left := 0
	right := 0
	valid := 0
	for right < length && left < length {
		// 获取right资源
		rC := rune(s[right])

		right++

		// 更新窗口
		if needCount, ok := tMap[rC]; ok {
			// 如果是t包含的字符的话
			window[rC] = window[rC] + 1
			// 如果这个字符已经达标了
			if window[rC] == needCount {
				valid++
			}
		}

		for leftNeedShrink(valid, len(tMap)) {
			// 走到这，意味着left和right已经找到了覆盖子串（不一定最小），所以先和result比较。
			s2 := s[left:right]
			if result == "" {
				result = s2
			} else if len(result) >= len(s2) {
				result = s2
			}

			// 获取left资源
			lC := rune(s[left])

			left++

			// 更新窗口
			if existCount, ok := window[lC]; ok && existCount > 0 {
				existCount = existCount - 1
				window[lC] = existCount
				// 如果该字符的存在个数 ＜ 要求个数，valid要减少，不命中下次的左滑
				if existCount < tMap[lC] {
					valid--
				}
			}
		}
	}
	return result
}

func leftNeedShrink(valid int, tMapLen int) bool {
	// 有效字符数 = t的要求字符数量，即 已经找到t的覆盖子串了
	return valid == tMapLen
}
