#
# @lc app=leetcode.cn id=438 lang=python3
#
# [438] 找到字符串中所有字母异位词
#

# @lc code=start
class Solution:
    def findAnagrams(self, s: str, p: str):
        from collections import defaultdict

        need = defaultdict(int)
        for c in p:
            need[c] += 1
        window = defaultdict(int)
        res = []

        left, right = 0, 0
        while right < len(s):
            c = s[right]
            right += 1
            if c in need:
                window[c] += 1
                if window == need:
                    res.append(left)

            while left < right and (right - left) > len(p) - 1:
                d = s[left]
                left += 1
                if d in window:
                    window[d] -= 1
        return res

    def test(self, s="cbaebabacd", p="abc"):
        return self.findAnagrams(s, p)


# @lc code=end

