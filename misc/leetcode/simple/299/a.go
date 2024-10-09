package simple_299

import "math"

// 1
func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i == n-j-1 {
				if grid[i][j] == 0 || grid[i][n-j-1] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}

	return true
}

// 2
func countHousePlacements(n int) int {

	if n == 1 {
		return 4
	}
	if n == 2 {
		return 9
	}

	mod := int(1e9 + 7)
	pre, next := 2, 3
	for i := 3; i <= n; i++ {
		tmp := (pre + next) % mod
		pre = next % mod
		next = tmp % mod
	}

	return next * next % mod
}

// 3

func cal(nums1, nums2 []int) int {
	diff, total, c := 0, 0, 0
	for i, v := range nums1 {
		total += v
		diff = max(diff+nums2[i]-v, 0)
		c = max(diff, c)
	}
	return total + c
}

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(cal(nums1, nums2), cal(nums2, nums1))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	xor := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock
	}
	dfs(0, -1)
	isAncestor := func(x, y int) bool { return in[x] < in[y] && in[y] <= out[x] }

	ans := math.MaxInt32
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			var x, y, z int
			if isAncestor(i, j) { // i 是 j 的祖先节点
				x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
			} else if isAncestor(j, i) { // j 是 i 的祖先节点
				x, y, z = xor[i], xor[i]^xor[j], xor[0]^xor[j]
			} else { // 删除的两条边分别属于两颗不相交的子树
				x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
			}
			ans = min(ans, max(max(x, y), z)-min(min(x, y), z))
			if ans == 0 {
				return 0 // 提前退出
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
