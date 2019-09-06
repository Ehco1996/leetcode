// package easy
import "strconv"
import "math"

/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第N个数字
 *
 * https://leetcode-cn.com/problems/nth-digit/description/
 *
 * algorithms
 * Easy (32.59%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    4.8K
 * Total Submissions: 14.2K
 * Testcase Example:  '3'
 *
 * 在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。
 *
 * 注意:
 * n 是正数且在32为整形范围内 ( n < 2^31)。
 *
 * 示例 1:
 *
 *
 * 输入:
 * 3
 *
 * 输出:
 * 3
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * 11
 *
 * 输出:
 * 0
 *
 * 说明:
 * 第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
 *
 *
 */


func intPow(x,y int)int{
	return int(math.Pow(float64(x),float64(y)))
}

 func findNthDigit(n int) int {


	// 找到第几个区间
	i := 1
	for n > i*intPow(10,i-1)*9 {
		n-=i*(intPow(10,i-1))*9 //小于区间的值要减去，知道得到确定的区间
		i++
	}

	// 找到这个数
	num :=intPow(10,i-1) + (n-1)/i

	s := strconv.Itoa(num)
	var idx int
	if r:=n %i;r==0{
		idx = i -1
	}else{
		idx = r-1
	}

	res,_ :=strconv.Atoi(string(s[idx]))
	return res

}
