package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	pow := func(x, n int) int {
		//x %= mod
		res := int(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	solve := func() {
		var n, v, s int
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			s = (s + v) % mod
		}
		Fprintln(out, s*pow(2, n-1)%mod)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
