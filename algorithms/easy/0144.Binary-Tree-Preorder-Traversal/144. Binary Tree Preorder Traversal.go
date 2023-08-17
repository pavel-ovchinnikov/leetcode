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

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)

	return result
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 2, 3, 4})
	fmt.Println(preorderTraversal(root))
}
