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

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rIndex := len(preorder)

	for i := range preorder {
		if rootVal < preorder[i] {
			rIndex = i
			break
		}
	}

	left := bstFromPreorder(preorder[1:rIndex])
	right := bstFromPreorder(preorder[rIndex:])

	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right}
}

func main() {
	fmt.Println(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
}
