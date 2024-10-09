package main

func stringCount(n int) int {
	mod := int(1e9 + 7)
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	return ((pow(26, n)-
		pow(25, n-1)*(75+n)+
		pow(24, n-1)*(72+n*2)-
		pow(23, n-1)*(23+n))%mod + mod) % mod
}
