package _100

import (
	"bufio"
	"fmt"
	"io"
)

func CF1178E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	cal := func(a, b, c, d byte) byte {
		if a == c || a == d {
			return a
		}
		if b == c || b == d {
			return b
		}
		return ' '
	}
	reserve := func(a []byte) []byte {
		b := make([]byte, len(a))
		copy(b, a)
		s, e := 0, len(b)-1
		for s < e {
			b[s], b[e] = b[e], b[s]
			s++
			e--
		}
		return b
	}

	ans := []byte{}
	for len(s) >= 4 {
		ans = append(ans, cal(s[0], s[1], s[len(s)-2], s[len(s)-1]))
		s = s[2 : len(s)-2]
	}
	if len(s) > 0 {
		fmt.Fprintln(w, string(ans)+string(s[0])+string(reserve(ans)))
	} else {
		fmt.Fprintln(w, string(ans)+string(reserve(ans)))
	}
}
