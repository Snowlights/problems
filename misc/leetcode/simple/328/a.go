package simple_328

// 1
func differenceOfSum(nums []int) int {
	var x, y int
	for _, v := range nums {
		x += v
		for v > 0 {
			y += v % 10
			v /= 10
		}
	}
	return abs(x - y)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// 2
func rangeAddQueries(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	for _, q := range queries {
		row1, col1, row2, col2 := q[0], q[1], q[2], q[3]
		for i := row1; i <= row2; i++ {
			for j := col1; j <= col2; j++ {
				mat[i][j]++
			}
		}
	}
	return mat
}

// 二为差分模版 https://leetcode.cn/problems/increment-submatrices-by-one/solution/er-wei-chai-fen-by-endlesscheng-mh0h/
//  func rangeAddQueries(n int, queries [][]int) [][]int {
//	// 二维差分模板
//	diff := make([][]int, n+2)
//	for i := range diff {
//		diff[i] = make([]int, n+2)
//	}
//	update := func(r1, c1, r2, c2, x int) {
//		diff[r1+1][c1+1] += x
//		diff[r1+1][c2+2] -= x
//		diff[r2+2][c1+1] -= x
//		diff[r2+2][c2+2] += x
//	}
//	for _, q := range queries {
//		update(q[0], q[1], q[2], q[3], 1)
//	}
//
//	// 用二维前缀和复原（原地修改）
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= n; j++ {
//			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
//		}
//	}
//	// 保留中间 n*n 的部分，即为答案
//	diff = diff[1 : n+1]
//	for i, row := range diff {
//		diff[i] = row[1 : n+1]
//	}
//	return diff
//}
//

// 3
func countGood(nums []int, k int) int64 {
	var ans int
	n := len(nums)
	s, e, cur := 0, 0, 0
	h := make(map[int]int)
	for e < n {
		for e < n && cur < k {
			cur += h[nums[e]]
			h[nums[e]]++
			e++
		}
		for s <= e && cur >= k {
			ans += n - e + 1
			h[nums[s]]--
			cur -= h[nums[s]]
			s++
		}
	}

	return int64(ans)
}

// 4
func maxOutput(n int, edges [][]int, price []int) int64 {
	ans := 0
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}
	// 返回带叶子的最大路径和，不带叶子的最大路径和
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		p := price[x]
		maxS1, maxS2 := p, 0
		for _, y := range g[x] {
			if y != fa {
				s1, s2 := dfs(y, x)
				// 前面最大带叶子的路径和 + 当前不带叶子的路径和
				// 前面最大不带叶子的路径和 + 当前带叶子的路径和
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p) // 这里加上 p 是因为 x 必然不是叶子
			}
		}
		return maxS1, maxS2
	}
	dfs(0, -1)
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
