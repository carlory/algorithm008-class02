/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	deque := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
	}

	maxs := make([]int, 0, len(nums)-k+1)
	maxs = append(maxs, deque[0])
	for i := k; i < len(nums); i++ {
		if nums[i-k] == deque[0] {
			deque = deque[1:]
		}
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
		maxs = append(maxs, deque[0])
	}
	return maxs
}
