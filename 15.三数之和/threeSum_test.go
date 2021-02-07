package _5_三数之和

import (
	"fmt"
	"sort"
	"testing"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// Related Topics 数组 双指针
// 👍 2944 👎 0
func TestThreeSum(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSumOne(nums))

	nums = []int{0}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSumOne(nums))

	nums = []int{}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSumOne(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}

			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/**
先给数组排序
先固定一个数, 然后用双指针来查找另外两个数的方式.
所以初始化时,我们选择固定第一个元素.同时将下一个
元素和末尾元素分别设上left和right指针.
*/

func threeSumOne(nums []int) [][]int {
	//从小到大排序
	sort.Ints(nums)

	//返回值
	var res [][]int

	length := len(nums) //数组长度

	//数字不满足3个.直接返回空
	if length < 3 {
		return res
	}

	//开始循环第一个固定值
	for index, value := range nums {
		//如果固定位的值已经大于0,因为已经排好序了,后面的两个指针
		//对应的值也肯定大于0,则和不可能为0,所以返回
		if value > 0 {
			return res
		}

		//排除固定位的重复值
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		//指针初始位置,固定位右边第一个和数组最后一个
		l := index + 1
		r := length - 1

		//开始移动两个指针
		for l < r {
			//判断3个数字之和的三种情况
			sum := value + nums[l] + nums[r]

			switch {
			case sum == 0:
				//将结果加入二元组
				res = append(res, []int{value, nums[l], nums[r]})

				//去重,如果l<r并且下一个数字一样,则继续挪动
				/*for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}*/

				l += 1
				r -= 1
			case sum > 0:
				//如果和大于0,说明right的值太大,需要左移
				r -= 1
			case sum < 0:
				// 如果和小于 0，那就说明 left 的值太小，需要右移
				l += 1
			}
		}
	}
	return res
}
