package _2_删除排序链表中的重复元素

import (
	"fmt"
	"testing"
)

//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
// 示例 1:
//
// 输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//
//
// 示例 2:
//
// 输入: 1->1->1->2->3
//输出: 2->3
// Related Topics 链表
// 👍 479 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestDeleteDuplicates(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	result := deleteDuplicates(l1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var maps = make(map[int]int)

	cur := head
	for cur != nil {
		if v, ok := maps[cur.Val]; ok {
			maps[cur.Val] = v + 1
		} else {
			maps[cur.Val] = 1
		}

		cur = cur.Next
	}

	result := &ListNode{
		Next: &ListNode{},
	}
	temp := result
	for head != nil {
		if v, ok := maps[head.Val]; ok && v == 1 {
			temp.Next = &ListNode{Val: head.Val}
			temp = temp.Next
		}
		head = head.Next
	}

	return result.Next

}

//leetcode submit region end(Prohibit modification and deletion)
