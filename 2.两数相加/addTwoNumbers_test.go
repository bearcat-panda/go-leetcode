package __两数相加

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

	result := addTwoNumbersTwo(l1, l2)

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
	for {

		if result == nil { //第一次临时节点 = 当前节点
			result = &ListNode{
				Next: &ListNode{},
			}
			//r = result
			temp = result
		} else {
			//所有的值都不为nil 或者 0
			if temp.Next == nil && l1 != nil && l2 != nil {
				temp.Next = &ListNode{} //添加节点
			}
			temp = temp.Next //走到下一个节点
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

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		//那个节点不为空. 就添加进位
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

func addTwoNumbersTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	var carry int
	for {

		if result == nil {
			result = &ListNode{
				Next: &ListNode{},
			}
			//r = result
			temp = result
		} else {
			if temp.Next == nil && l1 != nil && l2 != nil {
				temp.Next = &ListNode{}
			}
			temp = temp.Next
		}

		if l1 != nil && l2 != nil {
			temp.Val = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
		} else if l1 != nil && l2 == nil {
			temp.Val = l1.Val
		} else if l1 == nil && l2 != nil {
			temp.Val = l2.Val
		}

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		l1 = l1.Next
		l2 = l2.Next

	}
	return result
}

func addTwoNumbersBest(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	n1, n2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
