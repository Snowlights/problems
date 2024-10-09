package _00_900

import "strings"

// 828
func uniqueLetterString(s string) (ans int) {
	sum, last := 0, [26][2]int{}
	for i := range last {
		last[i] = [2]int{-1, -1}
	}
	for i, c := range s {
		c -= 'A'

		sum += i - last[c][0]*2 + last[c][1]

		ans += sum

		last[c][1] = last[c][0]

		last[c][0] = i
	}
	return
}

// 833
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	replace, replaceLen := make([]string, len(s)), make([]int, len(s))
	for i, v := range indices {
		if strings.HasPrefix(s[v:], sources[i]) {
			replace[v] = targets[i]
			replaceLen[v] = len(sources[i])
		}
	}

	ans := strings.Builder{}
	for i := 0; i < len(s); {
		if len(replace[i]) > 0 {
			ans.WriteString(replace[i])
			i += replaceLen[i]
		} else {
			ans.WriteString(string(s[i]))
			i++
		}
	}

	return ans.String()
}

// 834 换根dp
func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n) // g[x] 表示 x 的所有邻居
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	size := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		ans[0] += depth // depth 为 0 到 x 的距离
		size[x] = 1
		for _, y := range g[x] { // 遍历 x 的邻居 y
			if y != fa { // 避免访问父节点
				dfs(y, x, depth+1) // x 是 y 的父节点
				size[x] += size[y] // 累加 x 的儿子 y 的子树大小
			}
		}
	}
	dfs(0, -1, 0) // 0 没有父节点

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] { // 遍历 x 的邻居 y
			if y != fa { // 避免访问父节点
				ans[y] = ans[x] + n - 2*size[y]
				reroot(y, x) // x 是 y 的父节点
			}
		}
	}
	reroot(0, -1) // 0 没有父节点
	return ans
}

// 837
func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, k+maxPts)
	s := 0.0
	for i := k; i < k+maxPts; i++ {
		if i > n {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
		s += dp[i]
	}

	for i := k - 1; i >= 0; i-- {
		dp[i] = s / float64(maxPts)
		s -= dp[i+maxPts]
		s += dp[i]
	}

	return dp[0]
}

// 838
func longestMountain(arr []int) int {
	n := len(arr)
	l, r := make([]int, n), make([]int, n)

	for i, v := range arr {
		if i > 0 && v > arr[i-1] {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && arr[i] > arr[i+1] {
			r[i] = r[i+1] + 1
		} else {
			r[i] = 1
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if l[i] > 1 && r[i] > 1 {
			ans = max(ans, l[i]+r[i]-1)
		}
	}
	return ans
}

// 841
func canVisitAllRooms(rooms [][]int) bool {
	vis := make(map[int]bool)
	vis[0] = true
	q := []int{0}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range rooms[v] {
				if vis[vv] {
					continue
				}
				q = append(q, vv)
				vis[vv] = true
			}
		}
	}
	return len(rooms) == len(vis)
}
