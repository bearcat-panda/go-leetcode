package _3_罗马数字转整数

import (
	"fmt"
	"testing"
)

//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + I
//I 。
//
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5
// 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//
//
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
//
//
// 示例 1:
//
//
//输入: "III"
//输出: 3
//
// 示例 2:
//
//
//输入: "IV"
//输出: 4
//
// 示例 3:
//
//
//输入: "IX"
//输出: 9
//
// 示例 4:
//
//
//输入: "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//
//
// 示例 5:
//
//
//输入: "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
//
//
//
// 提示：
//
//
// 1 <= s.length <= 15
// s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
// 题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
// 题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
// IC 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
// 关于罗马数字的详尽书写规则，可以参考 罗马数字 - Mathematics 。
//
// Related Topics 数学 字符串
// 👍 1201 👎 0
func TestRomanToInt(t *testing.T) {
	var s = "III"
	fmt.Println(romanToInt(s) == 3, 3, romanToInt(s))

	s = "IV"
	fmt.Println(romanToInt(s) == 4, 4, romanToInt(s))

	s = "IX"
	fmt.Println(romanToInt(s) == 9, 9, romanToInt(s))

	s = "LVIII"
	fmt.Println(romanToInt(s) == 58, 58, romanToInt(s))

	s = "MCMXCIV"
	fmt.Println(romanToInt(s) == 1994, 1994, romanToInt(s))
}

//leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	var res int
	var flag bool
	for len(s) > 0 {
		/**
		先截取2个字符串. 存在就继续.不存在就截取1个
		*/
		var c string
		if len(s) >= 2 {
			c = string(s[:2]) //先截取2个字符串.
		}

	LABEL:
		for i, v := range roman {
			if v == c {
				res += intSlice[i]
				s = s[len(c):]
				flag = true
			}
		}

		if !flag {
			c = string(s[:1])
			goto LABEL
		}
		flag = false

	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/**
从右至左遍历罗马数字的字符,将罗马字符映射为对应的阿拉伯数字,
若当前的数字大于或等于前一个数字,则加,否则减
*/
func romanToIntOne(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var pre, res int

	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur >= pre {
			res += cur
		} else {
			res -= cur
		}

		pre = cur
	}

	return res
}
