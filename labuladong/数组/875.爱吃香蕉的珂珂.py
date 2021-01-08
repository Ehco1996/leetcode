#
# @lc app=leetcode.cn id=875 lang=python3
#
# [875] 爱吃香蕉的珂珂
#

# @lc code=start
class Solution:
    def minEatingSpeed(self, piles, H: int) -> int:
        """
        speed最大可能为多少，最少可能为多少呢？
        显然最少为 1，最大为max(piles)，
        因为一小时最多只能吃一堆香蕉。
        那么暴力解法就很简单了，
        只要从 1 开始穷举到max(piles)，
        一旦发现发现某个值可以在H小时内吃完所有香蕉，这个值就是最小速度

        这时可以用二分法啦查找速度了
        """
        max_s = max(piles)
        left = 1
        right = max_s + 1
        while left < right:
            mid = left + (right - left) // 2
            if self.can_eat_piles(piles, mid, H):
                right = mid
            else:
                left = mid + 1
        return left

    def can_eat_piles(self, piles, s, H):
        import math

        t = 0
        for p in piles:
            t += math.ceil(p / s)
        return t <= H


# @lc code=end

