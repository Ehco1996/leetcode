/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (49.13%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    42.5K
 * Total Submissions: 85.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
 *
 *
 * 示例 1:
 *
 * 给定二叉树 [3,9,20,null,null,15,7]
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回 true 。
 *
 * 示例 2:
 *
 * 给定二叉树 [1,2,2,3,3,null,null,4,4]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * 返回 false 。
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	dep := 0

	if root == nil {
		return dep
	}
	var helper func(node *TreeNode, cur int)
	helper = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		if cur > dep {
			dep = cur
		}
		helper(node.Left, cur+1)
		helper(node.Right, cur+1)
	}
	helper(root, 1)
	return dep
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(maxDepth(root.Left), maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)

}

// @lc code=end

