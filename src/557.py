'''
LeetCode: Reverse Words in a String III


description:
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

author: Ehco1996
time: 2017-11-27
'''


class Solution:
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        words = s.split(' ')
        newwords = ''

        for word in words:
            for w in word[::-1]:
                newwords += w
            newwords += ' '
        return newwords[:-1]

print(Solution().reverseWords("Let's take LeetCode contest"))