package main

import (
	"problems/testutil/leetcode"
	"slices"
)

func kthLargestPerfectSubtree(root *leetcode.TreeNode, k int) int {
	hs := []int{}
	var dfs func(*leetcode.TreeNode) int
	dfs = func(node *leetcode.TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH < 0 || leftH != rightH {
			return -1
		}
		hs = append(hs, leftH+1)
		return leftH + 1
	}
	dfs(root)

	if k > len(hs) {
		return -1
	}
	slices.Sort(hs)
	return 1<<hs[len(hs)-k] - 1
}
