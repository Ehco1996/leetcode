/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (52.57%)
 * Likes:    148
 * Dislikes: 0
 * Total Accepted:    31.7K
 * Total Submissions: 59.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层次遍历如下：
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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

func insert(a []int, idx int, val int) []int {
	a = append(a[:idx], append([]int{val}, a[idx:]...)...)
	return a
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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

		if depth%2 == 0 {
			res[depth] = append(res[depth], node.node.Val)
		} else {
			res[depth] = insert(res[depth], 0, node.node.Val)
		}

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

