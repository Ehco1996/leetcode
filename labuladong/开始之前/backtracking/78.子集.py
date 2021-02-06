#
# @lc app=leetcode.cn id=78 lang=python3
#
# [78] 子集
#

# @lc code=start
from typing import List
from copy import copy


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        now = []

        def backtrack(now, start):
            res.append(now)
            for i in range(start, len(nums)):
                now.append(nums[i])
                backtrack(copy(now), i + 1)
                now.pop()

        backtrack(copy(now), 0)
        return res

    def test(self):
        return self.subsets([1, 2, 3])


# @lc code=end

