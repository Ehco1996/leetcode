#
# @lc app=leetcode.cn id=239 lang=python3
#
# [239] 滑动窗口最大值
#

# @lc code=start
class MaxQueue:
    """递减队列。每次push元素都将比自己小的元素删除，
        这样队列里的元素一定是从大到小排列的"""

    def __init__(self):
        self.nums = []

    def pop(self, num):
        if self.nums and self.nums[0] == num:
            return self.nums.pop(0)

    def push(self, num):
        while self.nums and self.nums[-1] < num:
            self.nums.pop(-1)
        self.nums.append(num)

    def max(self):
        return self.nums[0]


class Solution:
    def maxSlidingWindow(self, nums, k: int):
        """单调队列"""
        res = []
        window = MaxQueue()

        for i in range(len(nums)):
            if i < k - 1:
                # 先用前k个数把window填满
                window.push(nums[i])
            else:
                window.push(nums[i])
                res.append(window.max())
                window.pop(nums[i - k + 1])
        return res


# @lc code=end

