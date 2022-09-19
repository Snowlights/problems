package _500_1600

import "sort"

// 1552
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	check := func(num int) bool {
		pre := position[0]
		cnt := 1
		for i := 1; i < n; i++ {
			if position[i]-pre >= num {
				pre = position[i]
				cnt++
			}
		}
		return cnt < m
	}
	idx := sort.Search(position[len(position)-1], func(i int) bool {
		return check(i)
	})
	return idx - 1
}

// 1557
func findSmallestSetOfVertices(n int, edges [][]int) []int {

	indegree := make(map[int]int)
	g := make(map[int][]int)
	vis := make(map[int]bool)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] != 0 || vis[i] {
			continue
		}
		ans = append(ans, i)
		vis[i] = true
		q := []int{i}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, v2 := range g[v] {
					if !vis[v2] {
						q = append(q, v2)
						vis[v2] = true
					}
				}
			}
		}

	}
	return ans
}
