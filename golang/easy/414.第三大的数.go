package easy

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (32.22%)
 * Likes:    68
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 34.5K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 *
 * 示例 1:
 *
 *
 * 输入: [3, 2, 1]
 *
 * 输出: 1
 *
 * 解释: 第三大的数是 1.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1, 2]
 *
 * 输出: 2
 *
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
 *
 * 示例 3:
 *
 *
 * 输入: [2, 2, 3, 1]
 *
 * 输出: 1
 *
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 *
 *
 */
import "math"

func thirdMax(nums []int) int {
	// 起是可以用最小堆（k=3）
	one, two, three := math.MinInt32, math.MinInt32, math.MinInt32
	for _, num := range nums {
		if num == one || num == two || num == three {
			continue
		}
		if num > one {
			one, two, three = num, one, two
		} else if num > two {
			two, three = num, two
		} else if num > three {
			three = num
		}
		fmt.Println(one, two, three)
	}

	if three != math.MinInt32 {
		return three
	}
	return one
}
