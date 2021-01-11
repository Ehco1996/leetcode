#
# @lc app=leetcode.cn id=283 lang=python3
#
# [283] 移动零
#

# @lc code=start
class Solution:
    def moveZeroes(self, nums) -> None:
        """
        Do not return anything, modify nums in-place instead.
        找到所有不是零的数安顺序排好，末尾补零就可以
        """
        idx = 0
        for i in range(0, len(nums)):
            if nums[i] != 0:
                nums[idx] = nums[i]
                idx += 1
        for i in range(idx, len(nums)):
            nums[i] = 0


# @lc code=end

