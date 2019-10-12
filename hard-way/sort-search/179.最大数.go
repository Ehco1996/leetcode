/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (32.94%)
 * Likes:    154
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 36.9K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
 *
 * 示例 1:
 *
 * 输入: [10,2]
 * 输出: 210
 *
 * 示例 2:
 *
 * 输入: [3,30,34,5,9]
 * 输出: 9534330
 *
 * 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 */

// @lc code=start
import "sort"
import "strconv"

type StrNum struct {
	Nums []string
}

func (a StrNum) Len() int      { return len(a.Nums) }
func (a StrNum) Swap(i, j int) { a.Nums[i], a.Nums[j] = a.Nums[j], a.Nums[i] }
func (a StrNum) Less(i, j int) bool {
	si := a.Nums[i]
	sj := a.Nums[j]
	return si+sj < sj+si
}

func largestNumber(nums []int) string {
	// 自定义排序顺序  string(a) + string(b) > string(b) + string(a) "12">"21"?
	sNums := make([]string, len(nums))
	for i, n := range nums {
		sNums[i] = strconv.Itoa(n)
	}
	s := StrNum{sNums}
	sort.Sort(s)

	res := ""
	for i := len(s.Nums) - 1; i >= 0; i-- {
		res += s.Nums[i]
	}
	if string(res[0]) == "0" {
		res = "0"
	}
	return res
}

// @lc code=end

