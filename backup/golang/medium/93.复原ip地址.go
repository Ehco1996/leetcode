package medium

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (45.34%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 32.6K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 *
 * 示例:
 *
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 *
 */
func restoreIpAddresses(s string) []string {
	var res []string
	dfs(&res, s, []string{}, 0)
	return res
}

func dfs(res *[]string, s string, path []string, position int) {
	if len(path) == 4 {
		if position == len(s) {
			ip := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			*res = append(*res, ip)
		}
		return
	}
	for _, wd := range getNextWords(s, position) {
		path = append(path, wd)
		dfs(res, s, path, position+len(wd))
		path = path[:len(path)-1]
	}
}

func getNextWords(s string, position int) (words []string) {
	var tmp string
	for i := position; i < len(s); i++ {
		tmp += string(s[i])
		num, _ := strconv.Atoi(tmp)
		if len(tmp) > 3 || num > 255 || len(tmp) >= 2 && string(tmp[0]) == "0" {
			break
		}
		words = append(words, tmp)
	}
	return
}
