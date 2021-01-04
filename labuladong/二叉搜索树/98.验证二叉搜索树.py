#
# @lc app=leetcode.cn id=98 lang=python3
#
# [98] 验证二叉搜索树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def isValidBST(self, root) -> bool:
        return self.helper(root, None, None)

    def helper(self, root, min_node, max_node):
        if not root:
            return True
        if min_node is not None and root.val <= min_node.val:
            return False
        if max_node is not None and root.val >= max_node.val:
            return False
        return self.helper(root.left, min_node, root) and self.helper(
            root.right, root, max_node
        )
# @lc code=end

