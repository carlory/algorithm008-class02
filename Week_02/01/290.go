/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
*/
package main

import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		word, ok1 := m1[pattern[i]]
		char, ok2 := m2[words[i]]
		if (ok1 && word != words[i]) || (ok2 && char != pattern[i]) {
			return false
		}

		if !ok1 {
			m1[pattern[i]] = words[i]
		}

		if !ok2 {
			m2[words[i]] = pattern[i]
		}
	}

	return true
}
