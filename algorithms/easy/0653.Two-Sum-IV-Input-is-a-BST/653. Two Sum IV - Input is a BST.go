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

func findTarget(root *TreeNode, k int) bool {
	mm := make(map[int]struct{})
	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		_, ok := mm[root.Val]
		if ok {
			return true
		}

		mm[k-root.Val] = struct{}{}

		return helper(root.Left) || helper(root.Right)
	}

	return helper(root)
}

func main() {
	root := structures.Ints2TreeNode([]int{5, 3, 7, 2, 4, 6, 8})
	fmt.Println(findTarget(root, 9))
}
