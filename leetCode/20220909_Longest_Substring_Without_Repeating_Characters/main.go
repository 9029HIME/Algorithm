package main

/*
1.题目
无重复字符的最长子串


2.描述
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。


3.示例
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


4.思路：
20220909_Longest_Substring_Without_Repeating_Characters/无重复字符串的最长子串长度.jpg
*/

func main() {
	println(lengthOfLongestSubstring("dvdf"))       //3
	println(lengthOfLongestSubstring("QWEQKOWLDP")) //8
	println(lengthOfLongestSubstring("aaaa"))       // 1
	println(lengthOfLongestSubstring("pwwkew"))     // 3
	println(lengthOfLongestSubstring("ckilbkd"))    //5
	println(lengthOfLongestSubstring("aab"))        //2
	println(lengthOfLongestSubstring("nfpdmpi"))    //5
	println(lengthOfLongestSubstring("tmmzuxt"))    // 5 if tmmzuxtm? 因为他，必须要删除新i前面的数据，否则最后面的t判断存在
	println(lengthOfLongestSubstring("abcabcbb"))   //3
}

func lengthOfLongestSubstring(s string) int {
	capacity := len(s)
	if capacity <= 1 {
		return capacity
	}
	set := make(map[string]int, capacity)
	result := 0
	tmpResult := 0
	i := 0
	j := 0

	for i < capacity && j < capacity {
		set[string(s[i])] = i

		next := string(s[j])
		if index, ok := set[string(next)]; ok && i != j {
			// 值存在，需要移位i
			for x := i; x <= index; x++ {
				deleteKey := string(s[x])
				delete(set, deleteKey)
			}
			i = index + 1
		} else {
			set[string(next)] = j
			tmpResult = j - i + 1
			if tmpResult > result {
				result = tmpResult
			}
			j++
		}
	}
	return result
}
