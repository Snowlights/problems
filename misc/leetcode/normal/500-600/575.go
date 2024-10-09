package _00_600

import "sort"

// 576
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {

	const (
		mod int = 1e9 + 7
	)
	// 移动i步之后在x,y的方案数
	dp := make([][][]int, maxMove+1)
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	ans := 0
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	dp[0][startRow][startColumn] = 1
	for i := 0; i < maxMove; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[i][j][k]
				for _, d := range dir {
					x, y := j+d[0], k+d[1]
					if 0 <= x && x < m && 0 <= y && y < n {
						dp[i+1][x][y] = (dp[i+1][x][y] + count) % mod
					} else {
						ans = (ans + count) % mod
					}
				}
			}
		}
	}
	return ans
}

// 581
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	left, right := 0, len(nums)-1
	for nums[left] == numsSorted[left] {
		left++
	}
	for nums[right] == numsSorted[right] {
		right--
	}
	return right - left + 1
}

// 583
func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	lcs := dp[m][n]
	return m + n - lcs*2
}

// 584
// select name from customer where referee_id != 2 or referee_id is null

// 589
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
	}
	return
}

// 595
// select name, population, area from World where area >= 3000000 or population >= 25000000
