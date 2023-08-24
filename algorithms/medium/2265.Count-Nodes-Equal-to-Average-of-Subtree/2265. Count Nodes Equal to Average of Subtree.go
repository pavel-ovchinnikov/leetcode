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

func averageOfSubtree(root *TreeNode) int {
	_, _, res := dfs(root)
	return res
}

func dfs(root *TreeNode) (sum, number, res int) {
	if root == nil {
		return 0, 0, 0
	}

	sumLeft, numberLeft, resLeft := dfs(root.Left)
	sumRight, numberRight, resRight := dfs(root.Right)

	sum = sumLeft + sumRight + root.Val
	number = numberLeft + numberRight + 1
	res = resLeft + resRight

	if sum/number == root.Val {
		res++
	}

	return sum, number, res
}

func main() {
	root := structures.Ints2TreeNode([]int{4, 8, 5, 0, 1, 0, 6})
	fmt.Println(averageOfSubtree(root))
}
