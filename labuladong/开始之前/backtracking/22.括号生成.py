#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
#

# @lc code=start
from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        """
        从（）开始生成，合法的括号组合就是在括号里不停插入新的（）
        """
        res = set()

        def dfs(path):
            if len(path) == 2 * n:
                res.add(path)
                return
            for i in range(len(path)):
                dfs(path[:i] + "()" + path[i:])

        dfs("()")
        return list(res)

    def test(self):
        return self.generateParenthesis(3)


# @lc code=end

