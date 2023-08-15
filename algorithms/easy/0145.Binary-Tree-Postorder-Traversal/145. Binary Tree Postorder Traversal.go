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

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	result = append(postorderTraversal(root.Left), result...)

	return result
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 0, 2, 3})
	fmt.Println(postorderTraversal(root))
}
