package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const mod int = 1e9 + 7

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, xor int
	f := [1 << 20]struct{ s0, sx, pre0 int }{}
	for i := range f {
		f[i].s0 = 1
	}
	cnt0 := 1
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		xor ^= v
		if xor == 0 {
			cnt0++
		} else {
			t := &f[xor]
			t.s0 = (t.s0 + t.sx*(cnt0-t.pre0)) % mod
			t.sx = (t.sx + t.s0) % mod
			t.pre0 = cnt0
		}
	}
	if xor > 0 {
		Fprintln(out, f[xor].s0)
	} else {
		ans := pow(2, cnt0-2)
		for _, t := range f {
			ans += t.sx
		}
		Fprintln(out, ans%mod)
	}
}

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}

func main() { run(os.Stdin, os.Stdout) }
