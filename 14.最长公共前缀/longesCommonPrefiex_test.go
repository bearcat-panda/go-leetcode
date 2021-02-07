package _4_最长公共前缀

import (
	"fmt"
	"strings"
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 0 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// Related Topics 字符串
// 👍 1445 👎 0

func TestLongestCommonPrefiex(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

//最长公共前缀. 先找到最短的字符串
//最段的字符串. 依次截断

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	var short string //最短的字符串

	//找到最短的字符串
	for _, v := range strs {
		if short == "" || len(short) > len(v) {
			short = v
		}
	}

	for i := len(short) - 1; i >= 0; i-- {
		//计数器
		var count int
		for _, v := range strs {
			if strings.HasPrefix(v, short[:i]) {
				count++
			}
		}

		if count == len(strs) {
			return short[:i]
		}
	}

	return ""
}

//leetcode submit region end(Prohibit modification and deletion)
