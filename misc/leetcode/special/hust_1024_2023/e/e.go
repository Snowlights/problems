package main

func subsetCounting(n int64, k int, m int) []int {
	const mod int = 998244353
	pow := func(x, n int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}

	f, mu, g := make([]int, m+1), make([]int, m+1), make([]int, m+1)
	vis, pr, vec := make([]bool, m+1), make([]int, 0, m+1), make([]int, 0, m+1)

	init := func() {
		for i := 2; i <= m; i++ {
			if !vis[i] {
				pr = append(pr, i)
				mu[i] = -1
			}
			for _, v := range pr {
				if v*i > m {
					break
				}
				vis[i*v] = true
				if i%v == 0 {
					break
				}
				mu[i*v] = -mu[i]
			}
		}
	}

	pre := func() {
		for i := 1; i <= m; i++ {
			d := gcd(i, m)
			if m/d&1 == 1 {
				f[i] = pow(2, (d)*int(n))
			}
		}
		for _, v := range pr {
			for j := m / v; j > 0; j-- {
				f[j*v] = (f[j*v] - f[j] + mod) % mod
			}
		}
		for i := 1; i <= m; i++ {
			if m%i == 0 {
				vec = append(vec, i)
			}
		}
	}

	last := func() {
		iv := pow(m, mod-2)
		for i := 1; i <= m; i++ {
			for _, de := range vec {
				if i%(m/de) == 0 {
					g[i] = (g[i] + f[de]*(m/de)%mod) % mod
				}
			}
			g[i] = g[i] * iv % mod
		}
		g[0] = g[m]
		g[m] = 0
	}

	if n > 0 {
		init()
		pre()
		last()
	} else {
		g[0] = 1
	}

	for i := 0; i <= m; i++ {
		f[i] = 0
	}
	for i := 1; i <= k; i++ {
		for j := 0; j < m; j++ {
			f[(j+i)%m] = (f[(j+i)%m] + g[j]) % mod
		}
		for j := 0; j < m; j++ {
			g[j] = (f[j] + g[j]) % mod
			f[j] = 0
		}
	}

	return g[:m]
}
