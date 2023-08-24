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

func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root, nil, nil)
}

func dfs(current, parent, grandparent *TreeNode) int {
	if current == nil {
		return 0
	}
	sum := 0
	if grandparent != nil && grandparent.Val%2 == 0 {
		sum += current.Val
	}
	sum += dfs(current.Left, current, parent)
	sum += dfs(current.Right, current, parent)

	return sum
}

func main() {
	root := structures.Ints2TreeNode([]int{6, 7, 8, 2, 7, 1, 3, 9, 0, 1, 4, 0, 0, 0, 5})
	fmt.Println(sumEvenGrandparent(root))
}
