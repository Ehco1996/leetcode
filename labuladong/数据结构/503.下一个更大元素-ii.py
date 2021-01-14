#
# @lc app=leetcode.cn id=503 lang=python3
#
# [503] 下一个更大元素 II
#

# @lc code=start
class Solution:
    def nextGreaterElements(self, nums):
        """循环数组的技巧在于构造一个两倍长度的数组"""
        stack = []
        double = nums + nums
        res = [-1] * len(double)
        for i in range(len(double)-1, -1, -1):
            while stack and stack[-1] <= double[i]:
                stack.pop(-1)
            if stack:
                res[i] = stack[-1]
            else:
                res[i] = -1
            stack.append(double[i])
        return res[: len(nums)]


# @lc code=end

