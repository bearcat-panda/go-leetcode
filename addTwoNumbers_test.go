package go_leetcode

import (
	"fmt"
	"testing"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
// 👍 5416 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result := addTwoNumbers(l1, l2)

	for {
		fmt.Println(result.Val)
		result = result.Next
		if result == nil {
			return
		}
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	temp = result
	for {

		if result == nil {
			result = &ListNode{
				Next: &ListNode{},
			}
			//r = result
			temp = result
		} else {
			if temp.Next == nil {
				temp.Next = &ListNode{}
			}
			temp = temp.Next
		}

		var carry int //进位
		if l1 != nil && l2 != nil {
			temp.Val = (l1.Val + l2.Val) % 10
			carry = (l1.Val + l2.Val) / 10
		} else if l1 != nil && l2 == nil {
			temp.Val = l1.Val
		} else if l1 == nil && l2 != nil {
			temp.Val = l2.Val
		}

		if l1 == nil && l2 == nil {
			break
		}

		//r = r.Next

		l1 = l1.Next
		if l1 != nil {
			if carry != 0 {
				l1.Val += carry
				carry = 0
			}
		}

		l2 = l2.Next
		if l2 != nil {
			if carry != 0 {
				l2.Val += carry
				carry = 0
			}
		}

	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
