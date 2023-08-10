package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	return dfs(root, nil)
}

func dfs(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}

	res := dfs(root.Left, root)
	root.Left = nil
	root.Right = dfs(root.Right, tail)

	return res
}

func main() {
	root := structures.Ints2TreeNode([]int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9})
	res := increasingBST(root)
	fmt.Println(structures.Tree2Preorder(res))
}
