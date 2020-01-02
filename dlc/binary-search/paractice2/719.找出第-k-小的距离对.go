/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 k 小的距离对
 *
 * https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance/description/
 *
 * algorithms
 * Hard (29.40%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    2K
 * Total Submissions: 6.6K
 * Testcase Example:  '[1,3,1]\n1'
 *
 * 给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。
 *
 * 示例 1:
 *
 *
 * 输入：
 * nums = [1,3,1]
 * k = 1
 * 输出：0
 * 解释：
 * 所有数对如下：
 * (1,3) -> 2
 * (1,1) -> 0
 * (3,1) -> 2
 * 因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。
 *
 *
 * 提示:
 *
 *
 * 2 <= len(nums) <= 10000.
 * 0 <= nums[i] < 1000000.
 * 1 <= k <= len(nums) * (len(nums) - 1) / 2.
 *
 *
 */

// @lc code=start
import "sort"

func smallestDistancePair(nums []int, k int) int {
	// 还是二分 排序后统计 差值小于mid(距离的中位数)的个数
	sort.Ints(nums)

	low := 0
	high := nums[len(nums)-1] - nums[0]

	for low < high {
		mid := (low + high) / 2
		cnt := 0
		left := 0
		for right := 0; right < len(nums); right++ {
			for ; nums[right]-nums[left] > mid; left++ {
				//以它为右端的满足距离小于等于 mid 的距离对数目即为 right - left
			}
			cnt += right - left
		}
		if cnt >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// @lc code=end

