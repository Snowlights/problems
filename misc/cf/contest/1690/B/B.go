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

	var T, n int
loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a, b, diff := make([]int, n), make([]int, n), math.MaxInt32
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := range b {
			Fscan(in, &b[i])
			if b[i] != 0 && a[i]-b[i] < diff {
				diff = a[i] - b[i]
			}
		}
		for i := range a {
			if b[i] > a[i] || a[i]-b[i] > diff {
				Fprintln(out, "NO")
				continue loop
			}
		}
		Fprintln(out, "YES")
	}

}

func main() { run(os.Stdin, os.Stdout) }
