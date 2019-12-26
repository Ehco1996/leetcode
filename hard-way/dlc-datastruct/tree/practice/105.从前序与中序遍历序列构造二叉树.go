/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (62.64%)
 * Likes:    296
 * Dislikes: 0
 * Total Accepted:    35.3K
 * Total Submissions: 55.8K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 先把序列拆分成root和左右两颗子树，然后反过来递归
	// 			in order
	// +------------+------+-------------+
	// | left child | root | right child |
	// +------------+------+-------------+

	// 			pre order
	// +------------+-------------+------+
	// | root| left child | right child  |
	// +------------+-------------+------+

	if len(inorder) == 0 && len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	pos := searchIdx(inorder, rootVal)
	if len(inorder) == 1 && len(preorder) == 1 {
		return root
	}
	root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
	root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
	return root
}

// @lc code=end

