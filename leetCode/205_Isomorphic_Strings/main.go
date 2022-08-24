package main

/**
1.题目
同构字符串(abb = lee,paper = title)

2.描述
给定两个字符串s和t，判断它们是否是同构的。
如果s中的字符可以按某种映射关系替换得到t，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

3.示例
输入：s = "egg", t = "add"
输出：true
输入：s = "foo", t = "bar"
输出：false

4.思路
首先肯定是得循环两个字符串的，能否用哈希表的特性来实现呢？
比较字符串长度，不相等必定返回false（第一次校验）
既然b=e，那么可以利用哈希表的特性，存储k=a,v=e的数据。当遍历到下一个b的时候，得到【相同下标位置】的【另一个字符串】的字符，看看是不是e，如果不是，直接返回false。

5.注意
badc baba这个用例，因为badc字符从头到尾在map都找不到对应的baba字符，所以会返回true。应该重复遍历一遍，此时遍历的起点是baba。
如果用Java写的话，可以先判断badc字符是否在map存在键，如果键不存在则判断baba字符在map里是否值，如果值存在则返回false。如果，如果值不存在则放入键值对。如果键存在则比较值是否相等（和上面的思路一样）。但是Go语言的map好像没有【快速找值是否存在】的api，Java本身是冗余了values集合的，所以Go只能基于上面的思路，重新相反遍历了。

*/

func main() {
	print(isIsomorphic("paper", "title"))
}

func isIsomorphic(s string, t string) bool {
	sLength := len(s)
	tLength := len(t)
	if sLength != tLength {
		return false
	}

	var map1 map[rune]rune = make(map[rune]rune)
	var map2 map[rune]rune = make(map[rune]rune)
	for i, v := range s {
		sChar := v
		tChar := t[i]
		if existTChar := map1[sChar]; existTChar == 0 {
			map1[sChar] = rune(tChar)
		} else if rune(tChar) != existTChar {
			// 既然你存在，就该比较一下是否相等了，如果不相等必定为false
			return false
		}

		if existSChar := map2[rune(tChar)]; existSChar == 0 {
			map2[rune(tChar)] = rune(sChar)
		} else if rune(sChar) != existSChar {
			// 既然你存在，就该比较一下是否相等了，如果不相等必定为false
			return false
		}

	}
	return true
}
