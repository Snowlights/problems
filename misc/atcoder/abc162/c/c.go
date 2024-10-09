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

	var k, ans int
	Fscan(in, &k)
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= k; j++ {
			for l := 1; l <= k; l++ {
				ans += gcd(gcd(i, j), l)
			}
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
