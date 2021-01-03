#
# @lc app=leetcode.cn id=652 lang=python3
#
# [652] 寻找重复的子树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findDuplicateSubtrees(self, root):
        """用后续遍历找出当前节点的树的字符串表现形式"""
        from collections import defaultdict

        self.visited = defaultdict(int)
        self.res = []
        self.traverse(root)
        return self.res

    def traverse(self, root):
        if not root:
            return "#"
        left = self.traverse(root.left)
        right = self.traverse(root.right)
        sub_tree = f"{left},{right},{root.val}"
        cnt = self.visited.get(sub_tree)
        if cnt == 1:
            self.res.append(root)
        self.visited[sub_tree] += 1
        return sub_tree


# @lc code=end

