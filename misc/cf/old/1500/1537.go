package _500

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF1537E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int
	var s string
	fmt.Fscan(r, &n, &k, &s)
	sz := 1
	for i := range s {
		if s[i] > s[i%sz] {
			break
		}
		if s[i] < s[i%sz] {
			sz = i + 1
		}
	}
	fmt.Fprint(w, strings.Repeat(s[:sz], k/sz+1)[:k])
}
