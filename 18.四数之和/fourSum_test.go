package _8_四数之和

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
// Related Topics 数组 哈希表 双指针
// 👍 739 👎 0

func TestFourSum(t *testing.T) {
	var nums = []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	//给数组排序
	sort.Ints(nums)

	var res [][]int
	/**
	每次固定2个数字
	然后利用指针和邻近的数  还有最后的数相加
	根据值的大小. 移动指针
	*/
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			one := nums[i]
			two := nums[j]

			if i > 0 && nums[i] == nums[i-1] { //去重
				continue
			}

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			//如果第一个数就大于0. (排序) 后面的数一定都大于0
			if one > target {
				return res
			}

			//双指针移动
			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := one + two + nums[left] + nums[right]

				switch {
				case sum == target:
					res = append(res, []int{one, two, nums[left], nums[right]})
					left++
					right--
				case sum > target:
					right--
				case sum < target:
					left++
				}
			}
		}

	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
