#
# @lc app=leetcode.cn id=106 lang=python3
#
# [106] 从中序与后序遍历序列构造二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def buildTree(self, inorder, postorder):
        """
        // 先把序列拆分成root和左右两颗子树，然后反过来递归
        // 			in order
        // +------------+------+-------------+
        // | left child | root | right child |
        // +------------+------+-------------+

        // 			post order
        // +------------+-------------+------+
        // | left child | right child | root |
        // +------------+-------------+------+
        """
        if len(inorder) == 0 and len(postorder) == 0:
            return

        root_val = postorder[-1]
        root = TreeNode(root_val)
        pos = inorder.index(root_val)
        if len(inorder) == 1 and len(postorder) == 1:
            return root
        root.left = self.buildTree(inorder[:pos], postorder[:pos])
        root.right = self.buildTree(inorder[pos + 1 :], postorder[pos:-1])
        return root


# @lc code=end

