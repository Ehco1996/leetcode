#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        """
        用栈来存，遇到左边的括号直接入栈，遇到右边的括号就从栈里往外抛，
        如果抛出的元素不是当前的括号的另一半的话说明不合法
        最后还要判断栈是否为空
        """
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

