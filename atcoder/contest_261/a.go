package contest_261

import (
	"fmt"
	"io"
)

func A(in io.Reader, out io.Writer) {

	var l1, l2 int
	var l3, l4 int
	fmt.Fscan(in, &l1, &l2, &l3, &l4)

	l := max(l1, l3)
	r := min(l2, l4)

	if r-l < 0 {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, r-l)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
