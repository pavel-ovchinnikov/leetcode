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

func sumOfLeftLeaves(root *TreeNode) int {
	var helper func(root *TreeNode, isLeft bool) int
	helper = func(root *TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}
		sum := 0
		if isLeft && root.Left == nil && root.Right == nil {
			sum += root.Val
		}
		sum += helper(root.Right, false)
		sum += helper(root.Left, true)
		return sum
	}

	return helper(root, false)
}

func main() {
	root := structures.Ints2TreeNode([]int{3, 9, 20, 2, 10, 15, 7})
	fmt.Println(sumOfLeftLeaves(root))
}
