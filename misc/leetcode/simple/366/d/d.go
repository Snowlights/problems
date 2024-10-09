package main

const mod int = 1e9 + 7

func maxSum(a []int, k int) (ans int) {
	cnt := [32]int{}
	for _, v := range a {
		for i := range cnt {
			cnt[i] += v >> i & 1
		}
	}
	for ; k > 0; k-- {
		v := 0
		for i := range cnt {
			if cnt[i] > 0 {
				v |= 1 << i
				cnt[i]--
			}
		}
		ans = (ans + v*v) % mod
	}

	ans = (ans%mod + mod) % mod
	return
}
