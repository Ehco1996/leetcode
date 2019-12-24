/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (59.72%)
 * Likes:    332
 * Dislikes: 0
 * Total Accepted:    65.4K
 * Total Submissions: 109K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
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

type Node struct {
	node  *TreeNode
	depth int
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	node := &Node{root, 0}
	queue := []*Node{node}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		depth := node.depth
		if len(res) < depth+1 {
			res = append(res, make([]int, 0))
		}
		res[depth] = append(res[depth], node.node.Val)

		if node.node.Left != nil {
			queue = append(queue, &Node{node.node.Left, depth + 1})
		}

		if node.node.Right != nil {
			queue = append(queue, &Node{node.node.Right, depth + 1})
		}
	}
	return res
}

// @lc code=end

