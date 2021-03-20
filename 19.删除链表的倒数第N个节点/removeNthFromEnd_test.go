package _9_删除链表的倒数第N个节点

import (
	"fmt"
	"testing"
)

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// Related Topics 链表 双指针
// 👍 1201 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestRemoveNthFromEnd(t *testing.T) {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	result := removeNthFromEnd(listNode, 2)
	for result != nil { //统计长度
		fmt.Print(result.Val)
		result = result.Next
	}

	fmt.Println()
	listNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	result = removeNthFromEndTwo(listNode, 2)
	for result != nil { //统计长度
		fmt.Print(result.Val)
		result = result.Next
	}

	fmt.Println()
	listNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	result = removeNthFromEndThree(listNode, 2)
	for result != nil { //统计长度
		fmt.Print(result.Val)
		result = result.Next
	}

}

/**
找到对应的节点. 去掉就好
从头到尾遍历节点.获取长度.
然后找到length -n 个节点. 这个节点的下一个节点
	就是需要去除的节点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length int
	var temp = head
	for temp != nil { //统计长度
		length++
		temp = temp.Next
	}

	temp = head
	var count = 1
	for temp != nil {
		if count == length-n { //找到被删除节点的前一个节点
			temp.Next = temp.Next.Next
		}
		count++
		temp = temp.Next
	}

	return head
}

/**
时间复杂度: O(L),其中L是链表的长度
空间复杂度;O(1)
*/

//方法二:栈
/**
我们可以在遍历链表的同时将所有节点依次入栈.根据栈[先进后出]的原则,
我们弹出栈的第n个节点就是需要删除的节点.并且目前栈顶的节点就是待删除
节点的前驱节点.
*/
func removeNthFromEndTwo(head *ListNode, n int) *ListNode {
	var array []*ListNode
	var temp = head

	for temp != nil { //将所有节点添加到数组中
		array = append(array, temp)
		temp = temp.Next
	}

	//被删除节点的前置节点
	pre := array[len(array)-1-n]
	pre.Next = pre.Next.Next

	return head
}

/**
时间复杂度: O(L),其中L是链表的长度
空间复杂度:O(L),其中L是链表的长度,主要为栈的开销.
*/

//方法三:双指针
/**
创建两个指针. 使其中一个指针先移动n个节点.
然后2个指针同时向后移动.  当有一个指针到达末尾节点
的时候. 前面的指针就是要被删除的指针的前置节点
*/

func removeNthFromEndThree(head *ListNode, n int) *ListNode {
	var left, right = head, head

	var count = 0
	for right != nil {
		if count >= n+1 { //如果右节点已经超过n个节点
			left = left.Next //左节点开始移动
		}
		count++
		right = right.Next
	}

	left.Next = left.Next.Next
	return head
}

/**
时间复杂度: O(L),其中L是链表的长度
空间复杂度: O(1)
*/
