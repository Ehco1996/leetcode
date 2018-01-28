'''
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.


author: Ehco1996
time: 2017-10-02
'''


class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int

        从右往左遵循左加右减原则
        """
        num_dict = {'I': 1, 'V': 5, 'X': 10,
                    'L': 50, 'C': 100, 'D': 500, 'M': 1000}
        #  反转字符串
        s = s[::-1]
        res = 0
        cur = s[0]
        for num in s:
            if num_dict[num] < num_dict[cur]:
                res -= num_dict[num]
            else:
                res += num_dict[num]
            cur = num
        return res


Solution().romanToInt('MDCLXVI')
