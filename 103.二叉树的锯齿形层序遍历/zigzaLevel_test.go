package _03_二叉树的锯齿形层序遍历

import (
	"fmt"
	"testing"
)

//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回锯齿形层序遍历如下：
//
//
//[
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
// 👍 412 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestZigLevelOrder(t *testing.T) {
	l1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	fmt.Println(zigzagLevelOrder(l1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	dfs(root, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

var res [][]int

func dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}
	res = append(res, []int{})
	res[level] = append(res[level], node.Val)

	if level%2 == 0 {

		if node.Right != nil {
			dfs(node.Right, level+1)
		}

		if node.Left != nil {
			dfs(node.Left, level+1)
		}

	} else {
		if node.Left != nil {
			dfs(node.Left, level+1)
		}

		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}

}
