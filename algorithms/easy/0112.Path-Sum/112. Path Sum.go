package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	sum := targetSum - root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func main() {
	root := structures.Ints2TreeNode([]int{5, 4, 8, 11, 1, 13, 4, 7, 2, 1, 1, 1, 1})
	fmt.Println(hasPathSum(root, 22))
}
