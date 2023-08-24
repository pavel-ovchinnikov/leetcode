package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func bstToGst(root *TreeNode) *TreeNode {
	bfs(root, 0)
	return root
}

func bfs(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	sum = bfs(root.Right, sum)
	sum += root.Val
	root.Val = sum
	sum = bfs(root.Left, sum)

	return sum
}

func main() {
	root := structures.Ints2TreeNode([]int{30, 36, 21, 36, 35, 26, 15, 0, 0, 0, 33, 0, 0, 0, 8})
	fmt.Println(structures.Tree2Postorder(bstToGst(root)))
}
