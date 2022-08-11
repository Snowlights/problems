package _700

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
)

func CF1704A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var m, n int
		fmt.Fscan(r, &n, &m)
		var a, b []byte
		fmt.Fscan(r, &a, &b)

		if m == n {
			if bytes.Compare(a, b) == 0 {
				fmt.Fprintln(w, "YES")
			} else {
				fmt.Fprintln(w, "NO")
			}
			continue
		}

		if bytes.Compare(a[n-m+1:], b[1:]) == 0 && bytes.Count(a[:n-m+1], b[:1]) >= 1 {
			fmt.Fprintln(w, "YES")
			continue
		}
		fmt.Fprintln(w, "NO")
	}
}

func CF1704B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, x int64
		fmt.Fscan(r, &n, &x)
		var val, ans int64
		maxV, minV := int64(math.MinInt64), int64(math.MaxInt64)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &val)
			maxV = max(maxV, val)
			minV = min(minV, val)
			if maxV-minV > 2*x {
				ans++
				maxV = val
				minV = val
			}
		}
		fmt.Fprintln(w, ans)
	}
}
