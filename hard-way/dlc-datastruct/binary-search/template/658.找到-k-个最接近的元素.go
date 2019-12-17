/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 *
 * https://leetcode-cn.com/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (41.74%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    4.9K
 * Total Submissions: 11.7K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * 给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。如果有两个数与 x
 * 的差值一样，优先选择数值较小的那个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,3,4,5], k=4, x=3
 * 输出: [1,2,3,4]
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,3,4,5], k=4, x=-1
 * 输出: [1,2,3,4]
 *
 *
 *
 *
 * 说明:
 *
 *
 * k 的值为正数，且总是小于给定排序数组的长度。
 * 数组不为空，且长度不超过 10^4
 * 数组里的每个元素与 x 的绝对值不超过 10^4
 *
 *
 *
 *
 * 更新(2017/9/19):
 * 这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。
 *
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	// 根据区间距离来二分
	// 比较区间两端的数值和x的距离，如果左边离得远了，就向右边走；如果右边离得远了，就像左边走
	l := 0
	r := len(arr) - k

	var mid int
	for l < r {
		mid = (l + r) / 2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[l : l+k]

}

// @lc code=end

