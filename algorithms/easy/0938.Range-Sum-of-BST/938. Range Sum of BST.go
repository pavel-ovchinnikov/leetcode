package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var sum = 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	l := rangeSumBST(root.Left, low, high)
	r := rangeSumBST(root.Right, low, high)
	return sum + l + r
}

func main() {
	root := structures.Ints2TreeNode([]int{10, 5, 15, 3, 7, -1, 18})
	fmt.Println(rangeSumBST(root, 7, 15))
}
