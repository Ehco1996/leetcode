'''
LeetCode: Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.


author: Ehco1996
time: 2018-01-29
'''


class Solution:

    @staticmethod
    def isRepeat(s):
        cahe = []
        for i in s:
            if i not in cahe:
                cahe.append(i)
            else:
                return True
        return False

    def findSubstring(self, s):
        num = len(s)
        subs = []
        for x in range(num):
            for i in range(num - x):
                subs.append(s[i:i + x + 1])
        return subs

    def findNoRepeatSubstring(self, s):
        subs = self.findSubstring(s)
        no_repeat_subs = []
        for sub in subs:
            if self.isRepeat(sub) == False:
                no_repeat_subs.append(sub)
        return no_repeat_subs

    def lengthOfLongestSubstringByHard(self, s):
        if len(s) == 0:
            return 0
        no_repeat_subs = self.findNoRepeatSubstring(s)
        no_repeat_subs = sorted(no_repeat_subs, key=len, reverse=True)
        return(len(no_repeat_subs[0]))

    def lengthOfLongestSubstring(self, s):
        used = {}
        max_length = start = 0
        for i, c in enumerate(s):
            if c in used and start <= used[c]:
                start = used[c] + 1
            else:
                max_length = max(max_length, i - start + 1)
            used[c] = i
        return max_length


s = Solution()
print(s.findSubstring('asdf'))
print(s.lengthOfLongestSubstring(
    "asdf"))
