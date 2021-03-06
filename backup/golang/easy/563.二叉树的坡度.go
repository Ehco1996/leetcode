package easy

/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 *
 * https://leetcode-cn.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (49.11%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 10.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，计算整个树的坡度。
 *
 * 一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
 *
 * 整个树的坡度就是其所有节点的坡度之和。
 *
 * 示例:
 *
 *
 * 输入:
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * 输出: 1
 * 解释:
 * 结点的坡度 2 : 0
 * 结点的坡度 3 : 0
 * 结点的坡度 1 : |2-3| = 1
 * 树的坡度 : 0 + 0 + 1 = 1
 *
 *
 * 注意:
 *
 *
 * 任何子树的结点的和不会超过32位整数的范围。
 * 坡度的值不会超过32位整数的范围。
 *
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
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * -1
	}
}

func findTilt(root *TreeNode) int {
	// 遍历
	tiltSum := 0

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := traverse(node.Left)
		right := traverse(node.Right)
		tiltSum += abs(left - right)
		return left + right + node.Val
	}
	traverse(root)
	return tiltSum
}
