package _6_最接近的三数之和

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
// 👍 681 👎 0

func TestThreeSumClosest(t *testing.T) {
	var nums = []int{-1, 2, 1, -4}
	var target = 1

	fmt.Println(threeSumClosest(nums, target))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {

	var res int
	//给数组排序
	sort.Ints(nums)
	//固定一个数. 其余的按双指针进行计算
	for index, value := range nums {

		left := index + 1
		right := len(nums) - 1

		for left < right {
			sum := value + nums[left] + nums[right]
			if target-sum == 0 { //相等
				res = sum
				return res
			}
			if sum-target > 0 { //右边大
				if target-sum < target-res {
					res = sum
				}

				right--
			}
			if sum-target < 0 { //左边大
				if target-sum < target-res {
					res = sum
				}

				left++
			}
		}

	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
