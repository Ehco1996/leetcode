#
# @lc app=leetcode.cn id=654 lang=python3
#
# [654] 最大二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def constructMaximumBinaryTree(self, nums):
        """最大值是root节点"""
        return self.build(nums, 0, len(nums) - 1)

    def build(self, nums, left, right):
        import sys

        if left > right:
            return None

        idx = -1
        max_value = sys.maxsize * -1

        # 找到数组中的最大值和最大索引
        for i in range(left, right + 1):
            if max_value < nums[i]:
                idx = i
                max_value = nums[i]

        root = TreeNode(max_value)
        root.left = self.build(nums, left, idx - 1)
        root.right = self.build(nums, idx + 1, right)
        return root


# @lc code=end
