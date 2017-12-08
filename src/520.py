'''
LeetCode:  String Compression


description:
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.

Example 1:
Input: "USA"
Output: True
Example 2:
Input: "FlaG"
Output: False

author: Ehco1996
time: 2017-12-9
'''


class Solution:
    def detectCapitalUse(self, word):
        """
        :type word: str
        :rtype: bool
        """
        if len(word) == 1:
            return True
        Up_words = [chr(i) for i in range(65, 91)]
        Low_words = [chr(i) for i in range(97, 123)]
        if word[0] in Up_words:
            if word[1] in Up_words:
                for w in word[2:]:
                    if w not in Up_words:
                        return False
            else:
                for w in word[2:]:
                    if w not in Low_words:
                        return False
        else:
            for w in word[1:]:
                if w not in Low_words:
                    return False
        return True


s = Solution()

print(s.detectCapitalUse('G'))
