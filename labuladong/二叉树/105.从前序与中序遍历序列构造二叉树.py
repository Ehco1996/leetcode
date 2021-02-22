#
# @lc app=leetcode.cn id=105 lang=python3
#
# [105] 从前序与中序遍历序列构造二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def test(self, preorder=[3, 9, 20, 15, 7], inorder=[9, 3, 15, 20, 7]):
        return self.buildTree(preorder, inorder)

    def buildTree(self, preorder, inorder):
        """
        先把序列拆分成root和左右两颗子树，然后反过来递归
                    in order
        +------------+------+-------------+
        | left child | root | right child |
        +------------+------+-------------+

                    pre order
        +------------+-------------+------+
        | root| left child | right child  |
        +------------+-------------+------+
        """
        if len(inorder) == 0 and len(preorder) == 0:
            return None

        root_val = preorder[0]
        root = TreeNode(root_val)
        pos = inorder.index(root_val)
        if len(inorder) == 1 and len(preorder) == 1:
            return root
        root.left = self.buildTree(preorder[1 : pos + 1], inorder[:pos])
        root.right = self.buildTree(preorder[pos + 1 :], inorder[pos + 1 :])
        return root


# @lc code=end
