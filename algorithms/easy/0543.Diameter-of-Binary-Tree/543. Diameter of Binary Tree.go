package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	dist := 0
	var solution func(*TreeNode) int
	solution = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lDist := solution(root.Left)
		rDist := solution(root.Right)
		dist = max(dist, lDist+rDist)
		return max(lDist, rDist) + 1
	}
	solution(root)
	return dist
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println(diameterOfBinaryTree(root))
}
