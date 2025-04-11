package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 40
	var n, k, v, ans, pre int
	Fscan(in, &n, &k)
	cnt := [mx]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j := uint(v); j > 0; j &= j - 1 {
			cnt[bits.TrailingZeros(j)]++
		}
	}

	suf := [mx + 1]int{}
	for i, c := range cnt {
		suf[i+1] = suf[i] + max(c, n-c)<<i
	}

	k++
	for i := mx - 1; i >= 0; i-- {
		if k>>i&1 > 0 {
			ans = max(ans, pre+cnt[i]<<i+suf[i])
			pre += (n - cnt[i]) << i
		} else {
			pre += cnt[i] << i
		}
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
