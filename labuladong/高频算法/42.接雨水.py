#
# @lc app=leetcode.cn id=42 lang=python3
#
# [42] 接雨水
#

# @lc code=start
class Solution:
    def trap(self, height) -> int:
        """
        每个格子能装的水量是min(max_left,max_right)-height[i]
        """
        if not height:
            return 0
        n = len(height)
        left = 0
        right = n - 1

        l_max = height[left]
        r_max = height[right]
        res = 0
        while left <= right:
            l_max = max(l_max, height[left])
            r_max = max(r_max, height[right])
            if l_max < r_max:
                res += l_max - height[left]
                left += 1
            else:
                res += r_max - height[right]
                right -= 1
        return res


# @lc code=end

