package _700

import (
	"bufio"
	"fmt"
	"io"
)

func CF1711A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		for i := int64(1); i < n; i++ {
			fmt.Fprint(w, i+1, " ")
		}
		fmt.Fprintln(w, 1)
	}
}
