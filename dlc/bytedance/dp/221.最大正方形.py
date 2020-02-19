# /*
#  * @lc app=leetcode.cn id=221 lang=golang
#  *
#  * [221] 最大正方形
#  *
#  * https://leetcode-cn.com/problems/maximal-square/description/
#  *
#  * algorithms
#  * Medium (38.57%)
#  * Likes:    229
#  * Dislikes: 0
#  * Total Accepted:    21.9K
#  * Total Submissions: 56.2K
#  * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
#  *
#  * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
#  *
#  * 示例:
#  *
#  * 输入:
#  *
#  * 1 0 1 0 0
#  * 1 0 1 1 1
#  * 1 1 1 1 1
#  * 1 0 0 1 0
#  *
#  * 输出: 4
#  *
#  */

class Solution:
    def maximalSquare(self, matrix):
		#  dp 存可以构成正方形的最大边长
		# dp(i,j) = min( dp(i-1,j-1) , dp(i,j-1),dp(i-1,j)) + 1
        if not matrix:
            return 0
        res, r, c = 0, len(matrix), len(matrix[0])
        dp = [[0]*c for _ in range(r)]

        for i in range(r):
            for j in range(c):
                if not i or not j or matrix[i][j] == '0':
                    dp[i][j] = int(matrix[i][j])
                else:
                    dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
                res = max(res, dp[i][j])
        return res**2

# // @lc code=end

