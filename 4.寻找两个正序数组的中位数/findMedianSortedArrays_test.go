package __寻找两个正序数组的中位数

import (
	"fmt"
	"sort"
	"testing"
)

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
//
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// 示例 3：
//
// 输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
//
// 示例 4：
//
// 输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
//
// 示例 5：
//
// 输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
//
// Related Topics 数组 二分查找 分治算法
// 👍 3618 👎 0

func TestFindMedianSortedArrays(t *testing.T) {
	var nums1 = []int{1, 3}
	var nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{}
	nums2 = []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{2}
	nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{1, 2, 5}
	nums2 = []int{3, 4, 8}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{1, 3, 6}
	nums2 = []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))

	nums1 = []int{3, 7}
	nums2 = []int{1, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(findMedianSortedArraysBin(nums1, nums2))
	fmt.Println(findMedianSortedArraysDevide(nums1, nums2))
}

//leetcode submit region begin(Prohibit modification and deletion)
//O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...) //合并
	sort.Ints(nums1)                //排序

	length := len(nums1)
	if length == 1 { //只有一个数
		return float64(nums1[0])
	}

	if length%2 == 0 {
		return float64(float64(nums1[length/2-1]+nums1[length/2]) / 2)
	} else {
		return float64(nums1[length/2])
	}

}

//1,2,5  3,4,8	12 |34| 58
//1,3  2,4		1 |23| 4
//1,3,6  2,4  	12 |3| 46
//3,7  1,4		1 |34| 7
//O(log(m+n))  //每次去掉无关的数据.
func findMedianSortedArraysDevide(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 1 { //如果只有一个数  直接返回
		if len(nums1) == 1 {
			return float64(nums1[0])
		}

		if len(nums2) == 1 {
			return float64(nums2[0])
		}
	} else if len(nums1)+len(nums2) == 2 { //如果有两个数.返回平均值
		if len(nums1) == 1 && len(nums2) == 1 {
			return float64(nums1[0]+nums2[0]) / 2
		}

		if len(nums1) == 2 {
			return float64(nums1[0]+nums1[1]) / 2
		}

		if len(nums2) == 2 {
			return float64(nums2[0]+nums2[1]) / 2
		}
	}

	//去掉最大值 和最小值
	n1Min, n1Max := nums1[0], nums1[len(nums1)-1]
	n2Min, n2Max := nums2[0], nums2[len(nums2)-1]
	if n1Min < n2Min {
		nums1 = nums1[1:]
	} else {
		nums2 = nums2[1:]
	}

	if n1Max > n2Max {
		nums1 = nums1[:len(nums1)-1]
	} else {
		nums2 = nums2[:len(nums2)-1]
	}
	return findMedianSortedArraysDevide(nums1, nums2)

}

func findMedianSortedArraysBin(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArraysBin(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
