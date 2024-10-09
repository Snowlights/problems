package _200

import (
	"bufio"
	"fmt"
	"io"
)

func CF1217A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var s, i, exp int64
		fmt.Fscan(r, &s, &i, &exp)

		fmt.Fprintln(w, max(0, exp-max(0, (exp+i-s+2)/2)+1))
	}
}
func CF1217B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, x int64
		fmt.Fscan(r, &n, &x)
		var h, d int64
		var m, g int64
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &d, &h)
			m = max(d, m)
			g = max(g, d-h)
		}
		if m >= x {
			fmt.Fprintln(w, 1)
		} else if g <= 0 {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintln(w, (x-m+g-1)/g+1)
		}
	}
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
func CF1217C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	var s string
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &s)
		ans, cnt0 := 0, 0
		for i, b := range s {
			if b == '0' {
				cnt0++
				continue
			}
			v := 0
			for j, b := range s[i:] {
				v = v<<1 | int(b&1)
				if v > cnt0+j+1 {
					break
				}
				// 由于 v 最高位为 1，v >= 2^j >= j+1 是恒成立的
				// 因此当 j+1 <= v <= j+1+cnt0 时，一定能添加足够的前导零组成长度为 v 的二进制
				ans++
			}
			cnt0 = 0
		}
		fmt.Fprintln(w, ans)
	}
}
