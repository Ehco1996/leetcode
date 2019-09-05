// package golang

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (59.44%)
 * Likes:    150
 * Dislikes: 0
 * Total Accepted:    14.4K
 * Total Submissions: 24.1K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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




func binaryTreePaths(root *TreeNode) []string {
	paths := []string{}

	if root == nil{
		return paths
	}

	if root.Left == nil && root.Right ==nil{
		paths = append(paths,strconv.Itoa(root.Val))
	}

	if root.Left!= nil{
		for _,path:=range binaryTreePaths(root.Left){
			s:= strconv.Itoa(root.Val)
			s += "->"
			s += path
			paths = append(paths,s)
		}
	}

	if root.Right!= nil{
		for _,path:=range binaryTreePaths(root.Right){
			s:= strconv.Itoa(root.Val)
			s += "->"
			s += path
			paths = append(paths,s)
		}
	}

	return paths

}
