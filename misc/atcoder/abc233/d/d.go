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

	var n, k, v, sum, ans int
	Fscan(in, &n, &k)
	h := map[int]int{0: 1}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		sum += v
		ans += h[sum-k]
		h[sum]++
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
