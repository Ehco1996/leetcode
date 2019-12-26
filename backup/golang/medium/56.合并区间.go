package medium

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (37.61%)
 * Likes:    187
 * Dislikes: 0
 * Total Accepted:    29.1K
 * Total Submissions: 75.7K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

import "sort"

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func merge(intervals [][]int) [][]int {
	// 排序后向右移动
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals[:], func(i, j int) bool {
		for x := range intervals[i] {
			if intervals[i][x] == intervals[j][x] {
				continue
			}
			return intervals[i][x] < intervals[j][x]
		}
		return false
	})

	res := [][]int{}

	for idx := 0; idx < len(intervals); idx++ {
		nums := intervals[idx]
		i := idx + 1
		r := nums[1]
		for ; i < len(intervals) && r >= intervals[i][0]; i++ {
			r = max(intervals[i][1], r)
		}
		res = append(res, []int{nums[0], r})
		idx = i - 1
	}
	return res
}
