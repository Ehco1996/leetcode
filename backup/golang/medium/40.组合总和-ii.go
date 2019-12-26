package medium

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (56.16%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    22.9K
 * Total Submissions: 40.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// dfs 加剪枝 注意去重
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
			if i != startIdx && candidates[i] == candidates[i-1] {
				// 这个数和上个数一样的情况 剪枝
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i+1, path, sum, target)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, path, 0, target)
	return res
}
