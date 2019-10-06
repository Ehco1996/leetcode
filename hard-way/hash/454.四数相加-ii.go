/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 *
 * https://leetcode-cn.com/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (53.58%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    10.1K
 * Total Submissions: 18.7K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] +
 * D[l] = 0。
 *
 * 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1
 * 之间，最终结果不会超过 2^31 - 1 。
 *
 * 例如:
 *
 *
 * 输入:
 * A = [ 1, 2]
 * B = [-2,-1]
 * C = [-1, 2]
 * D = [ 0, 2]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个元组如下:
 * 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 *
 *
 */

// @lc code=start
import "sort"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	//遍历 A 和 B 所有元素和的组合情况，并记录在 ab_map 中，ab_map 的 key 为两数和，value 为该两数和出现的次数
	// 遍历 C 和 D 所有元素和的组合情况，取和的负值判断其是否在 ab_map 中，若存在则取出 ab_map
	// 对应的 value 值，count = count + value

	cnt := 0
	h := make(map[int]int)

	for _, a := range A {
		for _, b := range B {
			h[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			if v, ok := h[(c+d)*-1]; ok {
				cnt += v
			}
		}
	}
	return cnt
}

// @lc code=end

