package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoofoothebar", []string{"foo", "bar", "the"}))
}

func findSubstring(s string, words []string) []int {
	return Slide(s, words)
}

func Slide(s string, words []string) []int {
	result := make([]int, 0, 1)
	length := len(s)
	size := len(words)
	wordSize := len(words[0])
	window := ""
	windowMaxSize := size * wordSize
	left := 0
	right := 0

	for left < length && right < length {
		rightC := string(s[right])
		right++
		// 更新窗口
		window = window + rightC
		if len(window) == windowMaxSize {
			wordsMap := make(map[string]int)
			for _, str := range words {
				wordsMap[str] = wordsMap[str] + 1
			}
			// 按照wordSize分割window，判断是否命中wordsMap
			for j := 0; j <= windowMaxSize-wordSize; j = j + wordSize {
				pieceOfWindow := window[j : j+wordSize]
				count, ok := wordsMap[pieceOfWindow]
				if ok && count > 0 {
					newValue := wordsMap[pieceOfWindow] - 1
					if newValue == 0 {
						// 已经匹配完了
						delete(wordsMap, pieceOfWindow)
					} else {
						wordsMap[pieceOfWindow] = newValue
					}
				} else {
					// 优化点1：没必要再比对了
					break
				}
			}
			// 如果wordsMap的size == 0，说明window命中所有word
			if len(wordsMap) == 0 {
				result = append(result, left)
			}

		}

		for LeftNeedShrink(left, right, windowMaxSize) {
			left++
			// 更新窗口
			window = window[1:]
		}
	}

	return result
}

func LeftNeedShrink(left int, right int, size int) bool {
	return right-left >= size
}
