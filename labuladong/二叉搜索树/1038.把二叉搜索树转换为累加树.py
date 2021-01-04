#
# @lc app=leetcode.cn id=1038 lang=python3
#
# [1038] 把二叉搜索树转换为累加树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def bstToGst(self, root):
        self.sum = 0
        self.traverse(root)
        return root

    def traverse(self, root):
        if not root:
            return
        self.traverse(root.right)
        self.sum += root.val
        root.val = self.sum
        self.traverse(root.left)


# @lc code=end

