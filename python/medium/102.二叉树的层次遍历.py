#
# @lc app=leetcode.cn id=102 lang=python3
#
# [102] 二叉树的层次遍历
#
# https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
#
# algorithms
# Medium (56.40%)
# Likes:    215
# Dislikes: 0
# Total Accepted:    29.2K
# Total Submissions: 51.6K
# Testcase Example:  '[3,9,20,null,null,15,7]'
#
# 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
#
# 例如:
# 给定二叉树: [3,9,20,null,null,15,7],
#
# ⁠   3
# ⁠  / \
# ⁠ 9  20
# ⁠   /  \
# ⁠  15   7
#
#
# 返回其层次遍历结果：
#
# [
# ⁠ [3],
# ⁠ [9,20],
# ⁠ [15,7]
# ]
#
#
#
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __repr__(self):
        return str(self.val)


class Solution:
    # def levelOrder(self, root: TreeNode) -> List[List[int]]:
    #     if not root:
    #         return []
    #     res = []

    #     depth = 1
    #     stack = [(1, root)]
    #     while stack:
    #         depth, node = stack.pop()
    #         if len(res) < depth:
    #             res.append([])
    #         res[depth - 1].append(node.val)
    #         node.right and stack.append((depth + 1, node.right))
    #         node.left and stack.append((depth + 1, node.left))
    #     return res

    def levelOrder(self, root):
        if not root:
            return []
        ans, level = [], [root]
        while level:
            ans.append([node.val for node in level])
            temp = []
            for node in level:
                temp.extend([node.left, node.right])
            level = [leaf for leaf in temp if leaf]
            print(temp, level)
        return ans

