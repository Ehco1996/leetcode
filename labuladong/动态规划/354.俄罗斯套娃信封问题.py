#
# @lc app=leetcode.cn id=354 lang=python3
#
# [354] 俄罗斯套娃信封问题
#

# @lc code=start
class Solution:
    def maxEnvelopes(self, envelopes):
        """
        1. 先按宽度排序，如果宽度一样就按照高度降序排序
        2. 把所有的高领出来昨晚一个数组
        3. 从这个数组里找最长递增子序列的长度
        """
        envelopes = sorted(envelopes, key=lambda x: (x[0], -x[1]))
        h_list = [h[1] for h in envelopes]
        print(envelopes, h_list)
        return self.length_of_LTS(h_list)

    def length_of_LTS(self, nums):
        """最长递增子序列
        dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度。
        """
        if not nums:
            return 0
        dp = [1] * len(nums)
        for i in range(len(nums)):
            for j in range(i):
                if nums[j] < nums[i]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return max(dp)

    def test(self):
        return self.maxEnvelopes([[4, 5], [4, 6], [6, 7], [2, 3], [1, 1]])


# @lc code=end

