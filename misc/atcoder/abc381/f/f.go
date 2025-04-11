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

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ps := [20]int{}
	for i := range ps {
		ps[i] = n
	}
	nxt := make([][20]int, n+2)
	nxt[n+1] = ps
	nxt[n] = ps
	for i := n - 1; i >= 0; i-- {
		ps[a[i]-1] = i
		nxt[i] = ps
	}

	f := make([]int, 1<<20)
	for i := range f {
		f[i] = n
	}
	f[0] = -1
	for i, fi := range f {
		if fi == n {
			continue
		}
		ans = max(ans, bits.OnesCount(uint(i)))
		for j, lb := len(f)-1^i, 0; j > 0; j ^= lb {
			lb = j & -j
			v := bits.TrailingZeros(uint(lb))
			f[i|lb] = min(f[i|lb], nxt[nxt[fi+1][v]+1][v])
		}
	}
	Fprint(out, ans*2)

}

func main() { run(os.Stdin, os.Stdout) }
