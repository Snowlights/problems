package _600

import (
	"bufio"
	"fmt"
	"io"
)

func CF1611F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n, s int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(in, &n, &s)
		a := make([]int64, n+1)
		for i := int64(1); i <= n; i++ {
			fmt.Fscan(in, &a[i])
			a[i] += a[i-1]
		}
		ll, rr := int64(-1), int64(-1)
		for left, right := int64(0), int64(0); right < n; left++ {
			for right < n && a[right+1]-a[left] >= -s {
				right++
			}
			if right-left > int64(rr-ll) {
				ll, rr = left, right
			}
		}
		if ll < 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, ll+1, rr)
		}
	}

}
