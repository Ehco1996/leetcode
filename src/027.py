'''
LeetCode: Remove Element
description:
Given an array and a value, remove all instances of that value in place and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example:
Given input array nums = [3,2,2,3], val = 3

Your function should return length = 2, with the first two elements of nums being 2.

author: Ehco1996
time: 2017-10-05
'''


class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        for i in range(len(nums)):
            if val in nums:
                nums.remove(val)
        return len(nums)


print(Solution().removeElement([3, 3, 2, 2], 3))
