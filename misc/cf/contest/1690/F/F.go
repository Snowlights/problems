package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"math"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	const eof = 0
	out := bufio.NewWriter(_w)
	defer out.Flush()
	// 多个数最大公约数
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	// 多个数最小公倍数
	lcm := func(a, b int) int {
		return a / gcd(a, b) * b
	}

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
	// 读一个长度为 n 的仅包含小写字母的字符串
	rsn := func(n int) []byte {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() { // 'A' 'Z'
		}
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}

	for t := r(); t > 0; t-- {
		n, res := r(), 1
		s := rsn(n)
		p, vis := make([]int, n), make([]bool, n)
		for i := 0; i < n; i++ {
			p[i] = r()
			p[i]--
		}
		circle := func(x int) (res []byte) {
			res, vis[x] = append(res, s[x]), true
			for i := p[x]; i != x; i = p[i] {
				res, vis[i] = append(res, s[i]), true
			}
			return res
		}
		minReStr := func(str []byte) int {
			n, res := len(str), math.MaxInt32
			for i := 1; i*i <= n; i++ {
				if n%i == 0 {
					t := n / i
					if bytes.Equal(bytes.Repeat(str[:i], t), str) && res > i {
						res = i
					}
					if bytes.Equal(bytes.Repeat(str[:t], i), str) && res > t {
						res = t
					}
				}
			}
			return res
		}
		for i := 0; i < n; i++ {
			if !vis[i] {
				res = lcm(res, minReStr(circle(i)))
			}
		}
		Fprintln(out, res)
	}

}

func main() { run(os.Stdin, os.Stdout) }
