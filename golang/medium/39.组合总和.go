package medium

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (66.55%)
 * Likes:    379
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 53.7K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// dfs
	res := [][]int{}
	if len(candidates) == 0 {
		return res
	}
	path := []int{}
	sort.Ints(candidates)
	var dfs func(startIdx int, path []int, sum, target int)
	dfs = func(startIdx int, path []int, sum, target int) {
		if sum == target {
			t := make([]int, len(path))
			copy(t, path)
			res = append(res, t)
			return
		} else if sum > target {
			return
		}
		// walkthrough
		for i := startIdx; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i, path, sum, target)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, path, 0, target)
	return res
}
