#
# @lc app=leetcode.cn id=710 lang=python3
#
# [710] 黑名单中的随机数
#

# @lc code=start
class Solution:
    """
    用hash把黑名单中的数都转成一个合法的数
    """

    def __init__(self, N: int, blacklist):
        m = dict()
        last = N - 1
        s = set(blacklist)
        for b in blacklist:
            while last in s:
                last -= 1
            m[b] = last
        # 所有black里的数都已经合法了
        self.m = m
        self.N = N

    def pick(self) -> int:
        import random

        idx = random.randint(0, self.N - 1)
        if idx in self.m:
            return self.m[idx]
        return idx


# Your Solution object will be instantiated and called as such:
# obj = Solution(N, blacklist)
# param_1 = obj.pick()
# @lc code=end

