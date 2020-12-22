#
# @lc app=leetcode.cn id=114 lang=python3
#
# [114] 二叉树展开为链表
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        1.将 root 的左子树和右子树拉平。
        2.将左子树作为右子树。
        3.将整个左子树作为右子树。
        """

        if not root:
            return None

        self.flatten(root.left)
        self.flatten(root.right)

        right = root.right
        left = root.left
        root.left = None
        root.right = left

        p = root
        while p.right:
            p = p.right
        p.right = right


# @lc code=end

