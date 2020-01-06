"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution:
    def __init__(self):
        self.dep = 0

    def maxDepth(self, root):
        if not root:
            return self.dep
        self.helper(root, 1)
        return self.dep

    def helper(self, node, now):
        if not node:
            return
        if now > self.dep:
            self.dep = now
        for n in node.children:
            self.helper(n, now + 1)

