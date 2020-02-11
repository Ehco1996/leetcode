/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (47.09%)
 * Likes:    840
 * Dislikes: 0
 * Total Accepted:    51.7K
 * Total Submissions: 106.8K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */

// @lc code=start

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
func trap(height []int) int {
	//https://www.bilibili.com/video/av18241289/
	res := 0
	lMax := 0
	rMax := 0
	for l, r := 0, len(height)-1; l < r; {
		lMax = max(height[l], lMax)
		rMax = max(height[r], rMax)
		if lMax < rMax {
			res += (lMax - height[l])
			l++
		} else {
			res += (rMax - height[r])
			r--
		}
	}
	return res
}

// @lc code=end

