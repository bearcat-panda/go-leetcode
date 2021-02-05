package _1_盛最多水的容器

import (
	"fmt"
	"testing"
)

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：height = [4,3,2,1,4]
//输出：16
//
//
// 示例 4：
//
//
//输入：height = [1,2,1]
//输出：2
//
//
//
//
// 提示：
//
//
// n = height.length
// 2 <= n <= 3 * 104
// 0 <= height[i] <= 3 * 104
//
// Related Topics 数组 双指针
// 👍 2161 👎 0

func TestMaxArea(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	fmt.Println(maxAreaOne(height))

	height = []int{1, 1}
	fmt.Println(maxArea(height))
	fmt.Println(maxAreaOne(height))

	height = []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea(height))
	fmt.Println(maxAreaOne(height))

	height = []int{1, 2, 1}
	fmt.Println(maxArea(height))
	fmt.Println(maxAreaOne(height))
}

//双循环 算出每一个. 然后算出最大的

func maxArea(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			value := min(height[i], height[j]) * (j - i)
			if value > max {
				max = value
			}
		}
	}

	return max
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

//双指针
//计算结果. 每次移动较小的指针
func maxAreaOne(height []int) int {
	left, right := 0, len(height)-1
	var max int
	for left < right {
		value := min(height[left], height[right]) * (right - left)
		if value > max {
			max = value
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
