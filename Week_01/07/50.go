/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
*/
package main

import "fmt"

func firstUniqChar(s string) byte {
	list := make([]int, 26)
	for _, r := range s {
		list[r-'a']++
	}
	for i, r := range s {
		if list[r-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	s := "a"
	for i, v := range s {
		fmt.Printf("%T, %T\n", s[i], v)
	}
}
