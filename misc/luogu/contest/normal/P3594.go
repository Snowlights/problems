//go:build main
// +build main

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

	n, lim, d := r(), r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}

	ans := d
	sd := 0
	for _, v := range a[:d] {
		sd += v
	}
	type pair struct{ l, s int }
	q := []pair{{0, sd}}
	s := sd
	left := 0
	for i := d; i < n; i++ {
		sd += a[i] - a[i-d]
		for len(q) > 0 && sd >= q[len(q)-1].s {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i - d + 1, sd})

		s += a[i]
		for s-q[0].s > lim {
			s -= a[left]
			left++
			for q[0].l < left {
				q = q[1:]
			}
		}

		ans = max(ans, i-left+1)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
