#
# @lc app=leetcode.cn id=107 lang=python3
#
# [107] 二叉树的层次遍历 II
#
# https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
#
# algorithms
# Easy (60.76%)
# Likes:    102
# Dislikes: 0
# Total Accepted:    16.8K
# Total Submissions: 27.6K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
#
# 例如：
# 给定二叉树 [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
#
# 返回其自底向上的层次遍历为：
#
# [
# ⁠ [15,7],
# ⁠ [9,20],
# ⁠ [3]
# ]
#
#
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        from collections import defaultdict

        if not root:
            return []

        h = defaultdict(list)

        cur_depth = 0
        max_depth = 0
        stack = [(cur_depth, root)]

        while stack:
            cur_depth, tree = stack.pop()
            if tree:
                h[cur_depth].append(tree.val)
                max_depth = max(max_depth, cur_depth)
                stack.append((cur_depth + 1, tree.right))
                stack.append((cur_depth + 1, tree.left))

        res = []
        print(h)
        for i in range(0, max_depth + 1):
            res.insert(0, h[i])
        return res

