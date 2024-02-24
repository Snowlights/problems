//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF480C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// 从数轴上a出发移动k次，输出不同方案数
	// 从x到y 限制
	// 1 <= y <= n
	// y != x
	// y != b
	// |x-y| < |x-b|
	var n, a, b, k int
	mod := int(1e9 + 7)
	Fscan(in, &n, &a, &b, &k)
	if a > b {
		a = n + 1 - a
		b = n + 1 - b
	}
	f := make([]int, b)
	f[a] = 1
	s := make([]int, b+1)
	for ; k > 0; k-- {
		for i, v := range f {
			s[i+1] = (s[i] + v) % mod
		}
		for y := 1; y < b; y++ {
			f[y] = (s[y+(b-y-1)/2+1] - f[y]) % mod
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
	}
	Fprint(out, (ans%mod+mod)%mod)
}

func main() { CF480C(os.Stdin, os.Stdout) }
