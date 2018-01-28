'''
LeetCode: Add Two Numbers

description:
Given an array of 2n integers, your task is to group these integers into n pairs of integer, 
say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.


Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).


author: Ehco1996
time: 2017-11-17
'''

class Solution:
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l = sorted(nums)
        s = 0 
        for i in range(0,len(l),2):
            s += l[i]
        return s
    



print (Solution().arrayPairSum([1,3,2,4,7,8]))