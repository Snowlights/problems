package simple_320

import "sort"

// 1
func unequalTriplets(nums []int) int {
	h := make(map[int]int)
	for _, v := range nums {
		h[v]++
	}
	a, b := 0, len(nums)
	ans := 0
	for _, v := range h {
		b -= v
		ans += a * b * v
		a += v
	}
	return ans
}

// 2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	a, ans := []int{}, [][]int{}

	var dfs func(o *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		a = append(a, n.Val)
		dfs(n.Right)
	}
	dfs(root)

	for _, q := range queries {
		min, max := -1, -1
		j := sort.SearchInts(a, q)
		if j < len(a) {
			max = a[j]
		}
		j = sort.SearchInts(a, q+1) - 1
		if j >= 0 {
			min = a[j]
		}

		ans = append(ans, []int{min, max})
	}
	return ans
}

// 3

func minimumFuelCost(roads [][]int, seats int) int64 {
	g := make(map[int][]int)
	for _, r := range roads {
		from, to := r[0], r[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	ans := 0
	var dfs func(fa, i int) int
	dfs = func(fa, i int) int {
		size := 1

		for _, v := range g[i] {
			if v == fa {
				continue
			}
			size += dfs(i, v)
		}

		if i != 0 {
			ans += (size + seats - 1) / seats
		}
		return size
	}
	dfs(-1, 0)
	return int64(ans)
}

// 4
// 问题中的变量
// n s的长度
// k 分割的段数
//
// 问题替换
// 使用s的前j个字符，划分为i段的方案数

// 最后一步发生了什么
// 取出了 长度为x的字符串，分割为一个字符串

// 去掉最后一步，划分子问题
// 使用s的前j-x个字符， 划分为i-1段的方案数

// 定义状态
// f[i][j] 划分为i段且使用s的前j个字符的方案数
// f[i][j] += f[i-1][j`]
// j` 是第i段开始的下表
// j - j` + 1 >= minLength
// s[j`] 是质数，且 s[j] 不是质数

// 初始值
// f[0][0] = 1
// ans = f[k][n]
func beautifulPartitions(s string, k int, l int) int {
	const (
		mod int = 1e9 + 7
	)
	n := len(s)

	isPrime := func(v byte) bool {
		return v == '2' || v == '3' || v == '5' || v == '7'
	}
	if k*l > n || !isPrime(s[0]) || isPrime(s[n-1]) {
		return 0
	}

	canPartition := func(j int) bool {
		return j == 0 || j == n || !isPrime(s[j-1]) && isPrime(s[j])
	}
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i := 1; i <= k; i++ {
		sum := 0
		// 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
		for j := i * l; j+(k-i)*l <= n; j++ {
			if canPartition(j - l) { // j'=j-l 双指针
				sum = (sum + f[i-1][j-l]) % mod
			}
			if canPartition(j) {
				f[i][j] = sum
			}
		}
	}
	return f[k][n]
}
