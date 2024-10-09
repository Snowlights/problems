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

	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	solve := func() {
		var n, v, ans int64
		Fscan(in, &n)
		for i := int64(0); i < n; i++ {
			Fscan(in, &v)
			ans |= v
		}
		Fprintln(out, ans*pow(2, n-1)%mod)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
