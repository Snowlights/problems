package _400_1500

import "strings"

// 1456
func maxVowels(s string, k int) int {
	pre, ans := make([]int, len(s)+1), 0
	for i, v := range s {
		pre[i+1] = pre[i]
		if strings.ContainsAny(string(v), "aeiou") {
			pre[i+1]++
		}
		if i+1 >= k {
			ans = max(ans, pre[i+1]-pre[i+1-k])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1462
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, 105)
	for _, v := range prerequisites {
		g[v[0]] = append(g[v[0]], v[1])
	}

	ans := make([]bool, len(queries))
	for i, v := range queries {
		q, vis := []int{v[0]}, make([]bool, 105)
		vis[v[0]] = true

		for len(q) > 0 {
			tmp := q
			q = nil
			for _, from := range tmp {
				if from == v[1] {
					ans[i] = true
					q = nil
					break
				}
				for _, to := range g[from] {
					if !vis[to] {
						q = append(q, to)
						vis[to] = true
					}
				}
			}
		}
	}
	return ans
}
