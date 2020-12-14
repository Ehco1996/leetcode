#
# @lc app=leetcode.cn id=322 lang=python3
#
# [322] 零钱兑换
#

# @lc code=start
class Solution:
    def coinChange_with_memory(self, coins, amount: int) -> int:
        """这个题目是符合最优子结构的
        比如你想求 amount = 11 时的最少硬币数（原问题），
        如果你知道凑出 amount = 10 的最少硬币数（子问题），
        你只需要把子问题的答案加一（再选一枚面值为 1 的硬币）就是原问题的答案。
        因为硬币的数量是没有限制的，所以子问题之间没有相互制，是互相独立的。
        """
        memo = dict()
        # 定义：要凑出金额 n，至少要 dp(n) 个硬币
        def dp(n):
            # init base
            if n == 0:
                return 0
            if n < 0:
                return -1
            inf = float("INF")
            res = inf
            for coin in coins:
                sub = dp(n - coin)
                if sub == -1:  # 子问题无解 跳过
                    continue
                res = min(res, 1 + sub)
            memo[n] = res if res != inf else -1
            return memo[n]

        return dp(amount)

    def coinChange(self, coins, amount: int) -> int:
        # 数组大小为 amount + 1，初始值也为 amount + 1
        # 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出
        dp = [amount + 1] * (amount + 1)
        # init base case
        dp[0] = 0
        # 定义：要凑出金额 n，至少要 dp(n) 个硬币
        for i in range(0, amount + 1):
            for coin in coins:
                if (i - coin) < 0:
                    # 无解
                    continue
                dp[i] = min(dp[i], 1 + dp[i - coin])
        return dp[amount] if dp[amount] != amount + 1 else -1

    def test(self, coins=[1, 2, 5], amount=11):
        return self.coinChange(coins, amount)


# @lc code=end

