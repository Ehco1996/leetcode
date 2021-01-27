#
# @lc app=leetcode.cn id=53 lang=python3
#
# [53] 最大子序和
#

# @lc code=start
class Solution:
    def maxSubArray(self, nums) -> int:
        """
        以nums[i]为结尾的「最大子数组和」为dp[i]
        dp(i) = max(dp[i-1]+now,now)
        """
        if not nums:
            return 0
        s = nums[0]
        cur = nums[0]
        for num in nums[1:]:
            cur = max(cur + num, num)
            s = max(s, cur)
        return s


# @lc code=end

