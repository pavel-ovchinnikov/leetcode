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

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil && node.Val != val {
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

func main() {
	root := structures.Ints2TreeNode([]int{4, 2, 7, 1, 3})
	out := searchBST(root, 2)
	fmt.Println(structures.Tree2Preorder(out))

}
