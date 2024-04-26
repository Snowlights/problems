//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF1886C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, pos int
	var s []byte
	f := func() byte {
		pos--
		n := len(s)
		if pos < n {
			return s[pos]
		}
		m := n*2 + 1
		k := int((float64(m) - math.Sqrt(float64(m*m-pos*8))) / 2)
		pos -= (n*2 - k + 1) * k / 2
		st := []byte{}
		s = append(s, 0)
		for i := 0; ; i++ {
			for len(st) > 0 && s[i] < st[len(st)-1] {
				st = st[:len(st)-1]
				k--
				if k == 0 {
					if pos < len(st) {
						return st[pos]
					}
					return s[i+pos-len(st)]
				}
			}
			st = append(st, s[i])
		}
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &pos)
		Fprintf(out, "%c", f())
	}

}

func main() { CF1886C(os.Stdin, os.Stdout) }
