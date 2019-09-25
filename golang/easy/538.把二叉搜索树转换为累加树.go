package easy

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 *
 * https://leetcode-cn.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (56.10%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 14.7K
 * Testcase Example:  '[5,2,13]'
 *
 * 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater
 * Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
 *
 * 例如：
 *
 *
 * 输入: 二叉搜索树:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * 输出: 转换为累加树:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
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
func convertBST(root *TreeNode) *TreeNode {
	// 从数的右下角开始累加
	sum := 0
	visit := func(node *TreeNode) {
		node.Val += sum
		sum = node.Val
	}

	var lastOrder func(node *TreeNode)
	lastOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		lastOrder(node.Right)
		visit(node)
		lastOrder(node.Left)
	}
	lastOrder(root)
	return root
}
