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

	const mod int = 998244353
	pow := func(x, n int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n int
	Fscan(in, &n)
	inv5, memo := pow(5, mod-2), make(map[int]int)
	var f func(x int) int
	f = func(x int) int {
		if x == 1 {
			return 1
		}
		val, ok := memo[x]
		if ok {
			return val
		}
		res := 0
		defer func() {
			memo[x] = res
		}()
		for i := 2; i <= 6; i++ {
			if x%i == 0 {
				res += f(x/i) % mod
			}
		}
		res = res * inv5 % mod
		return res
	}
	Fprintln(out, f(n))

}

func main() { run(os.Stdin, os.Stdout) }
