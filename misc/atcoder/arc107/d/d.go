package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var n, k int
	Fscan(in, &n, &k)
	f := []int{1}
	for i := 1; i <= n; i++ {
		f = slices.Insert(f, 0, 0)
		for j := i / 2; j > 0; j-- {
			f[j] = (f[j] + f[j*2]) % mod
		}
	}
	Fprint(out, f[k])

}

func main() { run(os.Stdin, os.Stdout) }
