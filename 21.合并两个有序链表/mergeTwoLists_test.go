package _1_合并两个有序链表

import (
	"fmt"
	"sort"
	"testing"
)

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表
// 👍 1520 👎 0

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res := mergeTwoLists(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res = mergeTwoListsOne(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//获取到所有数字
//先给数字排序
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var intArray []int                     //存储对应的值
	var listMap = make(map[int][]ListNode) //存储只对应的节点

	for l1 != nil {
		intArray = append(intArray, l1.Val)
		listMap[l1.Val] = append(listMap[l1.Val], ListNode{Val: l1.Val})
		l1 = l1.Next
	}

	for l2 != nil {
		intArray = append(intArray, l2.Val)
		listMap[l2.Val] = append(listMap[l2.Val], ListNode{Val: l2.Val})
		l2 = l2.Next
	}

	var res = &ListNode{}
	temp := res
	sort.Ints(intArray)

	for i, v := range intArray {
		if i > 0 && v == intArray[i-1] {
			continue
		}

		for _, lv := range listMap[v] {
			temp.Val = lv.Val
			temp.Next = &ListNode{}
			temp = temp.Next

		}

	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//递归计算
func mergeTwoListsOne(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	temp := res
	var digui = func(res, l1, l2 *ListNode) {}

	digui = func(temp, l1, l2 *ListNode) {
		if l1 == nil && l2 == nil {
			return
		} else if l1 != nil && l2 == nil {

			temp.Next = l1 //不能这样写. 因为temp的next还没有赋值.
			return
		} else if l1 == nil && l2 != nil {
			temp.Next = l2
			return

		} else {
			if l1.Val > l2.Val {
				temp.Next = l2
				l2 = l2.Next
			} else {
				temp.Next = l1
				l1 = l1.Next
			}
		}

		temp = temp.Next

		digui(temp, l1, l2)

	}

	digui(temp, l1, l2)

	return res
}

/**
两个链表头部值较小的一个节点与剩下元素的merge操作结果合并
如果l1或者l2一开始就是空链表,那么没有任何合并操作,所以我们只需要返回空链表.
否则,我们要判断l1和l2哪一个链表的头结点的值更小,然后递归的决定下一个添加到
结果里的节点.
*/
func mergeTwoListsTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoListsOne(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsOne(l2.Next, l1)
		return l2
	}
}

func mergeTwoListsThree(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head

	for {
		if l1 == nil && l2 == nil {
			return nil
		} else if l1 != nil && l2 == nil {
			head.Next = l1
			return res.Next
		} else if l1 == nil && l2 != nil {
			head.Next = l2
			return res.Next
		} else {
			if l1.Val > l2.Val {
				head.Next = l2
				l2 = l2.Next
			} else {
				head.Next = l1
				l1 = l1.Next
			}

			head = head.Next
		}
	}

	return res.Next
}
