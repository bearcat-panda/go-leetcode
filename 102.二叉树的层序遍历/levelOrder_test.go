package _02_二叉树的层序遍历

import (
	"fmt"
	"testing"
)

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层序遍历结果：
//
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 811 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestTreeNode(t *testing.T) {
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

	result := levelOrder(l1)
	fmt.Println(result)

	result = levelOrder1(l1)
	fmt.Println(result)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	//如果为空返回空数组
	if root == nil {
		return result
	}

	//如果只有一个节点
	if root.Left == nil && root.Right == nil {
		return append(result, []int{root.Val})
	}
	dfs(root, 0)

	return res

}

//leetcode submit region end(Prohibit modification and deletion)

var res [][]int

//按层级放到数组中
func dfs(node *TreeNode, n int) {
	if node != nil {
		res = append(res, []int{})
		res[n] = append(res[n], node.Val)
		n++
		if node.Left != nil {
			dfs(node.Left, n)
		}

		if node.Right != nil {
			dfs(node.Right, n)
		}
	}
}

//不使用递归调用
//每一层放到队列中
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})

		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}

	return res
}
