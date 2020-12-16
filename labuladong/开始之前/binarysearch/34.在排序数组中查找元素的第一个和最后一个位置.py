#
# @lc app=leetcode.cn id=34 lang=python3
#
# [34] 在排序数组中查找元素的第一个和最后一个位置
#

# @lc code=start
class Solution:
    def searchRange(self, nums, target):
        if not nums:
            return [-1, -1]
        target_idx = -1
        left_bound_idx = -1
        right_bound_idx = -1

        # 1. 先通过二分找到目标值的位置

        left = 0
        right = len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                target_idx = mid
                break

        # 2. 向左找左边界
        left = 0
        right = target_idx
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                # 找到之后不返回，锁定在左侧边界
                right = mid - 1
        # 判断一下越界的情况
        if left >= len(nums) or nums[left] != target:
            left_bound_idx = -1
        else:
            left_bound_idx = left

        # 3. 向右找右边界
        left = target_idx
        right = len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                # 找到之后不返回，锁定在右侧边界
                left = mid + 1
        # 判断一下越界的情况
        if right < 0 or nums[right] != target:
            right_bound_idx = -1
        else:
            right_bound_idx = right

        return [left_bound_idx, right_bound_idx]

    def test(self, nums=[5, 7, 7, 8, 8, 10], target=8):
        return self.searchRange(nums, target)


# @lc code=end

