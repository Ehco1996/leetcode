"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
class Solution():

    def __init__(self):
        self.res = []

    def preorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        if root:
            self.res.append(root.val)
            for node in root.children:
                self.preorder(node)
        return self.res

