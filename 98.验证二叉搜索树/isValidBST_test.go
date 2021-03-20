package _8_验证二叉搜索树

import (
	"fmt"
	"testing"
)

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 递归
// 👍 963 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestIsValid(t *testing.T) {
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

	fmt.Println(isValidBST(l1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && node.Left.Val < node.Val {
		return true
	}

	if node.Right != nil && node.Right.Val > node.Val {
		return true
	}

	return order(node.Left) && order(node.Right)
}

func order(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && node.Left.Val < node.Val {
		return true
	}

	if node.Right != nil && node.Right.Val > node.Val {
		return true
	}

	return order(node.Left) && order(node.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
