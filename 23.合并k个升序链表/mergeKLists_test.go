package _3_合并k个升序链表

import (
	"fmt"
	"testing"
)

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 堆 链表 分治算法
// 👍 1130 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestMergeKLists(t *testing.T) {
	//var lists = [[1,4,5],[1,3,4],[2,6]]
	lists := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	res := mergeKLists(lists)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

//每次合并两个数组.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res = &ListNode{}
	head := res

	for i := len(lists) - 1; i >= 1; i-- {
		for {
			if lists[i] == nil && lists[i-1] == nil {
				break
			} else if lists[i] != nil && lists[i-1] == nil {
				lists[i-1].Next = lists[i]
				break
			} else if lists[i] != nil && lists[i-1] != nil {
				if lists[i].Val > lists[i-1].Val {
					lists[i-1].Next = lists[i]
					lists[i] = lists[i].Next
				} else {
					lists[i].Next = lists[i-1]
					lists[i-1] = lists[i-1].Next
				}
			}

		}
	}

	return lists[0]
}
