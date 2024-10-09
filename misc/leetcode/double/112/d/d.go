package main

import "sort"

const (
	mod int = 1e9 + 7
	mx      = 30
)

func countKSubsequencesWithMaxBeauty(s string, k int) (ans int) {

	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int(i) % mod
	}
	// 阶乘
	pow := func(x, n int) int {
		//x %= mod
		res := int(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	// n阶乘的逆元
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int(i) % mod
	}
	// C(n, k) 组合：n中取k，有多少种取法
	C := func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	cnt, cc := make([]int, 26), make(map[int]bool)
	for _, v := range s {
		cnt[int(v-'a')]++
		cc[int(v-'a')] = true
	}
	if k > 26 || len(cc) < k {
		return 0
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	val, count := cnt[k-1], 0
	ans = 1
	for _, v := range cnt {
		if v == val {
			count++
			continue
		}
		if v < val {
			break
		}
		k--
		ans = ans * v % mod
	}

	ans = ans * C(count, k) % mod * pow(val, k) % mod

	ans = (ans%mod + mod) % mod
	return
}
