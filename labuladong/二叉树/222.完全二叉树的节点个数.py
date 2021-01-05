#
# @lc app=leetcode.cn id=222 lang=python3
#
# [222] 完全二叉树的节点个数
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def countNodes(self, root) -> int:
        """
        完全二叉树：
        除了最底层节点可能没填满外，
        其余每层节点数都达到最大值，
        并且最下面一层的节点都集中在该层最左边的若干位置。
        若最底层为第 h 层，则该层包含 1~ 2h 个节点。
        """
        if not root:
            return 0
        return 1 + self.countNodes(root.left) + self.countNodes(root.right)

    def countFullNodes(self, root):
        """
        满二叉树，既每层都是满的二叉树
        他的节点个数为 2^h -1 个 （h为层数）
        """
        if not root:
            return 0
        h = 0
        while root.left:
            h += 1
            root = root.left
        return 2 ** h - 1


# @lc code=end

