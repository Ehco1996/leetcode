/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (56.36%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    20.5K
 * Total Submissions: 36.4K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 输出: 3
 * 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
 *
 *
 * 示例 2:
 *
 * 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 输出: 5
 * 解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉树中。
 *
 *
 */
package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	ans := root

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		// dequeue
		queue = queue[1:]
		ans = node

		if node.Left != nil &&
			findNode(q.Val, node.Left) && findNode(p.Val, node.Left) {
			queue = append(queue, node.Left)
		}
		if node.Right != nil &&
			findNode(q.Val, node.Right) && findNode(p.Val, node.Right) {
			queue = append(queue, node.Right)
		}
	}
	return ans
}

func findNode(val int, node *TreeNode) bool {
	if node == nil {
		return false
	} else if node.Val == val {
		return true
	}
	return findNode(val, node.Left) || findNode(val, node.Right)
}
