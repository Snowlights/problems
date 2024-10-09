package _00

import (
	"bufio"
	"fmt"
	"io"
)

func CF729D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, a, b, k int
	var s string
	fmt.Fscan(r, &n, &a, &b, &k, &s)
	pos := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			continue
		}
		st := i
		for i++; i < n && s[i] == '0'; i++ {
		}
		for st += b - 1; st < i; st += b {
			pos = append(pos, st)
		}
	}
	pos = pos[a-1:]
	fmt.Fprintln(out, len(pos))
	for _, v := range pos {
		fmt.Fprint(out, v+1, " ")
	}
}
