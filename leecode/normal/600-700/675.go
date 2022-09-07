package _00_700

import "sort"

// 692
func topKFrequent(words []string, k int) []string {

	type node struct {
		word  string
		count int
	}
	nodeList, nodeMap := make([]*node, 0, len(words)), map[string]*node{}
	for _, word := range words {
		n, ok := nodeMap[word]
		if !ok {
			n = &node{word: word}
			nodeList = append(nodeList, n)
			nodeMap[word] = n
		}
		n.count++
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].count > nodeList[j].count || (nodeList[i].count == nodeList[j].count && nodeList[j].word > nodeList[i].word)
	})

	ans := make([]string, 0, k)
	for _, v := range nodeList[:k] {
		ans = append(ans, v.word)
	}
	return ans
}

// 695
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		grid[i][j] = 0
		tmp := 1
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
				tmp += dfs(x, y)
			}
		}
		return tmp
	}
	for i, v := range grid {
		for j, vv := range v {
			if vv == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
