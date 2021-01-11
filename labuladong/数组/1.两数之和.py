#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums, target: int):
        h = dict()
        for idx, num in enumerate(nums):
            if h.get(target - num) is not None:
                return [idx, h.get(target - num)]
            h[num] = idx
        return [-1, -1]

    def test(self, nums=[2, 7, 11, 15], target=9):
        return self.twoSum(nums, target)


# @lc code=end

