package _2_括号生成

import (
	"fmt"
	"testing"
)

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 回溯算法
// 👍 1552 👎 0

func TestGenerateParenthesis(t *testing.T) {
	var n = 1
	fmt.Println(generateParenthesis(n))
	fmt.Println(generateParenthesisOne(n))

	n = 2
	fmt.Println(generateParenthesis(n))
	fmt.Println(generateParenthesisOne(n))

	n = 3
	fmt.Println(generateParenthesis(n))
	fmt.Println(generateParenthesisOne(n))

	n = 4
	fmt.Println(generateParenthesis(n))
	fmt.Println(generateParenthesisOne(n))

}

//只考虑.外.左.右
//需要去重
func generateParenthesis(n int) []string {
	var array = make([]string, 0)
	var res = []string{}

	var digui func([]string, int) []string //定义一个函数
	digui = func(array []string, n int) []string {
		if n == 0 {
			return nil
		} else if n == 1 { //如果是1. 直接返回
			array = append(array, "()")
			res = array
			return array
		}
		array = digui(array, n-1)

		if len(array) != 0 {
			res = []string{}
			var outer, left, right string //外 左 右
			for _, v := range array {
				outer = "(" + v + ")"
				left = "()" + v
				right = v + "()"

				//如果不存在.就添加
				if !isContain(res, outer) {
					res = append(res, outer)
				}
				if !isContain(res, left) {
					res = append(res, left)
				}
				if !isContain(res, right) {
					res = append(res, right)
				}
			}
			array = res

		}
		return array
	}

	digui(array, n)
	return res
}

func isContain(array []string, value string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}

/**
不停地选括号,要么选左括号,要么选右括号
并且是有约束的:
	只有(有剩的,就可以选, ((((这么选.都还不能判为非法
	当剩下的)比(多时,才可以选),否则,)不能选,选了就非法了

*/

func generateParenthesisOne(n int) []string {
	res := []string{}

	var dfs func(lRemain int, rRemain int, path string)

	dfs = func(lRemain int, rRemain int, path string) {
		//结束条件
		if 2*n == len(path) {
			res = append(res, path)
			return
		}

		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}
