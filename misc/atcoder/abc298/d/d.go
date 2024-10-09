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
		res := int(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	inv10 := pow(10, mod-2)
	cur, pow10, s := 1, 1, []int{1}

	var q, op, v int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op)
		switch op {
		case 1:
			Fscan(in, &v)
			cur = (cur*10 + v) % mod
			pow10 = pow10 * 10 % mod
			s = append(s, v)
		case 2:
			cur = (cur - s[0]*pow10) % mod
			pow10 = pow10 * inv10 % mod
			s = s[1:]
		case 3:
			Fprintln(out, (cur+mod)%mod)
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
