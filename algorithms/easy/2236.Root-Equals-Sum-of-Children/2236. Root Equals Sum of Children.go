package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	fmt.Println(checkTree(structures.Ints2TreeNode([]int{10, 4, 6})))
	fmt.Println(checkTree(structures.Ints2TreeNode([]int{5, 3, 1})))
	fmt.Println(checkTree(structures.Ints2TreeNode([]int{1, 2, 3})))
}
