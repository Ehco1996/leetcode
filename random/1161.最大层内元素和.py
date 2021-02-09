#
# @lc app=leetcode.cn id=1161 lang=python3
#
# [1161] 最大层内元素和
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


from collections import deque


class Solution:
    def maxLevelSum(self, root: TreeNode) -> int:
        curr_level = max_level = 1
        max_sum = float("-inf")
        curr_sum = 0

        marker = None
        queue = deque([root, marker])

        while queue:
            x = queue.popleft()
            # continue current level
            if x != marker:
                curr_sum += x.val
                if x.left:
                    queue.append(x.left)
                if x.right:
                    queue.append(x.right)
            # end of current level, go to the next level
            else:
                if curr_sum > max_sum:
                    max_sum, max_level = curr_sum, curr_level
                curr_sum = 0
                curr_level += 1
                if len(queue) == 0:
                    break
                else:
                    queue.append(marker)

        return max_level

    def test(self):
        root = TreeNode(1)
        left = TreeNode(2)
        right = TreeNode(3)
        root.left = left
        root.right = right
        return self.maxLevelSum(root)


# @lc code=end
