'''
LeetCode: Keyboard Row


description:
Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]

author: Ehco1996
time: 2017-11-29
'''

class Solution:
    def findWords(self, words):
        """
        :type words: List[str]
        :rtype: List[str]
        """
        k_rows=['qwertyuiopQWERTYUIOP','asdfghjklASDFGHJKL','zxcvbnmZXCVBNM']
        res = []
        for row in k_rows:
            for word in words:
                i = 0
                for s in word:
                    if s in row:
                        i+=1
                    else:
                        break
                if i == len(word):
                    res.append(word)
        return res

print (Solution().findWords(["Hello", "Alaska", "Dad", "Peace"]))