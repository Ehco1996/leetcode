#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#

# @lc code=start
from copy import copy


class Solution:
    def permute(self, nums):
        res = []

        def backtrack(now, last):
            if len(last) == 0:
                res.append(now)
                return

            for num in last:
                old = copy(now)  # 备份一下当前的选择
                now.append(num)  # 做选择
                last = last[1:]
                backtrack(now, last)
                now = old  # 回退选择
                last.append(num)  # 回退选择

        backtrack([], nums)
        return res

    def test(self, nums=[1, 2, 3]):
        return self.permute(nums)


# @lc code=end
