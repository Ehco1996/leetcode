#
# @lc app=leetcode.cn id=230 lang=python3
#
# [230] 二叉搜索树中第K小的元素
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root, k) -> int:
        """
        利用二叉树的前序遍历(左中右)是升序来找
        """
        self.k = k
        self.res = 0
        self.traverse(root)
        return self.res

    def traverse(self, root):
        if not root:
            return
        self.traverse(root.left)
        self.k -= 1
        if self.k == 0:
            self.res = root.val
            return
        self.traverse(root.right)


# @lc code=end

