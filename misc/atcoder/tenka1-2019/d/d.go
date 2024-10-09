package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	const eof = 0
	out := bufio.NewWriter(_w)
	defer out.Flush()

	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB
	// 读一个字符
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	// 读一个整数，支持负数
	r := func() (x int) {
		b := rc()
		neg := false
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	var n, v, s, ans int
	mod, pow := 998244353, 1
	f := make([]int, 300*300+1)
	g := make([]int, 300*300+1)
	f[0], g[0] = 3, 3
	for n = r(); n > 0; n-- {
		v = r()
		s, pow = s+v, pow*3%mod
		for i := s; i >= 0; i-- {
			f[i] = (f[i]*2 + f[max(i-v, 0)]) % mod
			if i >= v {
				g[i] = (g[i] + g[i-v]) % mod
			}
		}
	}
	if s%2 == 0 {
		ans = pow - f[(s+1)/2] + g[s/2]
	} else {
		ans = pow - f[(s+1)/2]
	}
	Fprintln(out, (ans%mod+mod)%mod)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
