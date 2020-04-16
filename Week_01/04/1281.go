/*
给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。

示例 1：

输入：n = 234
输出：15
解释：
各位数之积 = 2 * 3 * 4 = 24
各位数之和 = 2 + 3 + 4 = 9
结果 = 24 - 9 = 15

示例 2：

输入：n = 4421
输出：21
解释：
各位数之积 = 4 * 4 * 2 * 1 = 32
各位数之和 = 4 + 4 + 2 + 1 = 11
结果 = 32 - 11 = 21


提示：

1 <= n <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer
*/
package main

import "fmt"

func subtractProductAndSum(n int) int {
	var (
		product = 1
		sum     = 0
	)

	for n != 0 {
		number := n % 10
		product *= number
		sum += number
		n = n / 10
	}

	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(234) == 15)
	fmt.Println(subtractProductAndSum(4421) == 21)
}
