#
# @lc app=leetcode.cn id=322 lang=python3
#
# [322] 零钱兑换
#

# @lc code=start
class Solution:
    def coinChange(self, coins, amount: int) -> int:
        """这个题目是符合最优子结构的
        比如你想求 amount = 11 时的最少硬币数（原问题），
        如果你知道凑出 amount = 10 的最少硬币数（子问题），
        你只需要把子问题的答案加一（再选一枚面值为 1 的硬币）就是原问题的答案。
        因为硬币的数量是没有限制的，所以子问题之间没有相互制，是互相独立的。
        """

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
            res = res if res != inf else -1
            return res

        return dp(amount)


# @lc code=end

