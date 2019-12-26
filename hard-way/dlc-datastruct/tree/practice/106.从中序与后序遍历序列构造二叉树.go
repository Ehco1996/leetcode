/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (65.11%)
 * Likes:    136
 * Dislikes: 0
 * Total Accepted:    19.9K
 * Total Submissions: 30.2K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

func searchIdx(nums []int, target int) int {
	for idx, num := range nums {
		if num == target {
			return idx
		}
	}
	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 先把序列拆分成root和左右两颗子树，然后反过来递归
	// 			in order
	// +------------+------+-------------+
	// | left child | root | right child |
	// +------------+------+-------------+

	// 			post order
	// +------------+-------------+------+
	// | left child | right child | root |
	// +------------+-------------+------+

	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	pos := searchIdx(inorder, rootVal)
	if len(inorder) == 1 && len(postorder) == 1 {
		return root
	}
	root.Left = buildTree(inorder[:pos], postorder[:pos])
	root.Right = buildTree(inorder[pos+1:], postorder[pos:len(postorder)-1])
	return root
}

// @lc code=end

