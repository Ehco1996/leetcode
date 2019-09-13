/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一个树的子树
 *
 * https://leetcode-cn.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (41.58%)
 * Likes:    113
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 18.3K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 * 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s
 * 也可以看做它自身的一棵子树。
 *
 * 示例 1:
 * 给定的树 s:
 *
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 给定的树 t：
 *
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 *
 * 示例 2:
 * 给定的树 s：
 *
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 * ⁠   /
 * ⁠  0
 *
 *
 * 给定的树 t：
 *
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 返回 false。
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSametree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	return s != nil && t != nil && s.Val == t.Val && isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	// 	一个树是另一个树的子树 则

	// 	要么这两个树相等
	// 	要么这个树是左树的子树
	// 	要么这个树是右树的子树

	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}

	return isSametree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

