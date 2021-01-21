#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for c in s:
            if c in ["(", "{", "["]:
                stack.append(c)
            else:
                if stack and self.leftof(c) == stack[-1]:
                    stack.pop(-1)
                else:
                    return False
        return len(stack) == 0

    def leftof(self, c):
        if c == ")":
            return "("
        if c == "}":
            return "{"
        if c == "]":
            return "["


# @lc code=end

