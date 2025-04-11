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

	const mod = 998244353
	var q, k, v int
	var op string
	Fscan(in, &q, &k)
	f := make([]int, k+1)
	f[0] = 1
	for ; q > 0; q-- {
		Fscan(in, &op, &v)
		if op == "+" {
			for i := k; i >= v; i-- {
				f[i] = (f[i] + f[i-v]) % mod
			}
		} else {
			for i := v; i <= k; i++ {
				f[i] = (f[i] - f[i-v] + mod) % mod
			}
		}
		Fprintln(out, f[k])
	}

}

func main() { run(os.Stdin, os.Stdout) }
