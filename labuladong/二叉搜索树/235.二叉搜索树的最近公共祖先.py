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
        if root is None or p == root or q == root:
            return root

        # 找到p，q在左树和右树的最近的公共祖先
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)

        # 每次递归都有以下几种情况

        # 1. 如果在左树和右树中都能找到公共祖先，那说明root.left =p root.right=q 既root就是pq的公共祖先
        if left and right:
            return root

        # 2. 如果p和q都不在以root为根的树中，说明没公共祖先
        if not left and not right:
            return None

        # 3 如果p和q只有一个存在于root为根的树中，那么找到的那个节点就是公共祖先
        if left:
            return left
        return right


# @lc code=end

