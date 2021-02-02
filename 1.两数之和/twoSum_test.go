package __两数之和

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
// 👍 9875 👎 0

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumBest(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	var numsMp = make(map[int]int)
	var result []int
	for i, v := range nums {
		numsMp[v] = i
	}
	//两个键的值是需要的. 添加下标
	for k, v1 := range numsMp {
		if v2, ok := numsMp[target-k]; ok {
			result = append(result, v1)
			result = append(result, v2)
			return result
		}
	}

	return result
}

func twoSumBest(nums []int, target int) []int {
	m := make(map[int]int) // 值-下标
	/*
		说了只有一种答案. 所以遍历第一个结果时. map没有.
		但是遍历到第二个结果时. 第一个结果已经在map中了
	*/
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
//方法一:暴力枚举
//思路及算法
/**
最容易想到的方法是枚举数组中的每一个数 x，寻找数组中是否存在 target - x。

当我们使用遍历整个数组的方式寻找 target - x 时，需要注意到每一个位于 x 之前的元素都已经和 x 匹配过，因此不需要再进行匹配。而每一个元素不能被使用两次，所以我们只需要在 x 后面的元素中寻找 target - x。

*/
//代码
func TwoSumOne(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/**
复杂度分析
时间复杂度：O(N^2)，其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。

空间复杂度：O(1)。
*/

//方法二:哈希表
/**
注意到方法一的时间复杂度较高的原因是寻找target-x的时间复杂度
过高.因此,我们需要一种更优秀的方法,能够快速寻找数组中是否存在
目标元素.如果存在,我们需要找出他的索引.
使用哈希表,,可以将寻找target-x的时间复杂度从O(N)降低到O(1).
这样我们创建一个哈希表,对于每一个x,我们首先查询哈希表中是否存在
target-x,然后将X插入到哈希表中,即可保证不会让X和自己匹配.
*/

//代码
func TwoSumTwo(nums []int, target int) []int {
	var numMap = make(map[int]int, len(nums))
	for i, v := range nums {
		if p, ok := numMap[target-v]; ok {
			return []int{i, p}
		}
		numMap[v] = i
	}
	return nil
}

/**
时间复杂度:O(N),其中N是数组中的元素数量,我们可以
以O(1)寻找target-x
空间复杂度:O(N),其中N是 数组中的元素数量.主要为哈希表的开销
*/
