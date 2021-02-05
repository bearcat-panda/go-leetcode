package __整数反转

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学
// 👍 2415 👎 0

func TestReverse(t *testing.T) {
	var x = 123
	fmt.Println(reverse(x))
	fmt.Println(reverseOne(x))
	x = -123
	fmt.Println(reverse(x))
	fmt.Println(reverseOne(x))
	x = 120
	fmt.Println(reverse(x))
	fmt.Println(reverseOne(x))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	//math.MaxInt32

	//转换为字符串. 如果是- .记录下来
	s := strconv.FormatInt(int64(x), 10)
	var symbol string
	switch string(s[0]) {
	case "-":
		symbol = "-"
		s = s[1:]
	}

	//进行翻转
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	//判断范围是否超出
	y, _ := strconv.Atoi(string(b))
	if y > math.MaxInt32 {
		return 0
	}
	//还原
	if symbol != "" {
		s = symbol + string(b)
	} else {
		s = string(b)
	}

	y, _ = strconv.Atoi(s)
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func reverseOne(x int) int {
	var res int
	for x != 0 {
		if int(int32(x)) != x {
			return 0
		}
		integer := x / 10   //整数
		remainder := x % 10 //余数
		x = integer

		//余数*10 + 其他余数
		res = res*10 + remainder
	}

	return res

}
