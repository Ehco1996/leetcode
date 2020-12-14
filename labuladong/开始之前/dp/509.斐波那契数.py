#
# @lc app=leetcode.cn id=509 lang=python3
#
# [509] 斐波那契数
#

# @lc code=start
class Solution:
    def __init__(self) -> None:
        self.m = dict()

    def _fib(self, N: int) -> int:
        """这种解法有大量的重复计算，但是最能表达其推导式原本的意思"""
        if N == 0:
            return 0
        if N == 1 or N == 2:
            return 1
        else:
            return self.fib(N - 1) + self.fib(N - 2)

    def fib_with_memory(self, n):
        """用备忘录去避免重复计算
        这是 自顶向下，从最大往最小分解
        """
        if self.m.get(n) is None:
            self.m[n] = self._fib(n)
        return self.m[n]

    def fib(self, n):
        """动态规划法
        这是 自低向上，从最小开始往最大计算
        """
        if n < 3:
            return self._fib(n)
        # base case
        dp = [0] * (n + 1)
        dp[1] = dp[2] = 1
        # 累积dp表
        for i in range(3, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]
        return dp[n]

    def fib_min_memory(self, n):
        """最优解: 发现不需要整张dp表，只需要前1、2两位就好"""
        if n < 3:
            return self._fib(n)
        # init base cace
        a, b = 1, 1
        for i in range(3, n + 1):
            a, b = b, a + b
        return b

    def test(self):
        for i in range(0, 20):
            assert self._fib(i) == self.fib_min_memory(i)


# @lc code=end
