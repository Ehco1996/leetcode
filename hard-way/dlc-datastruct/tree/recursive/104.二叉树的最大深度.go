/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (71.34%)
 * Likes:    415
 * Dislikes: 0
 * Total Accepted:    105.7K
 * Total Submissions: 147.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
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

// @lc code=end

