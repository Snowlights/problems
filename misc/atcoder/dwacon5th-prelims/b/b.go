package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
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

	var n, k, ans, sum int
	n, k = r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
		sum += a[i]
	}
	for i := 1 << (bits.Len(uint(sum)) - 1); i > 0; i >>= 1 {
		cnt, s := 0, 0
		for j := range a {
			s = 0
			for _, v := range a[j:] {
				s += v
				if s&i > 0 && s&ans == ans {
					cnt++
				}
			}
		}
		if cnt >= k {
			ans |= i
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
