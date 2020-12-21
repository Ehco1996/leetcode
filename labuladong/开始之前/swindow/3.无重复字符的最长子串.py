#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        from collections import defaultdict

        left, right = 0, 0
        res = 0
        rep = []
        window = defaultdict(int)

        while right < len(s):
            c = s[right]
            right += 1
            window[c] += 1
            if window[c] == 1:
                if len(rep) == 0:
                    res = max(res, len(window))
            else:
                rep.append(c)

            while left < right:
                if len(rep) == 0:
                    break
                d = s[left]
                window[d] -= 1
                if window[d] == 0:
                    del window[d]
                left += 1
                if d in rep:
                    rep.remove(d)
        return res

    def test(self, s="abcabcbb"):
        return self.lengthOfLongestSubstring(s)


# @lc code=end

