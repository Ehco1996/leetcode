#
# @lc app=leetcode.cn id=849 lang=python3
#
# [849] 到最近的人的最大距离
#

# @lc code=start
from typing import List


class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        """
        找两点之间的最大距离
        距离就用双指针
        """
        res = 0
        p1 = None
        p2 = None
        for p2 in range(len(seats)):
            if seats[p2] == 1:
                if p1 is None:  # seats[0] 为0
                    res = max(res, p2)
                else:
                    res = max(res, int((p2 - p1) / 2))
                p1 = p2
        if seats[p2] == 0:
            res = max(res, p2-p1)
        return res


# @lc code=end

