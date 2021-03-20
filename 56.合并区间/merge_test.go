package _6_合并区间

import "sort"

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104
//
// Related Topics 排序 数组
// 👍 847 👎 0

//先对数组进行排序.
//判断当前的数组中的最小值  是否比 结果数组中的最后一个数组的最大值 大.
// 大-更新之前的值为当前的最大值. 否则直接添加

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	for i := 0; i < len(intervals); i++ {

	}

}

//leetcode submit region end(Prohibit modification and deletion)
