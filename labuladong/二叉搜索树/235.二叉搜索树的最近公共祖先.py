#
# @lc app=leetcode.cn id=235 lang=python3
#
# [235] 二叉搜索树的最近公共祖先
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def lowestCommonAncestor(self, root, p, q):
        """
        为由于二叉搜索树具有很强的特性，每个节点的右孩子均大于左孩子，
        所以我们要找他们最先出现的祖先节点，这意味着他们一定在某一个节点的两边
        只需要找到某个节点中同时满足大于p节点的值还满足小于q节点的值即可
        """
        if not root:
            return

        # 保证p<q
        if p.val > q.val:
            p, q = q, p

        # 每次递归都有以下几种情况
        # 1 如果当前节点大于p小于q 就找到了
        if root.val >= p.val and root.val <= q.val:
            return root

        # 2 当前值大于p了，就要往左边找，变小
        if root.val >= p.val and root.left:
            return self.lowestCommonAncestor(root.left, p, q)

        # 3 当前值小于q了，就要往右边找，变大
        if root.val <= q.val and root.right:
            return self.lowestCommonAncestor(root.right, p, q)
        return root


# @lc code=end

