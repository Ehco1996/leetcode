'''
LeetCode: Valid Parentheses


description:
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

author: Ehco1996
time: 2017-11-1
'''


class Solution:
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        # 构造一个堆栈结构
        stack = []
        cmap = {
            '{': '}',
            '[': ']',
            '(': ')',
        }
        # 判断第一位是否是左括号，或者长度是否是双数
        if s[0] not in cmap.keys() or len(s) % 2 != 0:
            return False

        for c in s:
            # 当字符串是keys时
            if c in cmap.keys():
                stack.append(c)
            elif c in cmap.values():
                try:
                    # 取出左边
                    res = stack.pop()
                    # 判断左边和右边是否对应
                    if c != cmap[res]:
                        return False
                except IndexError:
                    return False

        # 判断堆栈是否为空
        if len(stack) != 0:
            return False
        return True


print(Solution().isValid('()'))
