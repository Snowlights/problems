package main

import "fmt"

func minimumSeconds(a []int) (ans int) {
	m := make(map[int][]int)
	for i, v := range a {
		m[v] = append(m[v], i)
	}

	ans = len(a)
	for _, v := range m {
		tmp := 0
		for i := 1; i < len(v); i++ {
			tmp = max(tmp, v[i]-v[i-1])
		}
		tmp = max(tmp, len(a)-v[len(v)-1]+v[0])
		ans = min(ans, tmp/2)
		fmt.Println(ans)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
