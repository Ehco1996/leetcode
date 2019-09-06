// package golang

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (50.97%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    9.9K
 * Total Submissions: 19.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 *
 * 示例：
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
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

func helper(root *TreeNode, sum *int, flag int) {
    if root == nil {
        return
    }
    helper(root.Left, sum, 1)
    if root.Left == nil && root.Right == nil && flag == 1{
        *sum += root.Val
    }
    helper(root.Right, sum, 2)
}



func sumOfLeftLeaves(root *TreeNode) int {
    var sum int
    // flag 表是左节点还是右节点，1 左，2 右
    helper(root, &sum, 0)
    return sum
}




//     3
//   9   20
//      15  7
//         1 2
