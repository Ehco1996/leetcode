#
# @lc app=leetcode.cn id=76 lang=python3
#
# [76] 最小覆盖子串
#

# @lc code=start
from collections import defaultdict


class Solution:
    def minWindow(self, s: str, t: str) -> str:

        window = defaultdict(int)
        need = defaultdict(int)
        for c in t:
            need[c] += 1

        valid = 0  # 表示窗口中满足need条件的字符串的字符个数
        left, right = 0, 0

        start = 0  # 记录找到的字符串的开始下标
        length = len(s) + 1

        while right < len(s):
            c = s[right]
            # 移动窗口
            right += 1
            if c in need:
                # 更新窗口里的数
                window[c] += 1
                if need[c] == window[c]:
                    valid += 1

            # 判断左侧是不是要收缩了（当前是否已经满足条件了）
            while valid == len(need):
                # 更新最小子串
                if (right - left) < length:
                    start = left
                    length = right - left

                d = s[left]
                left += 1
                # 将最左侧的数据从window中移除
                if d in need:
                    if need[d] == window[d]:
                        valid -= 1
                    window[d] -= 1
        if length != len(s) + 1:
            return s[start : start + length]
        return ""

    def test(self, s="ADOBECODEBANC", t="ABC"):
        return self.minWindow(s, t)


# @lc code=end

