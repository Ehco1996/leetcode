package easy

/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (42.08%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    4.9K
 * Total Submissions: 11.5K
 * Testcase Example:  '[1,null,2,2]'
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 *
 *
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 * 返回[2].
 *
 * 提示：如果众数超过1个，不需考虑输出顺序
 *
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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

// func findWeight(node *TreeNode, weight, val int) int {
// 	if node == nil || node.Val != val {
// 		return weight
// 	}
// 	weight++
// 	left := findWeight(node.Left, weight, val)
// 	right := findWeight(node.Right, weight, val)
// 	return weight + right + left
// }


func findMode(root *TreeNode) []int {
	// 先序遍历
	var result []int
	var maxCount = 0
	var commonNum int

	count := 0

	visit := func(val int) {
		if val == commonNum {
			count++
		} else {
			commonNum = val
			count = 1
		}
		if count > maxCount {
			maxCount = count
			result = []int{val}
		} else if count == maxCount {
			result = append(result, val)
		}
	}

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		visit(node.Val)
		inOrder(node.Right)
	}
	inOrder(root)

	return result
}
