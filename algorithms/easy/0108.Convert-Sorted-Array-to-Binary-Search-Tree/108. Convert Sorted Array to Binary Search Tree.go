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

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	node := &TreeNode{Val: nums[mid], Left: nil, Right: nil}
	node.Left = helper(nums, start, mid-1)
	node.Right = helper(nums, mid+1, end)

	return node
}

func main() {
	head := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(structures.Tree2Inorder(head))
}
