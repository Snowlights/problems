package double_90

// 1
func oddString(words []string) string {
	n := len(words[0])
	h := make(map[string][]string)
	b := make([]byte, n-1)
	for _, word := range words {
		for i := 0; i < n-1; i++ {
			b[i] = word[i] - word[i+1]
		}
		h[string(b)] = append(h[string(b)], word)
	}
	for _, v := range h {
		if len(v) == 1 {
			return v[0]
		}
	}
	return ""
}

// 2
func twoEditWords(queries []string, dictionary []string) []string {
	ans := []string{}
	for _, q := range queries {
		for _, v := range dictionary {
			res := 0
			for i := range v {
				if q[i] != v[i] {
					res++
				}
			}
			if res <= 2 {
				ans = append(ans, q)
				break
			}
		}
	}

	return ans
}

// 3

func destroyTargets(nums []int, space int) int {
	h := make(map[int][]int)
	for _, v := range nums {
		h[v%space] = append(h[v%space], v)
	}
	ans, l := 0, 0
	for _, v := range h {
		mm := v[0]
		for _, vv := range v {
			mm = min(mm, vv)
		}

		if ll := len(v); ll > l || (ll == l && ans > mm) {
			l = ll
			ans = mm
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 4
// 数组的第二大元素
// AtCoderABC140E 数组中下一个大的元素(第二个大)的位置
func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	s, t := []int{}, []int{}
	for i, v := range nums {
		for len(t) > 0 && v > nums[t[len(t)-1]] {
			ans[t[len(t)-1]] = v
			t = t[:len(t)-1]
		}

		j := len(s) - 1
		for j >= 0 && v > nums[s[j]] {
			j--
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}
	return ans
}
