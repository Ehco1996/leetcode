#
# @lc app=leetcode.cn id=5 lang=python3
#
# [5] 最长回文子串
#

# @lc code=start
class Solution:
    def longestPalindrome(self, s: str) -> str:
        """
        寻找回文串的问题核心思想是：从中间开始向两边扩散来判断回文串
        for 0 <= i < len(s):
            # 找到以 s[i] 为中心的回文串
            palindrome(s, i, i)
            # 找到以s[i+1] 为中心的回文串
            palindrome(s, i, i + 1)
            更新答案
        """
        res = ""
        for i in range(len(s)):
            s1 = self.palindrome(s, i, i)
            s2 = self.palindrome(s, i, i + 1)
            if len(s1) > len(res):
                res = s1
            if len(s2) > len(res):
                res = s2
        return res

    def palindrome(self, s, l, r):
        while l >= 0 and r < len(s) and s[l] == s[r]:
            l -= 1
            r += 1
        return s[l + 1 : r]


# @lc code=end

