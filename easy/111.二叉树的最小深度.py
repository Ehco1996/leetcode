#
# @lc app=leetcode.cn id=111 lang=python3
#
# [111] 二叉树的最小深度
#
# https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
#
# algorithms
# Easy (38.73%)
# Likes:    128
# Dislikes: 0
# Total Accepted:    19.7K
# Total Submissions: 50.7K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，找出其最小深度。
#
# 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
#
# 说明: 叶子节点是指没有子节点的节点。
#
# 示例:
#
# 给定二叉树 [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
# 返回它的最小深度  2.
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0

        stack = [(1, root)]
        cur_depth = 1

        while stack:
            depth, tree = stack.pop()
            cur_depth = max(cur_depth, depth)

            if not tree.left and not tree.right:
                return cur_depth
            else:
                tree.left and stack.insert(0, (cur_depth + 1, tree.left))
                tree.right and stack.insert(0, (cur_depth + 1, tree.right))
