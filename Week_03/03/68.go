/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
祖先的定义：若节点p在节点root的左（右）子树中，或p=root，则称root为p的祖先。

最近公共祖先的定义：设节点root为节点p，q的某个公共祖先，若其左子节点root.left和右子节点root.right都不是p，q的公共祖先，则称为root是”最近公共祖先“。

根据上述定义，若root是p，q的最近公共祖先，则只可能为以下情况之一：

1. p和q在root的子树中，且分列root的异侧（即分别在左右子树中）
2. p = root，且q在root的左子树或右子树中；
3. q = root，且p在root的左子树或右子树中；

通过递归对二叉树进行后序遍历，当遇到节点p或q时返回，从底到顶回溯，当节点p，q在节点root的异端时，节点root即为最近公共祖先，则向上返回root。

递归解析：

1. 终止条件
	1. 当越过叶子节点，则直接返回null
	2. 当root等于p或q，则直接返回root
2. 递推工作
	1. 开始递归左子树，返回值记为left
	2. 开始递归右子树，返回值记为right
3. 返回值：根据left和right，分为4种情况：
	1. 当left和right同时为空时，说明root的左/右子树中都不包含p、q，返回null
	2. 当left和right同时不为空时，说明p、q分列在root的异侧，因此root为最近公共祖先，返回root
	3. 当left为空，right不为空时，p、q都不在root的左子树中，直接返回right，具体可分为两种情况：
		1. p、q其中一个在root的右子树中，此时right指向p（假设为p）
		2. p、q两节点都在在root的右子树中，此时的right指向最近公共祖先节点
	4. 当left不为空，right为空时，与3同理
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if Left == nil {
		return right
	}

	if right == nil {
		return Left
	}

	return root
}
