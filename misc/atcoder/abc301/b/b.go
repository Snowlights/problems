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

	var n int
	Fscan(in, &n)
	a, ans := make([]int, n), make([]int, 0, n*2)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans = append(ans, a[0])
	for i := 1; i < n; i++ {
		val := abs(a[i] - a[i-1])
		if val == 1 {
			ans = append(ans, a[i])
			continue
		}
		c, d := a[i], a[i-1]
		if c > d {
			for x := d + 1; x < c; x++ {
				ans = append(ans, x)
			}
		} else {
			for x := d - 1; x > c; x-- {
				ans = append(ans, x)
			}
		}
		ans = append(ans, a[i])
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { run(os.Stdin, os.Stdout) }
