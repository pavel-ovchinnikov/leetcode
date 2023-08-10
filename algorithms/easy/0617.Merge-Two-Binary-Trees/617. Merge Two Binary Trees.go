package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	left := mergeTrees(getLeft(root1), getLeft(root2))
	right := mergeTrees(getRight(root1), getRight(root2))

	return &TreeNode{
		Val:   sumTreeNode(root1, root2),
		Left:  left,
		Right: right}
}

func sumTreeNode(root1 *TreeNode, root2 *TreeNode) (sum int) {
	if root1 != nil {
		sum += root1.Val
	}

	if root2 != nil {
		sum += root2.Val
	}

	return sum
}

func getLeft(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return root.Left
}
func getRight(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return root.Right
}

func main() {
	root1 := structures.Ints2TreeNode([]int{1, 3, 2, 5})
	root2 := structures.Ints2TreeNode([]int{2, 1, 3, 0, 4, 0, 7})
	fmt.Println(structures.Tree2Preorder(mergeTrees(root1, root2)))
}
