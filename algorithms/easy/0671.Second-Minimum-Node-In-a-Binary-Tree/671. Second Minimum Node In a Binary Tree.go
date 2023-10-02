package main

import (
	"fmt"
	"math"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func findSecondMinimumValue(root *TreeNode) int {
	first := root.Val
	second := math.MaxInt
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val != first && root.Val < second {
			second = root.Val
		}
		helper(root.Left)
		helper(root.Right)
	}

	helper(root)

	if second == math.MaxInt {
		return -1
	}

	return second
}

func main() {
	root := structures.Ints2TreeNode([]int{2, 2, 5, 1, 3, 5, 7})
	fmt.Println(findSecondMinimumValue(root))
}
