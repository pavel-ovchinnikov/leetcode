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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}

	maxIndex := l
	maxNum := nums[maxIndex]
	for i := l + 1; i < r+1; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	root := &TreeNode{Val: maxNum}
	root.Left = helper(nums, l, maxIndex-1)
	root.Right = helper(nums, maxIndex+1, r)
	return root
}

func main() {

	fmt.Println(structures.Tree2Postorder(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})))
}
