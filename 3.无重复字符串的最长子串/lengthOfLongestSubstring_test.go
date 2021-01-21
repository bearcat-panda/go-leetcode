package __无重复字符串的最长子串

import (
	"fmt"
	"strings"
	"testing"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4834 👎 0

func TestTengthOfLongestSubstring(t *testing.T) {

	var s = "pwwkew"
	length := lengthOfLongestSubstring(s)
	length = lengthOfLongestSubstringBit(s)
	fmt.Println(length)

	s = "abcabcbb"
	length = lengthOfLongestSubstring(s)
	fmt.Println(length)

	s = "bbbbb"
	length = lengthOfLongestSubstring(s)
	fmt.Println(length)

	s = ""
	length = lengthOfLongestSubstring(s)
	fmt.Println(length)
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	// "pwwkew"
	//依次累加字符串 放到map中
	var maps = make(map[string]int)
	var start = 0
	for i, _ := range s {
	LABEL:
		sub := s[start : i+1] //当前子串
		var flag bool         //是否是不重复的子串
		for subI, _ := range sub {
			if strings.Count(sub, sub[subI:subI+1]) > 1 {
				flag = true //重复
				start = i
				goto LABEL

			}
		}

		if !flag { //不包含重复字符串
			maps[sub] = len(sub)
		}

	}

	var max int //获取最大的子串长度
	for _, v := range maps {
		if v > max {
			max = v
		}
		//fmt.Println(k, v)
	}

	return max

}

//解法一 位图
func lengthOfLongestSubstringBit(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool
	//最终结果. 左侧起始位. 右侧起始位
	result, left, right := 0, 0, 0

	for left < len(s) {
		//右侧字符对应的bitSet被标记为true,说明此字符在X位置重复,需要左侧向前移动,
		//直到将X标记为false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			fmt.Println(s[right])
			bitSet[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}

		if left+result >= len(s) || right >= len(s) {
			break
		}
	}

	return result
}

//解法二  滑动窗口
func lengthOfLongestSubstringSliceWindow(s string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

/*apiv1 := router.Group("/api/version1")
{
	//测试项目状态
	apiv1.GET("ping", Ping)

	tagRouter := apiv1.Group("/tag")
	tag := v1.NewTag()
	{
		tagRouter.Get("/list", tag.Get)
	}

}*/
