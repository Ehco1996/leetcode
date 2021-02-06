#
# @lc app=leetcode.cn id=77 lang=python3
#
# [77] 组合
#

# @lc code=start
from copy import copy
from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []
        now = []

        def backtrack(now, start):
            if k == len(now):
                res.append(now)
                return
            for i in range(start, n + 1):
                now.append(i)
                backtrack(copy(now), i + 1)
                now.pop()

        backtrack(copy(now), 1)
        return res

    def test(self):
        return self.combine(n=4, k=2)


# @lc code=end

