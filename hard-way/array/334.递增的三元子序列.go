/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 *
 * https://leetcode-cn.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (34.86%)
 * Likes:    76
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 26.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
 *
 * 数学表达式如下:
 *
 * 如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
 * 使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
 *
 * 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5]
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: [5,4,3,2,1]
 * 输出: false
 *
 */

// @lc code=start
import "math"

func increasingTriplet(nums []int) bool {
	// 有3个槽位，a,b,c
	// 满足条件 a < b < c，即可
	// 需要将合适的元素填入这3个槽位
	a, b := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

