package main

func matrixSumQueries(n int, qs [][]int) int64 {
	ans := 0
	vis := [2]map[int]bool{{}, {}}
	for i := len(qs) - 1; i >= 0; i-- {
		q := qs[i]
		t, idx, v := q[0], q[1], q[2]
		if !vis[t][idx] {
			ans += v * (n - len(vis[t^1]))
			vis[t][idx] = true
		}
	}

	return int64(ans)
}
