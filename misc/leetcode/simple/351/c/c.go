package main

const mod int = 1e9 + 7

func numberOfGoodSubarraySplits(a []int) (ans int) {

	ans, pre := 1, -1
	for i, v := range a {
		if v == 1 {
			if pre != -1 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre == -1 {
		return 0
	}
	ans = (ans%mod + mod) % mod
	return
}
