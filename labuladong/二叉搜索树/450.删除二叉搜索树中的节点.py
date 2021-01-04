#
# @lc app=leetcode.cn id=450 lang=python3
#
# [450] 删除二叉搜索树中的节点
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def deleteNode(self, root, key: int):
        if not root:
            return
        if root.val == key:
            # case1 自己是叶子节点
            if root.left is None and root.right is None:
                return
            # case2 只有一个节点
            if root.left is None:
                return root.right
            if root.right is None:
                return root.left
            # case3 两个节点都有值，找到右边最小的来接替自己，或者左边最大的
            min_node = self.get_min_node(root.right)
            root.val = min_node.val
            root.right = self.deleteNode(root.right, min_node.val)
        if root.val > key:
            root.left = self.deleteNode(root.left, key)
        if root.val < key:
            root.right = self.deleteNode(root.right, key)
        return root

    def get_min_node(self, node):
        while node.left is not None:
            node = node.left
        return node


# @lc code=end

