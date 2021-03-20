package _06_反转链表

import (
	"fmt"
	"testing"
)

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1589 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestRever(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result := reverseList(l1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

	l1 = &ListNode{
		Val: 1,
	}

	result = reverseList(l1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

	l1 = nil

	result = reverseList(l1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//判断链表是否为空
	if head == nil {
		return nil
	}

	//如果只有一个节点.直接返回
	if head.Next != nil && head.Next.Next == nil {
		return head
	}

	//当前的头结点
	cur := &ListNode{
		Val: head.Val,
	}

	head = head.Next
	for head != nil {
		temp := &ListNode{
			Val: head.Val,
		}
		//当前节点作为头结点
		temp.Next = cur

		//更换头结点
		cur = temp

		head = head.Next

	}

	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
