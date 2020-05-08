/*
我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/chou-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "math"

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		dp[i] = int(math.Min(float64(dp[p2]*2), math.Min(float64(dp[p3]*3), float64(dp[p5]*5))))
		if dp[i] == dp[p2]*2 {
			p2++
		}

		if dp[i] == dp[p3]*3 {
			p3++
		}

		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n-1]
}
