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
