package __最长回文子串

import (
	"fmt"
	"testing"
)

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划
// 👍 3120 👎 0

func TestLongestPalindrome(t *testing.T) {
	var s = "ccbaaabca"
	fmt.Println(longestPalindrome(s))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
	s = "a"
	fmt.Println(longestPalindrome(s))
	s = "ac"
	fmt.Println(longestPalindrome(s))
	s = "babad"
	fmt.Println(longestPalindrome(s))
}

//如果是最长回文子串. 最大的一定是从字符串中间开始的.
//babad  先判断两边的字符是否和中间的字符相同.
// 左边相同 -- . 减到0
//右边相同 ++, 加到0
//如果不同.判断左右两边是否相同. 相同.左-右+

func longestPalindrome(s string) string {
	//字符串的长度
	length := len(s)
	if length == 1 { //只有一个字符
		return s
	}

	if length == 2 { //有两个字符
		if s[0] == s[1] { //并且相同
			return s
		} else { //并且不相同
			return string(s[0])
		}
	}

	left, right := length/2, length/2

	var result = []byte{s[length/2]}
	for {

		//结束条件
		if left == 0 || right == length-1 {
			break
		}
		//左边
		left--
		if s[length/2] == s[left] && left == length/2-1 {
			result = append(result, s[left])
			left--
		}

		//右边
		right++
		if s[length/2] == s[right] && right == length/2+1 {
			result = append(result, s[right])
			right++
		}

		//左右两边相同
		if s[left] == s[right] {
			result = append([]byte{s[left]}, result...)
			result = append(result, s[right])
		}

	}

	return string(result)

}
