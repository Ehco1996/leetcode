#
# @lc app=leetcode.cn id=116 lang=python3
#
# [116] 填充每个节点的下一个右侧节点指针
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""


class Solution:
    def helper(self, node1, node2):
        if node1 is None or node2 is None:
            return
        node1.next = node2

        self.helper(node1.left, node1.right)
        self.helper(node2.left, node2.right)
        self.helper(node1.right, node2.left)

    def connect(self, root):
        if not root:
            return
        self.helper(root.left, root.right)
        return root


# @lc code=end

