/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	sum := []byte{}
	i, j, carry := len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 {
		var v1, v2 int
		if i >= 0 {
			v1 = int(num1[i] - '0')
		}
		if j >= 0 {
			v2 = int(num2[j] - '0')
		}
		v := v1 + v2 + carry
		sum = append([]byte(strconv.Itoa(v%10)), sum...)
		carry = v / 10
	}
	if carry > 0 {
		sum = append([]byte("1"), sum...)
	}
	return string(sum)
}
