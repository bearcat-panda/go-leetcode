package _2_08环路检测

import (
	"fmt"
	"testing"
)

//给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的
//位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
//
//
//
// 进阶：
//
//
// 你是否可以不用额外空间解决此题？
//
//
//
// Related Topics 链表
// 👍 64 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func TestCycle(t *testing.T) {

	child := &ListNode{
		Val: 2,
	}

	child.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  -4,
			Next: child,
		},
	}
	node := &ListNode{
		Val:  3,
		Next: child,
	}

	detectCycle(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	maps := make(map[*ListNode]int)

	var index = 0
	for head != nil {
		if v, ok := maps[head]; ok {
			fmt.Println(v, true)
			return nil
		} else {
			maps[head] = index
		}
		head = head.Next
		index++
	}

	fmt.Println(-1, false)
	return nil

}

//leetcode submit region end(Prohibit modification and deletion)

//快慢指针
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
			}
			return p
		}
	}
	return nil
}
