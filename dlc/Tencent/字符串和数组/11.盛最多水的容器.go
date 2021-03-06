/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (56.85%)
 * Likes:    787
 * Dislikes: 0
 * Total Accepted:    78.9K
 * Total Submissions: 136.9K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最x多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(leftIdx, leftVal, rightIdx, rightVal int) int {
	return min(leftVal, rightVal) * (rightIdx - leftIdx)
}

func maxArea(height []int) int {
	// 滑动窗口 短边向长边方向移动
	res := 0
	for leftIdx, rightIdx := 0, len(height)-1; leftIdx < rightIdx; {
		now := calc(leftIdx, height[leftIdx], rightIdx, height[rightIdx])
		res = max(now, res)
		if height[leftIdx] < height[rightIdx] {
			leftIdx++
		} else {
			rightIdx--
		}
	}
	return res
}
