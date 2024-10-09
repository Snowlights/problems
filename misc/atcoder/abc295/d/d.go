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

	var mask, ans int
	var s string
	Fscan(in, &s)
	h := map[int]int{0: 1}
	for _, v := range s {
		mask ^= 1 << int(v-'0')
		ans += h[mask]
		h[mask]++
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
