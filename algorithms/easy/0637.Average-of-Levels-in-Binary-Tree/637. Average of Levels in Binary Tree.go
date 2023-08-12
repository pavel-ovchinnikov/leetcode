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

type Node struct {
	val   int
	level int
}

func averageOfLevels(root *TreeNode) []float64 {
	sumNodeByLevel := []float64{}
	countNodeByLevel := []int{}
	out := make(chan Node)
	go func() {
		dfs(root, 0, out)
		close(out)
	}()

	for node := range out {
		if node.level >= len(sumNodeByLevel) {
			sumNodeByLevel = append(sumNodeByLevel, float64(node.val))
			countNodeByLevel = append(countNodeByLevel, 1)
		} else {
			sumNodeByLevel[node.level] += float64(node.val)
			countNodeByLevel[node.level] += 1.0
		}
	}

	res := make([]float64, len(sumNodeByLevel))
	for i := range sumNodeByLevel {
		res[i] = sumNodeByLevel[i] / float64(countNodeByLevel[i])
	}
	return res
}

func dfs(root *TreeNode, level int, out chan Node) {
	if root == nil {
		return
	}
	out <- Node{root.Val, level}
	dfs(root.Left, level+1, out)
	dfs(root.Right, level+1, out)
}

func main() {
	root := structures.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(averageOfLevels(root))
}
