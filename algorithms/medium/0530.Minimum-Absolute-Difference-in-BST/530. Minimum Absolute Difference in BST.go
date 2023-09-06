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

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return -1
	}

	minDiff := math.MaxInt
	dfs(root, &minDiff)

	return minDiff
}

func dfs(root *TreeNode, midDiff *int) (minValue, maxValue int) {
	minValue, maxValue = root.Val, root.Val
	if root.Left != nil {
		minV, maxV := dfs(root.Left, midDiff)
		*midDiff = min(*midDiff, root.Val-maxV)
		minValue = minV
	}
	if root.Right != nil {
		minV, maxV := dfs(root.Right, midDiff)
		*midDiff = min(*midDiff, minV-root.Val)
		maxValue = maxV
	}
	return minValue, maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := structures.Ints2TreeNode([]int{4, 2, 6, 1, 3})
	fmt.Println(getMinimumDifference(root))
}
