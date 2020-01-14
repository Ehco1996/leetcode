/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms字符串相乘
 * Medium (40.85%)
 * Likes:    252
 * Dislikes: 0
 * Total Accepted:    39.1K
 * Total Submissions: 94.7K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 示例 1:
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * 说明：
 *
 *
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 *
 *
 */

// @lc code=start

func BigMulti(a, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}
	// string转换成[]byte，容易取得相应位上的具体值
	bsi := []byte(a)
	bsj := []byte(b)

	temp := make([]int, len(bsi)+len(bsj))
	//两数相乘，结果位数不会超过两乘数位数和，即temp的长度最大为 len(num1)+len(num2)
	// 选最大的，免得位数不够
	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			// 对应每个位上的乘积，直接累加存入 temp 中相应的位置
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}

	//统一处理进位
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10 //对该结果进位（进到前一位）
		temp[i] = temp[i] % 10    //对个位数保留
	}

	// a 和 b 较小的时候，temp的首位为0
	// 为避免输出结果以0开头，需要去掉temp的0首位
	if temp[0] == 0 {
		temp = temp[1:]
	}
	//转换结果：将[]int类型的temp转成[]byte类型,
	//因为在未处理进位的情况下，temp每位的结果可能超过255(go中，byte类型实为uint8，最大为255),所以temp选用[]int类型
	//但在处理完进位后，不再会出现溢出
	res := make([]byte, len(temp)) //res 存放最终结果的ASCII码

	for i := 0; i < len(temp); i++ {
		res[i] = byte(temp[i] + '0')
	}

	return string(res)
}

func multiply(num1 string, num2 string) string {
	// 没什么卵用 为啥不能直接用乘法？ 那干脆也别用啥高级语言了，天天用纸带打孔算了
	// 其实就是用代码模拟手动算乘法的步骤 不写
	return BigMulti(num1, num2)
}

// @lc code=end

