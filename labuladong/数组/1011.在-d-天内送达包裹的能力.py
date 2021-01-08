#
# @lc app=leetcode.cn id=1011 lang=python3
#
# [1011] 在 D 天内送达包裹的能力
#

# @lc code=start
class Solution:
    def shipWithinDays(self, weights, D: int) -> int:
        """
        可能的最小值是 max(weights)
        最大值是sum(weights)+1
        """
        left = max(weights)
        right = sum(weights) + 1

        while left < right:
            mid = left + (right - left) // 2
            if self.can_ship(mid, weights, D):
                right = mid
            else:
                left = mid + 1
        return left

    def can_ship(self, cap, weights, D):
        t = 0
        now = 0
        for w in weights:
            if w + now <= cap:
                now += w
            else:
                t += 1
                now = w

        if now > 0:
            t += 1
        return t <= D


# @lc code=end

