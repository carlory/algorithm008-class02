/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rlen, clen := len(matrix), len(matrix[0])
	res := make([]int, 0, rlen*clen)
	seen := make([][]bool, rlen)
	for i := range seen {
		seen[i] = make([]bool, clen)
	}
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	var r, c, di int
	for i := 0; i < rlen*clen; i++ {
		res = append(res, matrix[r][c])
		seen[r][c] = true
		cr := r + dr[di]
		cc := c + dc[di]
		if cr >= 0 && cr < rlen && cc >= 0 && cc < clen && !seen[cr][cc] {
			r = cr
			c = cc
		} else {
			di = (di + 1) % 4
			r += dr[di]
			c += dc[di]
		}
	}
	return res
}
