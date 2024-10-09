package main

import "sort"

func countServers(n int, logs [][]int, x int, qs []int) (ans []int) {

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	type q struct {
		i, t int
	}
	qList := make([]*q, 0, len(qs))
	for i, v := range qs {
		qList = append(qList, &q{i, v})
	}
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].t < qList[j].t
	})

	ans = make([]int, len(qs))
	h := make(map[int]int)
	l, r := 0, 0

	for _, q := range qList {
		for r < len(logs) && logs[r][1] <= q.t {
			h[logs[r][0]]++
			r++
		}
		for l < r && logs[l][1] < q.t-x {
			h[logs[l][0]]--
			if h[logs[l][0]] == 0 {
				delete(h, logs[l][0])
			}
			l++
		}
		ans[q.i] = n - len(h)
	}

	return
}
