/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (62.55%)
 * Likes:    307
 * Dislikes: 0
 * Total Accepted:    39.7K
 * Total Submissions: 62.3K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */

// @lc code=start
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	// 把三角形倒过来看，把每一层相邻的两个元素中小的拿个加到下一层去，
	// 最终在终点[0][0]位置的就是最小值

	hight := len(triangle)
	if hight == 0 {
		return 0
	}

	for i := hight - 1; i > 0; i-- {
		row := triangle[i]
		// 计算row，并往row-1行加
		for idx, _ := range triangle[i-1] {
			triangle[i-1][idx] += min(row[idx], row[idx+1])
		}
	}
	return triangle[0][0]

}

// @lc code=end

