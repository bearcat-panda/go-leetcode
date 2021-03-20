package _0_有效的括号

import (
	"fmt"
	"strings"
	"testing"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串
// 👍 2143 👎 0

func TestIsValid(t *testing.T) {
	var s = "()"
	fmt.Println(isValid(s))
	fmt.Println(isValidOne(s))

	s = "()[]{}"
	fmt.Println(isValid(s))
	fmt.Println(isValidOne(s))

	s = "(]"
	fmt.Println(isValid(s))
	fmt.Println(isValidOne(s))

	s = "([)]"
	fmt.Println(isValid(s))
	fmt.Println(isValidOne(s))

	s = "{[]}"
	fmt.Println(isValid(s))
	fmt.Println(isValidOne(s))
}

//遇到相同的替换
func isValidOne(s string) bool {
	s = strings.ReplaceAll(s, "()", "")
	s = strings.ReplaceAll(s, "[]", "")
	s = strings.ReplaceAll(s, "{}", "")
	if s == "" {
		return true
	}

	return false
}

/**
从左只有扫描.
遇到"(", 他能否找到匹配,还不得而知.
遇到"[", 他能否找到匹配,还不得而知.
遇到"{", 他能否找到匹配,还不得而知.
所以.需要一个数据结构去暂存. 需要匹配的
	不能忘了扫描过的左括号,他们等着被匹配,需要一个容器暂存
	这个容器为什么是栈?
	当遇到右括号,我们期待他匹配掉 [最近出现的左括号]
	于是匹配容器中的[最近出现的左括号],后者不用等待匹配了,可以离开容器.
	他是后进的.现在先出了,所以是栈

像 “对对碰”
匹配了就拿掉，如果最后清空了栈中的左括号，则有效。如果栈中还剩左括号未匹配，则无效。
*/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	//为了方便对应
	var maps = map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var array []int32

	for _, v := range s {
		if _, ok := maps[v]; ok {
			array = append(array, v)
		} else if v == ')' || v == ']' || v == '}' {
			if maps[array[len(array)-1]] == v {
				array = array[:len(array)-1]
			}
		}

	}
	return len(array) == 0
}
