// package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 *
 * https://leetcode-cn.com/problems/string-compression/description/
 *
 * algorithms
 * Easy (36.07%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    6.2K
 * Total Submissions: 17.1K
 * Testcase Example:  '["a","a","b","b","c","c","c"]'
 *
 * 给定一组字符，使用原地算法将其压缩。
 *
 * 压缩后的长度必须始终小于或等于原数组长度。
 *
 * 数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。
 *
 * 在完成原地修改输入数组后，返回数组的新长度。
 *
 *
 *
 * 进阶：
 * 你能否仅使用O(1) 空间解决问题？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["a","a","b","b","c","c","c"]
 *
 * 输出：
 * 返回6，输入数组的前6个字符应该是：["a","2","b","2","c","3"]
 *
 * 说明：
 * "aa"被"a2"替代。"bb"被"b2"替代。"ccc"被"c3"替代。
 *
 *
 * 示例 2：
 *
 *
 * 输入：
 * ["a"]
 *
 * 输出：
 * 返回1，输入数组的前1个字符应该是：["a"]
 *
 * 说明：
 * 没有任何字符串被替代。
 *
 *
 * 示例 3：
 *
 *
 * 输入：
 * ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
 *
 * 输出：
 * 返回4，输入数组的前4个字符应该是：["a","b","1","2"]。
 *
 * 说明：
 * 由于字符"a"不重复，所以不会被压缩。"bbbbbbbbbbbb"被“b12”替代。
 * 注意每个数字在数组中都有它自己的位置。
 *
 *
 * 注意：
 *
 *
 * 所有字符都有一个ASCII值在[35, 126]区间内。
 * 1 <= len(chars) <= 1000。
 *
 *
 */

func compress(chars []byte) int {
	// 双指针+滑动窗口
	size := 0
	lens := len(chars)
	// 由于最后一个字符也需要判断，所以将右指针终点放到数组之外一格
	for l, r := 0, 0; r <= lens; r++ {
		// 当遍历完成，或右指针元素不等于左指针元素时，更新数组
		if r == lens || chars[r] != chars[l] {
			// 更新字符
			chars[size] = chars[l]
			size++
			// 更新计数，当个数大于 1 时才更新
			if r-l > 1 {
				for _, b := range []byte(strconv.Itoa(r - l)) {
					chars[size] = b
					size++
				}
			}
			// 挪动左指针
			l = r
		}
	}
	return size
}
