//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF204A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	f := func(n int) int {
		if n < 10 {
			return n
		}
		res := 9 + n/10
		if int(strconv.Itoa(n)[0]-'0') > n%10 {
			res--
		}
		return res
	}
	var l, r int
	Fscan(in, &l, &r)
	Fprint(out, f(r)-f(l-1))

}

func main() { CF204A(os.Stdin, os.Stdout) }
