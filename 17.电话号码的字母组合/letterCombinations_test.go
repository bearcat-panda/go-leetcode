package _7_电话号码的字母组合

import (
	"fmt"
	"testing"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 深度优先搜索 递归 字符串 回溯算法
// 👍 1120 👎 0

func TestLetterCombinationsOne(t *testing.T) {
	var digits = "23"
	fmt.Println(letterCombinationsOne(digits))
	fmt.Println(letterCombinations(digits))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
先拿第一个的数据. 遍历
组合第二个的数组.
*/

func letterCombinations(digits string) []string {
	var maps = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var result = []string{}
	var def = func(res, digits string) {}
	def = func(res, digits string) {
		if len(digits) == 0 { //最后的结果
			result = append(result, res)
			return //函数已经结束. 不返回会数组越界
		}

		for _, v := range maps[string(digits[0])] { //每次只取第一个
			res += string(v)
			def(res, digits[1:])   // 获取后面的
			res = res[:len(res)-1] //获取最开始的字符串
		}
	}

	def("", digits)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func letterCombinationsOne(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var dfs func(res string, digits string)
	dfs = func(res string, digits string) {
		if len(digits) == 0 {
			result = append(result, res)
			return
		}
		for _, v := range mp[string(digits[0])] {
			res = res + string(v)
			dfs(res, digits[1:])
			res = res[:len(res)-1] //回溯
		}
	}
	dfs("", digits)
	return result
}
