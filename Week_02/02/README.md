# 刷题总结

## [350. 求两个数组的交集II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)

```
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
```

### 数组元素不可排序

在不看题目的示例的前提下，我们是无法确定要求交集的两个数组元素的类型的。

通用的解法如下：

1. 在哈希表中记录一个数组中存在的数字和对应的出现次数；
2. 遍历第二个数组，检查数字在哈希表中是否存在，如果存在其计数结果大于0，则将该数组添加到答案并减少哈希表中的计数。

### 数组元素可排序

若数组元素有序，可采用下面解法：

1. 初始化指针i, j, k
2. 指针i指向nums1, 指针j指向nums2
3. 如果`nums1[i] < nums2[j]`, 则指针i向后移动一位
4. 如果`nums1[i] > nums2[j]`, 则指针j向后移动一位
5. 如果`nums1[i] == nums2[j]`, 将元素拷贝到`nums1[k]`, 并将i, j, k 分别向后移动一位
6. 返回结果`nums1[:k]`

当然，若数组元素无序，可使用排序算法对数组进行原地排序，然后使用上面算法解答。

## TODO List

- [ ] 350. 求两个数组的交集II的进阶解法
- [ ] [392. 判断子序列的动态规划解法](https://leetcode-cn.com/problems/is-subsequence/solution/cpan-duan-zi-xu-lie-hou-xu-tiao-zhan-by-zzzzzz-5/)
- [ ] [17.09 第k个数](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)的堆解法和动态规划解法 
