// package easy

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 *
 * https://leetcode-cn.com/problems/license-key-formatting/description/
 *
 * algorithms
 * Easy (37.65%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 9.7K
 * Testcase Example:  '"5F3Z-2e-9-w"\n4'
 *
 * 给定一个密钥字符串S，只包含字母，数字以及 '-'（破折号）。N 个 '-' 将字符串分成了 N+1 组。给定一个数字
 * K，重新格式化字符串，除了第一个分组以外，每个分组要包含 K 个字符，第一个分组至少要包含 1 个字符。两个分组之间用
 * '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。
 *
 * 给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。
 *
 * 示例 1：
 *
 *
 * 输入：S = "5F3Z-2e-9-w", K = 4
 *
 * 输出："5F3Z-2E9W"
 *
 * 解释：字符串 S 被分成了两个部分，每部分 4 个字符；
 * 注意，两个额外的破折号需要删掉。
 *
 *
 * 示例 2：
 *
 *
 * 输入：S = "2-5g-3-J", K = 2
 *
 * 输出："2-5G-3J"
 *
 * 解释：字符串 S 被分成了 3 个部分，按照前面的规则描述，第一部分的字符可以少于给定的数量，其余部分皆为 2 个字符。
 *
 *
 *
 *
 * 提示:
 *
 *
 * S 的长度不超过 12,000，K 为正整数
 * S 只包含字母数字（a-z，A-Z，0-9）以及破折号'-'
 * S 非空
 *
 *
 *
 *
 */
// func licenseKeyFormatting1(S string, K int) string {
// 	var res string

// 	count := K

// 	for i := 0; i < len(S); i++ {
// 		now := string(S[i])
// 		if now != "-" && count > 0 {
// 			res += now
// 			count--
// 			if count == 0 && i != (len(S)-1) {
// 				res += "-"
// 				count = K
// 			}
// 		}
// 	}

// 	for i := len(res) - 1; i >= 0; i-- {
// 		now := string(res[i])
// 		if now == "-" {
// 			res = res[i+1:] + "-" + res[:i]
// 			break
// 		}
// 	}
// 	return res
// }

func licenseKeyFormatting(S string, K int) string {

	var s string

	for _, c := range S {
		now := string(c)
		if now != "-" {
			s += now
		}
	}
	if len(s) < K {
		return s
	}

	var res string

	idx := len(s) % K
	if idx > 0 {
		res = res + s[:idx] + "-"
	}
	count := K
	for ; idx < len(s); idx++ {
		now := string(s[idx])
		if count > 0 {
			res += now
			count--
			if count == 0 && idx != len(s)-1 {
				res += "-"
				count = K
			}
		}
	}

	return strings.ToUpper(res)
}
