package _50

import (
	"bufio"
	"fmt"
	"io"
)

func CF482A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int64
	fmt.Fscan(r, &n, &k)
	left, right := 1, k+1
	for i := int64(1); i <= k+1; i++ {
		if i&1 == 1 {
			fmt.Fprint(w, left, " ")
			left++
		} else {
			fmt.Fprint(w, right, " ")
			right--
		}
	}
	for i := k + 2; i <= n; i++ {
		fmt.Fprint(w, i, " ")
	}

}

func CF482B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	qs := make([]struct{ l, r, q int }, m)
	for i := range qs {
		fmt.Fscan(r, &qs[i].l, &qs[i].r, &qs[i].q)
		qs[i].l--
	}

	ans := make([]int, n)
	for mask := 1; mask < 1<<30; mask <<= 1 {
		d := make([]int, n+1)
		for _, q := range qs {
			if q.q&mask > 0 {
				d[q.l]++
				d[q.r]--
			}
		}
		sum := make([]int, n+1)
		for i, v := range d[:n] {
			d[i+1] += v
			sum[i+1] = sum[i]
			if v > 0 {
				ans[i] |= mask
			} else {
				sum[i+1]++
			}
		}
		for _, q := range qs {
			if q.q&mask == 0 && sum[q.l] == sum[q.r] {
				fmt.Fprint(w, "NO")
				return
			}
		}
	}
	fmt.Fprintln(w, "YES")
	for _, v := range ans {
		fmt.Fprint(w, v, " ")
	}

}
