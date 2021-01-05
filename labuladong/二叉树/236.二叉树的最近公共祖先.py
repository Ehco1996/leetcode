#
# @lc app=leetcode.cn id=236 lang=python3
#
# [236] 二叉树的最近公共祖先
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
        if root is None or p == root or q == root:
            return root
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)

        # 每次递归都有以下几种情况
        # 1 如果p和q都在以root为根的树中，那么left和right一定分别是p和q
        if left and right:
            return root

        # 2 如果p和q都不在以root为根的树中，直接返回null。
        if not left and not right:
            return None

        # 3 如果p和q只有一个存在于root为根的树中，函数返回该节点。
        if left:
            return left
        return right


# @lc code=end

