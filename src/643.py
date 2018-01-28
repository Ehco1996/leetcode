'''
LeetCode: Reverse Integer

description: Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value. And you need to output the maximum average value.



author: Ehco1996
time: 2017-08-11
'''


class Solution(object):
    def findMaxAverage(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: float
        """

        sums = sum(nums[:k])

        res = sums * 1.0
        for i in range(k, len(nums)):
            sums += nums[i] - nums[i - k] * 1.0
            res = max(res, sums)

        return res / k


l = [1, 12, -5, -6, 50, 3, -30, 25]
k = 4
print(Solution().findMaxAverage(l, k))
