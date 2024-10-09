package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, ans int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		Fscan(in, &s)
		a := make([]int, n+1)
		for i, v := range s {
			a[i+1] = a[i]
			if v == 'W' {
				a[i+1]++
			}
		}
		ans = math.MaxInt32
		for i := k; i <= n; i++ {
			if a[i]-a[i-k] < ans {
				ans = a[i] - a[i-k]
			}
		}
		Fprintln(out, ans)
	}

}

func main() { run(os.Stdin, os.Stdout) }
