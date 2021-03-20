package _98_划分为k个相等的子集

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 示例 1：
//
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
//
//
// 提示：
//
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
//
// Related Topics 递归 动态规划
// 👍 324 👎 0

func TestCanParKSub(t *testing.T) {
	var nums = []int{4, 3, 2, 3, 5, 2, 1}
	fmt.Println(canPartitionKSubsets(nums, 4))
}

/**
如果数组长度小于传入的值.
直接返回false

找到一个最大的值.作为结果.
	从数组中去掉这个值. 计算其他的.
这个最大的值加一个最小的值  依次计算
*/

//leetcode submit region begin(Prohibit modification and deletion)
func canPartitionKSubsets1(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}

	sort.Ints(nums)

	var left, right = 0, len(nums) - 1

	for left < len(nums)/2 {

		var max, min = nums[len(nums)-1], nums[0]
		var result = max
		if left == 0 {
			result = max
			right--
		} else {
			result = max + min
			left++
			right--
		}

		for left < right {
			if nums[left]+nums[right] > result {
				right--
			} else if nums[left]+nums[right] < right {
				left++
			} else {
				right--
				left++
				k--
			}
		}
	}

	if k == 1 {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

var target int

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)

	sum := 0
	for _, v := range nums {
		sum += v
	}

	target = sum / k
	if k*target < sum || nums[len(nums)-1] > target {
		return false
	}

	l := len(nums)
	for nums[l-1] == target {
		l--
		k--
	}

	return false
}

func helper(newGroups []int, nums []int, l int) bool {
	if l == 0 {
		return true
	}

	limitIndex := l - 1

	for i := 0; i < len(newGroups); i++ {
		temp := newGroups[i] + nums[limitIndex]
		if temp <= target {
			if temp < target && temp+nums[0] > target {
				continue
			}
			newGroups[i] = temp
			if helper(newGroups, nums[:limitIndex-1], l-1) {
				return true
			}
			newGroups[i] = temp - nums[limitIndex]

		}
	}

	return false
}
