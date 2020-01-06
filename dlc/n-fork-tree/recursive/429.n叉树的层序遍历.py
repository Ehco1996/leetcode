"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution(object):
    def __init__(self):
        self.res = []

    def levelOrder(self, root):
        """
        :type root: Node
        :rtype: List[List[int]]
        """
        if not root:
            return self.res
        queue = [(root, 0)]
        while len(queue) > 0:
            node, dep = queue[0]
            queue = queue[1:]
            if len(self.res) <= dep:
                self.res.append([])
            self.res[dep].append(node.val)
            for n in node.children:
                queue.append((n, dep + 1))
        return self.res

