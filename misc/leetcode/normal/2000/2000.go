package _000

// 2014
func longestSubsequenceRepeatedK(s string, k int) (ans string) {
	n := len(s)
	pos, cnt := [26]int{}, [26]int{}
	for i := range pos {
		pos[i] = len(s)
	}
	nxt := make([][26]int, n)
	for i := n - 1; i >= 0; i-- {
		nxt[i] = pos
		pos[s[i]-'a'] = i
		cnt[s[i]-'a']++
	}

	a := make([]byte, 0, n)
	for i := 25; i >= 0; i-- {
		for ; cnt[i] >= k; cnt[i] -= k {
			a = append(a, 'a'+byte(i))
		}
	}
	if len(a) == 0 {
		return
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					if do(ids[:r]) {
						return
					}
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}

	for m := len(a); m > 0 && ans == ""; m-- {
		permutations(len(a), m, func(ids []int) (Break bool) {
			t := make([]byte, m)
			for i, id := range ids {
				t[i] = a[id]
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for ; j < len(t); j++ {
				i = nxt[i][t[j]-'a']
				if i == len(s) {
					break
				}
			}
			if j >= k*m {
				ans = string(t)
				return true
			}
			return false
		})
	}
	return
}
