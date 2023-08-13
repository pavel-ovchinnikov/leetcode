package main

import (
	"fmt"
	"math"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func main() {
	root := structures.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(maxDepth(root))
}
