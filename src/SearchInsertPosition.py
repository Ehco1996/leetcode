'''
LeetCode: Search Insert Position

description:
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Input: [1,3,5,6], 5
Output: 2
Input: [1,3,5,6], 2
Output: 1
Input: [1,3,5,6], 7
Output: 4
Input: [1,3,5,6], 0
Output: 0

author: Ehco1996
time: 2017-11-04
'''


class Solution:
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """

        if target in nums:
            return nums.index(target)
        if target < nums[0]:
            return 0
        else:
            for i in nums[::-1]:
                if target > i:
                    return nums.index(i) + 1


nums = [1, 3, 5, 6]
target = 7
print(Solution().searchInsert(nums, target))
