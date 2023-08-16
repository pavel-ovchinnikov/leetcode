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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil &&
		q != nil &&
		p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right) {
		return true
	}

	return false
}

func main() {
	root1 := structures.Ints2TreeNode([]int{1, 2, 3})
	root2 := structures.Ints2TreeNode([]int{1, 2, 3})
	fmt.Println(isSameTree(root1, root2))
}
