package main

func beautifulString(s string) int {
	const mod int64 = 998244353
	const mx int = 1e6
	// 逆元的定义
	// mod2 * a % mod = 1, mod2 就是a的逆元
	// mod2 = (mod + 1) / a
	// n的阶乘
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	// 阶乘
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	// n阶乘的逆元
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	// C(n, k) 组合：n中取k，有多少种取法
	C := func(n, k int64) int64 {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}
	var ans, one, zero, d0, d1 int64
	n := int64(len(s))
	if s[0] == '0' {
		d0, zero = 1, 1
	} else {
		d1, one = 1, 1
	}
	ans = 1
	for i := int64(1); i < n; i++ {
		switch s[i] {
		case '0':
			d0 = (2*d0%mod + C(i-1, i-one)) % mod
			d1 = (2*d1%mod - C(i-1, i-one)) % mod
			ans = (ans + d0) % mod
			zero++
		case '1':
			d0 = (2*d0%mod - C(i-1, i-zero)) % mod
			d1 = (2*d1%mod + C(i-1, i-zero)) % mod
			ans = (ans + d1) % mod
			one++
		}
	}
	return int((ans%mod + mod) % mod)
}
