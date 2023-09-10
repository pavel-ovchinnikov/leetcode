package main

import (
	"fmt"
	"math"

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

func isValidBST(root *TreeNode) bool {
	minVal, maxVal := math.MinInt, math.MaxInt
	return isValid(root, minVal, maxVal)

}

func isValid(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minVal {
		return false
	}
	if root.Val >= maxVal {
		return false
	}

	isLeft := isValid(root.Left, minVal, root.Val)
	isRight := isValid(root.Right, root.Val, maxVal)

	return isLeft && isRight
}

func main() {
	root := structures.Ints2TreeNode([]int{2, 1, 3})
	fmt.Println(isValidBST(root))
}
