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

func findTilt(root *TreeNode) int {
	var helper func(root *TreeNode) (int, int)
	helper = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftVal, sum1 := helper(root.Left)
		rightVal, sum2 := helper(root.Right)
		return (root.Val + leftVal + rightVal), (abs(leftVal-rightVal) + sum1 + sum2)
	}
	_, res := helper(root)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := structures.Ints2TreeNode([]int{21, 7, 14, 1, 1, 2, 2, 3, 3})
	fmt.Println(findTilt(root))
}
