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
	// 多个数最大公约数
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}
	// 多个数最小公倍数
	lcm := func(a, b int) int {
		return a / gcd(a, b) * b
	}

	var n int
	Fscan(in, &n)
	ans := 1
	for i := 2; i <= n; i++ {
		ans = lcm(ans, i)
	}
	Fprintln(out, ans+1)
}

func main() { run(os.Stdin, os.Stdout) }
