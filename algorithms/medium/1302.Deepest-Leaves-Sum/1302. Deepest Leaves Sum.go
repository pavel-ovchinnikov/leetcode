package main

import "github.com/pavel-ovchinnikov/LeetCode/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func deepestLeavesSum(root *TreeNode) int {
	sum := 0

	if root == nil {
		return sum
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)

		sum = 0
		for i := 0; i < l; i++ {
			node := queue[i]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
	}
	return sum
}

func main() {
	root := structures.Ints2TreeNode([]int{6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5})
	deepestLeavesSum(root)
}
