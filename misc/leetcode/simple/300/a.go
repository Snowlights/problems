package simple_300

// 1
func decodeMessage(key string, message string) string {
	s, h := byte('a'), make(map[byte]byte)
	for _, v := range key {
		if string(v) == " " {
			continue
		}
		if _, ok := h[byte(v)]; !ok {
			h[byte(v)] = s
			s++
		}
	}
	res := make([]byte, len(message))
	for i, v := range message {
		if string(v) == " " {
			res[i] = ' '
			continue
		}
		res[i] = h[byte(v)]
	}
	return string(res)
}

// 2

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
		for j := range arr[i] {
			arr[i][j] = -1
		}
	}
	xstart, xend, ystart, yend := 0, m-1, 0, n-1
	for head != nil {
		for i := ystart; i <= yend && head != nil; i++ {
			arr[xstart][i] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		for i := xstart + 1; i <= xend && head != nil; i++ {
			arr[i][yend] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}

		for i := yend - 1; i >= ystart && head != nil; i-- {
			arr[xend][i] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		for i := xend - 1; i >= xstart+1 && head != nil; i-- {
			arr[i][xstart] = head.Val
			head = head.Next
		}
		if head == nil {
			break
		}
		xstart++
		ystart++
		xend--
		yend--
	}

	return arr
}

// 3
func peopleAwareOfSecret(n int, delay int, forget int) int {

	mod := int(1e9 + 7)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if i < delay {
			dp[i] = dp[i-1]
			continue
		}
		if i >= forget {
			dp[i] = (dp[i-1] + dp[i-delay] - dp[i-forget] + mod) % mod
			continue
		}
		if i >= delay {
			dp[i] = (dp[i-1] + dp[i-delay]) % mod
		}
	}

	if n-1-forget < 0 {
		return int(dp[n-1])
	}

	return int((dp[n-1] - dp[n-1-forget] + mod) % mod)
}

// 4
func countPaths(grid [][]int) int {
	n, m, mod := len(grid), len(grid[0]), int(1e9+7)
	arr := make([][]int, len(grid))
	for i := range arr {
		arr[i] = make([]int, len(grid[0]))
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if arr[i][j] != 0 {
			return arr[i][j]
		}
		tmp := 1
		for _, d := range dir {
			if x, y := i+d[0], j+d[1]; 0 <= x && x < n && 0 <= y && y < m && grid[i][j] < grid[x][y] {
				tmp = (tmp + dfs(x, y)) % mod
			}
		}
		arr[i][j] = tmp
		return arr[i][j]
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return ans
}
