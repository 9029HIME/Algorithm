package main

/**
1.题目
判断子序列

2.描述
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

3.示例
输入：s = "abc", t = "ahbgdc"
输出：true
输入：s = "axc", t = "ahbgdc"
输出：false

4.思路
其实和IDEA里搜索文件名差不多的原理，假设字符串abcdefg，当我搜索cfg的时候，在abcdefg里必须要先出现c，再出现f，然后再出现g，f必须在c后面，g必须在f后面。那么可以采用双指针的解法
a b c d e f g
↑ ↑ ↑

c
↑

a b c d e f g
↑ ↑ ↑ ↑ ↑ ↑

c f g
↑ ↑ ↑

a b c d e f g
↑ ↑ ↑ ↑ ↑ ↑ ↑

c f g
↑ ↑ ↑
如果遍历完整个abcdefg，还没有遍历完cfg的话，就说明cfg并不是abcdefg的子序列


*/
func main() {
	print(isSubsequence("ace", "abcdefg"))
}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	var index int   // 指向长字符串字符的指针
	var success int // 命中段字符串的个数，如果个数=段字符串长度，说明是子序列
	// i是指向短字符串字符的指针
	for i := 0; i < len(s); i++ {
		sV := s[i]
		for index = index; index < len(t); index++ {
			tV := t[index]
			if sV == tV {
				index++
				success++
				break
			}
		}
	}
	if success == len(s) {
		return true
	}

	return false
}
