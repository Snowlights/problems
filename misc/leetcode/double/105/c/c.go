package main

import "sort"

func maxStrength(a []int) int64 {
	if len(a) == 1 {
		return int64(a[0])
	}

	ans, flag, cal := 1, 0, false
	sort.Ints(a)
	for _, v := range a {
		if v < 0 {
			if flag == 0 {
				flag = v
				continue
			}
			ans *= flag * v
			flag = 0
			cal = true
			continue
		}
		if v > 0 {
			ans *= v
			cal = true
		}
	}

	if !cal {
		return 0
	}
	return int64(ans)
}
