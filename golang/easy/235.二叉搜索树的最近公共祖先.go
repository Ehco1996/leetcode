// package golang
/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (60.45%)
 * Likes:    153
 * Dislikes: 0
 * Total Accepted:    23.1K
 * Total Submissions: 38.2K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * 输出: 6
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 *
 *
 * 示例 2:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
 *
 *
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func isAnc(root,p *TreeNode)bool{
	if root == p{
		return true
	}
	if root==nil{
		return false
	}
	return isAnc(root.Left,p) || isAnc(root.Right,p)



}


 func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	 // 暴力
	anc := root

	for root!=nil{
		if isAnc(root.Left,p) && isAnc(root.Left,q){
			anc = root.Left
			root = root.Left
		}else if isAnc(root.Right,p) && isAnc(root.Right,q){
			anc = root.Right
			root = root.Right
		}else{
			break
		}
	}
	return anc

}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 思路：此题其实是一道后序遍历的特性题，
	// 但是由于树被限定为了 二叉搜索树，此题就变得很简单了。
	// 因为由于二叉搜索树具有很强的特性，每个节点的右孩子均大于左孩子，
	// 所以我们要找他们最先出现的祖先节点，这意味着他们一定在某一个节点的两边，
	// 对于这个节点就是他们最先出现的祖先节点，此节点满足二叉搜索树种的特点，
	// 只需要找到某个节点中同时满足大于p节点的值还满足小于q节点的值即可（因为题目规定p<q）
	if root == nil{
		return nil
	}

	// switch
	if p.Val > q.Val{
		p,q = q,p
	}

	if root.Val>=p.Val && root.Val<=q.Val{
		return root
	}

	if root.Val>=p.Val && root.Val>q.Val && root.Left!=nil{
		return lowestCommonAncestor(root.Left,p,q)
	}

	if root.Val<p.Val && root.Val<=q.Val && root.Right!=nil{
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root

}


