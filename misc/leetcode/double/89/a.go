package double_89

import "fmt"

// 1
func count(time string, limit int) (ans int) {
next:
	for i := 0; i < limit; i++ {
		for j, c := range fmt.Sprintf("%02d", i) {
			if time[j] != '?' && byte(c) != time[j] {
				continue next
			}
		}
		ans++
	}
	return
}

func countTime(time string) int {
	return count(time[:2], 24) * count(time[3:], 60)
}

// 2
func productQueries(n int, queries [][]int) []int {
	mod := int(1e9 + 7)
	prefix := make([]int, 0)
	for n > 0 {
		val := n & -n
		n -= val
		prefix = append(prefix, int(val))
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		val := 1
		for _, v := range prefix[q[0] : q[1]+1] {
			val *= v
			val %= mod
		}
		ans[i] = val
	}
	return ans
}

// 3
func minimizeArrayValue(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func componentValue(nums []int, edges [][]int) int {
	g, total, mx := make(map[int][]int), 0, 0
	for _, v := range nums {
		total += v
		mx = max(mx, v)
	}
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	var target int

	var dfs func(int, int) int
	dfs = func(fa, son int) int {
		t := nums[son]
		for _, v := range g[son] {
			if v == fa {
				continue
			}
			res := dfs(son, v)
			if res == -1 {
				return -1
			}
			t += res
		}

		if t > target {
			return -1
		} else if t == target {
			return 0
		}
		return t
	}

	for i := total / mx; ; i-- {
		if total%i == 0 {
			target = total / i
			if dfs(-1, 0) == 0 {
				return i - 1
			}
		}
	}
}
