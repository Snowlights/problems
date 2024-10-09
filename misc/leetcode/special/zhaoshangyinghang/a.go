package zhaoshangyinghang

// 1
func deleteText(article string, index int) string {

	if article[index] == ' ' {
		return article
	}

	s, e := index, index
	for 0 <= s && article[s] != ' ' {
		s--
	}
	if s == -1 {
		s++
	}

	for e < len(article) && article[e] != ' ' {
		e++
	}
	for e+1 < len(article) && article[e+1] == ' ' {
		e++
	}
	ans := article[:s] + article[e:]
	for len(ans) > 0 && ans[0] == ' ' {
		ans = ans[1:]
	}

	return ans
}

// 2

func numFlowers(roads [][]int) int {

	g := make(map[int]int)
	for _, v := range roads {
		g[v[0]]++
		g[v[1]]++
	}
	ans := 0
	for _, v := range g {
		ans = max(ans, v+1)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
