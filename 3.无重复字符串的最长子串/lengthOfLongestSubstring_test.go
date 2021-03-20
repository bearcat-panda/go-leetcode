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
	length = lengthOfLongesSubstring(s)
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
func lengthOfLongestSubstring_(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

//方法一:双指针
/**
思路和算法
	我们先用一个例子考虑如何在较优的时间复杂度内通过本题.
	我们不妨以示例一中的字符串abcabcbb为例,找出每一个字符开始的,不包含
重复字符串的最长子串,那么其中最长的那个字符串即为答案.对于示例一中的字符串,
我们列举出这些结果,其中括号中表示选中的字符以及最长的字符串:
	以 (a)bcabcbb}(a)bcabcbb 开始的最长字符串为 (abc)abcbb}(abc)abcbb；
	以 a(b)cabcbb}a(b)cabcbb 开始的最长字符串为 a(bca)bcbb}a(bca)bcbb；
	以 ab(c)abcbb}ab(c)abcbb 开始的最长字符串为 ab(cab)cbb}ab(cab)cbb；
	以 abc(a)bcbb}abc(a)bcbb 开始的最长字符串为 abc(abc)bb}abc(abc)bb；
	以 abca(b)cbb}abca(b)cbb 开始的最长字符串为 abca(bc)bb}abca(bc)bb；
	以 abcab(c)bb}abcab(c)bb 开始的最长字符串为 abcab(cb)b}abcab(cb)b；
	以 abcabc(b)b}abcabc(b)b 开始的最长字符串为 abcabc(b)b}abcabc(b)b；
	以 abcabcb(b)}abcabcb(b) 开始的最长字符串为 abcabcb(b)}abcabcb(b)。

	发现了什么?如果我们依次递增的枚举子串的起始位置,那么子串的结束位置也是递增的!
这里的原因在于,假设我们选择字符串中的第K个字符作为起始位置,并且得到了不包含重复
字符的最长子串的结束位置为rk.那么当我们选择第k+1个字符作为起始位置时,首先从k+1到
rk的字符显示是不重复的,并且由于少了原本的第k个字符,我们可以尝试继续增大rk,直到右侧
出现了重复字符为止.
	这样一来,我们就可以使用[双指针]来解决这个问题了:
	.我们使用两个指针表示字符串的某个子串(的左右边界).其中左指针代表着上文中[枚举子串的起始位置],
而右指针即为上文中的rk:
	.在每一步的操作中,我们会将左指针向右移动一格,表示我们开始枚举下一个字符作为起始位置,然后我们
可以不断地向右移动右指针,但需要保证这两个指针对应的子串中没有重复的字符.在移动结束后,这个子串就对应
着以左指针开始的,不包含重复字符的最长子串.我们记录下这个子串的长度;
	.在枚举结束后,我们找到的最长的子串的长度即为答案
在上面的流程中，我们还需要使用一种数据结构来判断 是否有重复的字符，常用的数据结构为哈希集合（即 C++ 中的 std::unordered_set，
Java 中的 HashSet，Python 中的 set, JavaScript 中的 Set）。在左指针向右移动的时候，我们从哈希集合中移除一个字符，在右指针向右
移动的时候，我们往哈希集合中添加一个字符。

至此，我们就完美解决了本题。


*/
func lengthOfLongesSubstring(s string) int {
	//哈希集合,记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)

	//rk右指针,初始值为-1,相当于我们在字符串的左侧,还没有开始移动
	//ans最长字符串的长度
	rk, ans := -1, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			//左指针向右移动一格,移除一格字符
			delete(m, s[i-1])
		}
		//右指针向右移动.并且当前字符不存在
		for rk+1 < n && m[s[rk+1]] == 0 {
			//不断的向右移动指针
			m[s[rk+1]]++ //当前字符的值不为0
			rk++
		}
		//第i到rk个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)

	}
	return ans
}

//复杂度分析
/**
.时间复杂度:O(N),其中N是字符串的长度.左指针和右指针分别会遍历整个字符串一次.
.空间复杂度:O(|Σ|),其中Σ表示字符集(及字符串中可以出现的字符),|Σ|表示字符集的大小.
在本题中没有明确说明字符集,因此可以默认为所有ASCll码在[0,128)内的字符,即|Σ|=128.
我们需要用到哈希集合来存储出现过的字符,而字符最多有|Σ|个,因此空间复杂度为O(|Σ|)
*/
