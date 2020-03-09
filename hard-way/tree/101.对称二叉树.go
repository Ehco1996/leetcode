/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (49.10%)
 * Likes:    544
 * Dislikes: 0
 * Total Accepted:    78.9K
 * Total Submissions: 160.2K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 * 说明:
 *
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
func isSymmetric(root *TreeNode) bool {
	// 对称以为着每个节点的左树和右树是一样的
	var helper func(l, r *TreeNode) bool
	helper = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l != nil && r != nil && l.Val == r.Val {
			return helper(l.Left, r.Right) && helper(r.Left, l.Right)
		}
		return false
	}
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

// @lc code=end

