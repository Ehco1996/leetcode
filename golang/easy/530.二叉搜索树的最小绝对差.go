/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (53.90%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    5.7K
 * Total Submissions: 10.5K
 * Testcase Example:  '[1,null,3,2]'
 *
 * 给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。
 *
 * 示例 :
 *
 *
 * 输入:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * 输出:
 * 1
 *
 * 解释:
 * 最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 *
 *
 * 注意: 树中至少有2个节点。
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

func getNowMin(now, val int) int {
	if now <= 0 || val < now {
		return val
	} else {
		return now
	}
}

func getMinimumDifference(root *TreeNode) int {
	// 搜索树 前序遍历 出来是一个递增序列

	last := -1
	nowMin := -1

	visit := func(node *TreeNode) {
		if last == -1 {
		} else if node.Val > last {
			nowMin = getNowMin(nowMin, node.Val-last)
		} else {
			nowMin = getNowMin(nowMin, last-node.Val)
		}
		last = node.Val
	}

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		visit(node)
		inOrder(node.Right)
	}
	inOrder(root)
	return nowMin
}

