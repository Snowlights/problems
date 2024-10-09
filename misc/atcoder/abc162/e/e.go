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
	var n, k, ans int
	Fscan(in, &n, &k)
	cnt := map[int]int{}
	for i := k; i > 0; i-- {
		c := pow(k/i, n)
		for j := 2 * i; j <= k; j += i {
			c -= cnt[j] % mod
		}
		cnt[i] = c % mod
		ans = (ans + i*cnt[i]) % mod
	}
	Fprintln(out, (ans%mod+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
