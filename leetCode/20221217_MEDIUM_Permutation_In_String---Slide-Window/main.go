package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[string]int)
	for _, c := range s1 {
		if count, ok := m[string(c)]; ok {
			m[string(c)] = count + 1
		} else {
			m[string(c)] = 1
		}
	}
	window := make(map[string]int)

	size := len(s1)
	length := len(s2)
	if size > length {
		return false
	}
	left := 0
	right := 0
	// 窗口内有多少个字符满足要求了
	valid := 0
	for right < length && left < length {
		// 获取right资源
		rightC := string(s2[right])

		right++

		// 更新窗口数据
		if count, ok := m[rightC]; ok {
			window[rightC] = window[rightC] + 1
			// 代表这个字符已经满足要求了
			if window[rightC] == count {
				valid++
			}
		}

		// 左边窗口开始滑（需要命中条件）
		for LeftNeedShrink(left, right, size) {
			if valid == len(m) {
				return true
			}
			// 获取left资源
			leftC := string(s2[left])

			left++

			// 更新窗口数据
			if count, ok := m[leftC]; ok {
				// 代表这个字符原本是满足要求的，结果左指针滑动导致不满足了
				if window[leftC] == count {
					valid--
				}
				window[leftC] = window[leftC] - 1
			}
		}
	}
	return false
}

func LeftNeedShrink(left int, right int, size int) bool {
	// 发现窗口变大了，需要收缩left，保持窗口固定大小
	if right-left+1 > size {
		return true
	}
	return false
}
