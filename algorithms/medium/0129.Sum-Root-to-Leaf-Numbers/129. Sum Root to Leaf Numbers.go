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

func sumNumbers(root *TreeNode) int {
	res := 0
	var solution func(root *TreeNode, sum int)
	solution = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		sum = sum*10 + root.Val

		if root.Left == nil && root.Right == nil {
			res += sum
			return
		}

		solution(root.Left, sum)
		solution(root.Right, sum)
	}
	solution(root, 0)
	return res
}

func main() {
	root := structures.Ints2TreeNode([]int{4, 9, 0, 5, 1})
	fmt.Println(sumNumbers(root))
}
