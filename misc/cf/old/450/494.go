package _50

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF494A(in io.Reader, out io.Writer) {
	var s string
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	fmt.Fscan(r, &s)
	cnt, d, minD := 0, 0, len(s)
	for _, b := range s {
		if b == '(' {
			d++
			continue
		}
		d--
		if d < 0 {
			fmt.Fprintln(w, -1)
			return
		}
		if b == '#' {
			cnt++
			minD = d
		} else if d < minD {
			minD = d
		}
	}
	if minD < d {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, strings.Repeat("1\n", cnt-1), d+1)
	}
}
