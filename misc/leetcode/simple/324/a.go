package simple_324

// 1
func similarPairs(words []string) int {

	ans := 0
	h := make(map[int][26]int)
	for i, v := range words {
		t := [26]int{}
		for _, vv := range v {
			t[vv-'a'] = 1
		}
		h[i] = t
	}

	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if h[i] == h[j] {
				ans++
			}
		}
	}

	return ans
}

// 2
func smallestValue(n int) int {
	ans, last := n, n
	for {
		tmp := 0
		for last != 1 {
			for j := 2; j <= last; j++ {
				if last%j == 0 {
					tmp += j
					last /= j
					break
				}
			}
		}

		if tmp >= ans {
			break
		}
		ans = tmp
		last = tmp
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 3

func isPossible(n int, edges [][]int) bool {

	g := make(map[int][]int)
	for _, e := range edges {
		from, to := e[0], e[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	h := make([]int, 0)
	for i := 1; i <= n; i++ {
		if len(g[i])%2 == 1 {
			h = append(h, i)
		}
	}

	if len(h) > 4 || len(h)%2 == 1 {
		return false
	}
	equal := func(a, b int) bool {
		for _, to := range g[h[a]] {
			if to == h[b] {
				return false
			}
		}
		return true
	}
	line := func(a, b int) bool {
		a, b = h[a], h[b]
	loop:
		for i := 1; i <= n; i++ {
			if i == a || i == b {
				continue
			}
			for _, to := range g[i] {
				if to == a || to == b {
					continue loop
				}
			}
			return true
		}
		return false
	}
	if len(h) == 2 {
		return equal(0, 1) || line(0, 1)
	}

	return len(h) == 0 || (equal(0, 1) && equal(2, 3)) ||
		(equal(0, 2) && equal(1, 3)) ||
		(equal(0, 3) && equal(1, 2))
}

// 4
func cycleLengthQueries(n int, queries [][]int) []int {

	ans := make([]int, len(queries))
	for i, v := range queries {
		from, to := v[0], v[1]
		h, idx := make(map[int]int), 0
		for from != 0 {
			h[from] = idx
			idx++
			from /= 2
		}
		if val, ok := h[to]; ok {
			ans[i] = val + 1
			continue
		}
		tIdx := 0
		for to != 0 {
			if val, ok := h[to]; ok {
				ans[i] = tIdx + val + 1
				break
			}
			to /= 2
			tIdx++
		}
	}
	return ans
}
