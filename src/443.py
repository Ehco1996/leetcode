'''
LeetCode:  String Compression


description:
Given an array of characters, compress it in-place.
The length after compression must always be smaller than or equal to the original array.
Every element of the array should be a character (not int) of length 1.
After you are done modifying the input array in-place, return the new length of the array.


Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".

author: Ehco1996
time: 2017-12-7
'''


class Solution:
    def compress(self, chars):
        """
        :type chars: List[str]
        :rtype: int
        """

        num_count = {}
        res = []
        for s in chars:
            if s not in num_count.keys():
                num_count[s] = 1
            else:
                num_count[s] += 1

        for k, v in num_count.items():
            res.append(str(k))
            if v == 1:
                pass
            else:
                for w in str(v):
                    res.append(w)
        return len(res)

    def compressv2(self, chars):
        anchor = write = 0
        for read, c in enumerate(chars):
            if read + 1 == len(chars) or chars[read + 1] != c:
                chars[write] = chars[anchor]
                write += 1
                if read > anchor:
                    for digit in str(read - anchor + 1):
                        chars[write] = digit
                        write += 1
                anchor = read + 1
        return write


s = Solution()
t = ["a", "a", "b", "b", "c", "c", "c"]

print(s.compress(t))
