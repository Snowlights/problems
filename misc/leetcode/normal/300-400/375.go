package _00_400

import (
	"math"
	"math/rand"
	"strings"
)

// 375
func getMoneyAmount(n int) int {

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				tmp := max(dp[i][k-1], dp[k+1][j]) + k
				dp[i][j] = min(dp[i][j], tmp)
			}
		}
	}

	return dp[1][n]
}

// 376
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// 377
func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

// 380
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.indices[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

// 384
type Solution struct {
	h             map[int]bool
	original, arr []int
}

func Constructor384(nums []int) Solution {
	h := make(map[int]bool)
	for _, v := range nums {
		h[v] = true
	}
	return Solution{h, nums, nums}
}

func (this *Solution) Reset() []int {
	this.arr = this.original
	return this.arr
}

func (this *Solution) Shuffle() []int {
	val := make([]int, 0, len(this.original))
	for v, _ := range this.h {
		val = append(val, v)
	}
	this.arr = val
	return this.arr
}

// 392
func isSubsequence(s string, t string) bool {

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j == len(t) {
			break
		}
		i++
		j++
	}
	return i == len(s)
}

// 394
func decodeString(s string) string {
	numQ, cQ := make([]int, 0), make([]string, 0)
	num, ans := 0, ""
	for _, v := range s {
		if '0' <= v && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			numQ = append(numQ, num)
			num = 0
			cQ = append(cQ, ans)
			ans = ""
		} else if v == ']' {
			count := numQ[len(numQ)-1]
			numQ = numQ[:len(numQ)-1]
			str := cQ[len(cQ)-1]
			cQ = cQ[:len(cQ)-1]
			ans = str + strings.Repeat(ans, count)
		} else {
			ans += string(v)
		}
	}

	return ans
}

// 399
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type pair struct {
		to  string
		val float64
	}
	g, h := make(map[string][]pair), make(map[string]bool)
	for i, e := range equations {
		g[e[0]] = append(g[e[0]], pair{to: e[1], val: values[i]})
		g[e[1]] = append(g[e[1]], pair{to: e[0], val: 1 / values[i]})
		h[e[0]], h[e[1]] = true, true
	}

	path, vis := []float64{}, make(map[string]bool)
	var bfs func(from, to string, tmp []float64)
	bfs = func(from, to string, tmp []float64) {
		for _, v := range g[from] {
			if v.to == to {
				tmp = append(tmp, v.val)
				path = make([]float64, len(tmp))
				copy(path, tmp)
				return
			}
			if !vis[v.to] {
				vis[v.to] = true
				bfs(v.to, to, append(tmp, v.val))
				vis[v.to] = false
			}
		}
	}

	ans := make([]float64, 0)
	for _, q := range queries {
		from, to := q[0], q[1]
		if !h[from] || !h[to] {
			ans = append(ans, -1)
			continue
		}
		path, vis = []float64{}, make(map[string]bool)
		vis[from] = true
		bfs(from, to, nil)
		if len(path) == 0 {
			ans = append(ans, -1)
		} else {
			val := 1.0
			for _, v := range path {
				val *= v
			}
			ans = append(ans, val)
		}
	}
	return ans
}
