package main

func maxSum(a []int, m int, k int) int64 {
	sum, ans := 0, 0
	cnt := make(map[int]int)
	for _, v := range a[:k-1] {
		cnt[v]++
		sum += v
	}
	for i, v := range a[k-1:] {
		cnt[v]++
		sum += v
		if len(cnt) >= m {
			ans = max(sum, ans)
		}
		cnt[a[i]]--
		if cnt[a[i]] == 0 {
			delete(cnt, a[i])
		}
		sum -= a[i]
	}

	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
