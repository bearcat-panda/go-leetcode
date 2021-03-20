package __回文数

import (
	"fmt"
	"testing"
)

//给你一个整数 x ，如果 x 是一个回文整数，返回 ture ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
// 示例 4：
//
//
//输入：x = -101
//输出：false
//
//
//
//
// 提示：
//
//
// -231 <= x <= 231 - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学
// 👍 1380 👎 0

func TestIsPalindrome(t *testing.T) {
	x := 121
	fmt.Println(isPalindrome(x))

	x = -121
	fmt.Println(isPalindrome(x))

	x = 10
	fmt.Println(isPalindrome(x))

	x = -101
	fmt.Println(isPalindrome(x))
}

//不用字符串. 这不就是整数的反转吗
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var res int
	var tmp = x
	for tmp != 0 {
		integer := tmp / 10 //整数
		rember := tmp % 10  //余数
		tmp = integer

		res = res*10 + rember
	}

	return x == res
}
