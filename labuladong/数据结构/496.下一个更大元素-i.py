#
# @lc app=leetcode.cn id=496 lang=python3
#
# [496] 下一个更大元素 I
#

# @lc code=start
class Solution:
    def nextGreaterElement(self, nums1, nums2):
        """单调栈呀，倒着入"""
        # 存放nums2中比每个元素的下一个更大元素
        h = dict()
        stack = []
        for num in nums2[::-1]:
            while stack and stack[-1] <= num:
                # 如果栈里的元素比当前的元素还小，就pop出去
                stack.pop(-1)
            if stack:
                h[num] = stack[-1]
            else:
                h[num] = -1
            stack.append(num)
        return [h[num] for num in nums1]


# @lc code=end

