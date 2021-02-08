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
    def connect(self, root: "Node") -> "Node":
        if not root:
            return root
        self.helper(root.left, root.right)
        return root

    def helper(self, node1, node2):
        if not node1 or not node2:
            return
        node1.next = node2

        self.connect(node1)  # 连左树
        self.connect(node2)  # 连右树
        self.helper(node1.right, node2.left)  # 跨树连


# @lc code=end

