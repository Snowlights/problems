package double_81

import "strings"

// 1
func countAsterisks(s string) int {
	ans := 0
	for i, v := range strings.Split(s, "|") {
		if i%2 == 0 {
			ans += strings.Count(v, "*")
		}
	}
	return ans
}

// 2
func maximumXOR(nums []int) int {
	tmp := 0
	for _, v := range nums {
		tmp |= v
	}
	return tmp
}

// 3
func countPairs(n int, edges [][]int) int64 {
	ans := int64(0)
	g, vis := make(map[int][]int), make(map[int]bool)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		q, l := []int{i}, 1
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, vv := range g[v] {
					if !vis[vv] {
						vis[vv] = true
						l++
						q = append(q, vv)
					}
				}
			}
		}
		ans += int64(l * (n - l))
	}

	return ans / 2
}

// 4
func distinctSequences(n int) int {
	ans, mod := 0, int(1e9+7)
	if n == 1 {
		return 6
	}
	f := [1e4 + 5][7][7]int{}
	for last := 1; last <= 6; last++ {
		for last1 := 1; last1 <= 6; last1++ {
			if last != last1 && gcd(last, last1) == 1 {
				f[2][last][last1] = 1
			}
		}
	}

	for i := 3; i <= n; i++ {
		for last := 1; last <= 6; last++ {
			for last1 := 1; last1 <= 6; last1++ {
				if last != last1 && gcd(last1, last) == 1 {
					for last2 := 1; last2 <= 6; last2++ {
						if last != last2 {
							f[i][last][last1] = (f[i][last][last1] + f[i-1][last1][last2]) % mod
						}
					}
				}
			}
		}
	}
	for _, row := range f[n] {
		for _, v := range row {
			ans = (ans + v) % mod
		}
	}

	return ans
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
