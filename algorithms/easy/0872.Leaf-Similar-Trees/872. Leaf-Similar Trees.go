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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := []int{}
	leafs2 := []int{}

	var dfs func(root *TreeNode, sequence *[]int)
	dfs = func(root *TreeNode, sequence *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*sequence = append(*sequence, root.Val)
		}

		dfs(root.Left, sequence)
		dfs(root.Right, sequence)
	}

	dfs(root1, &leafs1)
	dfs(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i, num := range leafs1 {
		if num != leafs2[i] {
			return false
		}
	}

	return true
}

func main() {
	root1 := structures.Ints2TreeNode([]int{1, 2, 3})
	root2 := structures.Ints2TreeNode([]int{1, 3, 2})
	fmt.Println(leafSimilar(root1, root2))
}
