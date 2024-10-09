package _600

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func CF1647A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)

		if n <= 2 {
			fmt.Fprintln(w, n)
			continue
		}

		body := make([]byte, 0)
		left, right := n/3, n%3
		if left > 0 {
			body = append(body, bytes.Repeat([]byte("21"), int(left))...)
		}
		switch right {
		case 1:
			body = append([]byte{'1'}, body...)
		case 2:
			body = append(body, []byte{'2'}...)
		}
		fmt.Fprintln(w, string(body))
	}

}
