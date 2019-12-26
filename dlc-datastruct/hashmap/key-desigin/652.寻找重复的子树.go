/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (50.03%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 7.1K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 两棵树重复是指它们具有相同的结构以及相同的结点值。
 *
 * 示例 1：
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   2   4
 * ⁠      /
 * ⁠     4
 *
 *
 * 下面是两个重复的子树：
 *
 * ⁠     2
 * ⁠    /
 * ⁠   4
 *
 *
 * 和
 *
 * ⁠   4
 *
 *
 * 因此，你需要以列表的形式返回上述重复子树的根结点。
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

import "strconv"

func treePath(node *TreeNode) string {

	res := ""
	var walker func(node *TreeNode)
	walker = func(node *TreeNode) {
		if node == nil {
			res += "*"
			return
		}
		res += strconv.Itoa(node.Val)
		walker(node.Left)
		walker(node.Right)
	}

	walker(node)
	return res
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	h := make(map[string][]*TreeNode)

	var walker func(node *TreeNode)
	walker = func(node *TreeNode) {
		if node == nil {
			return
		}
		key := treePath(node)
		h[key] = append(h[key], node)
		walker(node.Left)
		walker(node.Right)
	}
	walker(root)

	for _, nodes := range h {
		if len(nodes) > 1 {
			result = append(result, nodes[0])
		}
	}
	return result
}

// @lc code=end

