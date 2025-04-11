//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func CF1778C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s, &t)
		u := 0
		for _, b := range s {
			u |= 1 << (b - 'a')
		}
		ans := 0
		for sub, ok := u, true; ok; ok = sub != u {
			if bits.OnesCount(uint(sub)) <= k {
				res, cnt := 0, 0
				for i, b := range s {
					if sub>>(b-'a')&1 > 0 || b == t[i] {
						cnt++
						res += cnt
					} else {
						cnt = 0
					}
				}
				ans = max(ans, res)
			}
			sub = (sub - 1) & u
		}
		Fprintln(out, ans)
	}

}

func main() { CF1778C(os.Stdin, os.Stdout) }
