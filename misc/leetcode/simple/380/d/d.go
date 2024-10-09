package main

func beautifulIndices(s, a, b string, k int) []int {
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (pos []int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				pos = append(pos, i-lenP+1)
				c = match[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}

	aPos, bPos := kmpSearch(s, a), kmpSearch(s, b)
	ans := make([]int, 0)

	//for _, v := range aPos {
	//	idx := sort.SearchInts(bPos, v)
	//	if idx < len(bPos) && abs(bPos[idx] - v) <= k ||
	//		idx > 0 && abs(bPos[idx-1] - v) <= k{
	//		ans = append(ans, v)
	//	}
	//}

	j, m := 0, len(bPos)
	for _, i := range aPos {
		for j < m && bPos[j] < i-k {
			j++
		}
		if j < m && abs(bPos[j]-i) <= k {
			ans = append(ans, i)
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
