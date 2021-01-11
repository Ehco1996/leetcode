#
# @lc app=leetcode.cn id=316 lang=python3
#
# [316] 去除重复字母
#

# @lc code=start
class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        """
        通stack来记录顺序
        用set 来表示该祖母是否进去了
        再用hash看剩下的字符串里是否还有符合要求的（字典顺序）

        比如 s = bcabc 当前stack = [bc],cur= a
        此时应该将stack清空，因为 stack里的bc再接下来的abc里还能遇到
        而且a<b a<c
        """
        from collections import defaultdict

        stack = []
        count = defaultdict(int)
        in_stack = set()

        for c in s:
            count[c] += 1

        for c in s:
            count[c] -= 1

            if c in in_stack:
                continue

            while stack and stack[-1] > c and count[stack[-1]] > 0:
                d = stack.pop(-1)
                in_stack.remove(d)

            stack.append(c)
            in_stack.add(c)
        return "".join(stack)

    def test(self, s="bcabc"):
        return self.removeDuplicateLetters(s)


# @lc code=end

