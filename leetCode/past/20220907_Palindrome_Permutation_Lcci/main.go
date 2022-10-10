package main

/*
1.题目
回文排列

2.描述
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。

3.示例
输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

aacccbb
4.思路
需要理解，回文串的充分必要条件是：
	1.最多只能有1个字符出现奇数次
	2.在1的前提下，其他字符出现的次数必须是2的倍数
所以依旧是用哈希表来做这两个判断
*/
func main() {
	print(canPermutePalindrome("aab"))
}

func canPermutePalindrome(s string) bool {
	var counter map[rune]int = make(map[rune]int)
	for _, v := range s {
		count := counter[v]
		counter[v] = count + 1
	}

	if len(counter) == 1 {
		return true
	}
	// 是否已存在奇数字符了
	hasSingle := false
	for _, v := range counter {
		if v%2 != 0 {
			if hasSingle {
				return false
			} else {
				hasSingle = true
			}
		}
	}

	return true
}
