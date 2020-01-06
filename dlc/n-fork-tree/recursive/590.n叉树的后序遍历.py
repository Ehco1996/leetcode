"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution:
    def __init__(self):
        self.res = []

    def postorder(self, root):
        if root:
            for node in root.children:
                self.postorder(node)
            self.res.append(root.val)
        return self.res

