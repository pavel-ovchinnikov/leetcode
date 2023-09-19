package main

import (
	"fmt"
	"strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var handler func(root *TreeNode, path string)
	handler = func(root *TreeNode, path string) {
		if root == nil {
			return
		}

		if len(path) != 0 {
			path += "->"
		}

		path += strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil {
			res = append(res, path)
			return
		}

		handler(root.Left, path)
		handler(root.Right, path)
	}

	handler(root, "")

	return res
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println(binaryTreePaths(root))
}
