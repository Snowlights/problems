package main

import "math/bits"

const mod = 10

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

const mx = 100_000

var f, invF, p2, p5 [mx + 1]int

func init() {
	f[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		// 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2
		// 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}
		f[i] = f[i-1] * x % mod
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}

	invF[mx] = pow(f[mx], 3) // 欧拉定理
	for i := mx; i > 0; i-- {
		x := i
		x >>= bits.TrailingZeros(uint(x))
		for x%5 == 0 {
			x /= 5
		}
		invF[i-1] = invF[i] * x % mod
	}
}

var pow2 = [4]int{2, 4, 8, 6}

func comb(n, k int) int {
	res := f[n] * invF[k] * invF[n-k]
	e2 := p2[n] - p2[k] - p2[n-k]
	if e2 > 0 {
		res *= pow2[(e2-1)%4]
	}
	if p5[n]-p5[k]-p5[n-k] > 0 {
		res *= 5
	}
	return res
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%mod == 0
}
