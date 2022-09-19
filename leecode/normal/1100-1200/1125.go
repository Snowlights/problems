package _100_1200

// 1129
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	rg, bg := map[int]map[int]bool{}, map[int]map[int]bool{}
	for _, e := range redEdges {
		if _, ok := rg[e[0]]; ok {
			rg[e[0]][e[1]] = true
		} else {
			rg[e[0]] = map[int]bool{e[1]: true}
		}
	}
	for _, e := range blueEdges {
		if _, ok := bg[e[0]]; ok {
			bg[e[0]][e[1]] = true
		} else {
			bg[e[0]] = map[int]bool{e[1]: true}
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	visited := map[node]bool{}
	que := []node{node{0, 0, 0}}
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		if ans[cur.id] < 0 {
			ans[cur.id] = cur.steps
		}
		for nxt, _ := range rg[cur.id] {
			if visited[node{nxt, 0, 1}] || cur.color == 1 {
				continue
			}
			visited[node{nxt, 0, 1}] = true
			que = append(que, node{nxt, cur.steps + 1, 1})
		}
		for nxt, _ := range bg[cur.id] {
			if visited[node{nxt, 0, 2}] || cur.color == 2 {
				continue
			}
			visited[node{nxt, 0, 2}] = true
			que = append(que, node{nxt, cur.steps + 1, 2})
		}
	}

	return ans
}

type node struct {
	id, steps, color int // red:color=1, blue:color=2
}

// 1137
func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	switch n {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	default:
		for i := 3; i <= n; i++ {
			a, b, c = b, c, a+b+c
		}
		return c
	}
}

// 1143
func longestCommonSubsequence(t1, t2 string) int {

	m, n := len(t1), len(t2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// t1长度为i时，t2长度为j时候的最长长度
	// 相等+1，不想等，之前的最长
	for i, v := range t1 {
		for j, vv := range t2 {
			if v == vv {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
