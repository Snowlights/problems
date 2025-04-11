package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const mod = 1_000_000_007
const mx = 200_000

var f [mx]int
var invF [mx]int

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

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

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k int
	Fscan(in, &n, &m, &k)
	Fprint(out, (m*n*(m*(n*n-1)+n*(m*m-1)))/6%mod*comb(m*n-2, k-2)%mod)

}

func main() { run(os.Stdin, os.Stdout) }
