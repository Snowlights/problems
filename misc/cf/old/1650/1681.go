package _650

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func CF1691A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var a, aMax, b, bMax int64
		var val int64
		fmt.Fscan(r, &a)
		for i := int64(0); i < a; i++ {
			fmt.Fscan(r, &val)
			aMax = max(val, aMax)
		}
		fmt.Fscan(r, &b)
		for i := int64(0); i < b; i++ {
			fmt.Fscan(r, &val)
			bMax = max(val, bMax)
		}
		if aMax > bMax {
			fmt.Fprintln(w, "Alice")
			fmt.Fprintln(w, "Alice")
		} else if aMax == bMax {
			fmt.Fprintln(w, "Alice")
			fmt.Fprintln(w, "Bob")
		} else {
			fmt.Fprintln(w, "Bob")
			fmt.Fprintln(w, "Bob")
		}
	}

}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func CF1691B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, m int64
		fmt.Fscan(r, &n)
		a := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		val, total := 0, 0
		fmt.Fscan(r, &m)
		for i := int64(0); i < m; i++ {
			fmt.Fscan(r, &val)
			total += val
			total %= len(a)
		}
		fmt.Fprintln(w, a[total%len(a)])
	}
}

func CF1691C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		a, b := make([]int64, n), make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &b[i])
		}

		ans, ok := CF1691CS(a, b, n)
		if !ok {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintln(w, len(ans))
			for _, v := range ans {
				fmt.Fprintln(w, v[0]+1, v[1]+1)
			}
		}

	}
}

func CF1691CS(a, b []int64, n int64) ([][]int64, bool) {
	ans := make([][]int64, 0)

	for i := int64(0); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] >= a[j] && b[i] >= b[j] {
				a[i], a[j] = a[j], a[i]
				b[i], b[j] = b[j], b[i]
				ans = append(ans, []int64{i, j})
			} else if a[i] <= a[j] && b[i] <= b[j] {
				continue
			} else {
				return nil, false
			}
		}
	}

	return ans, true
}

func CF1691D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, x int64
	fmt.Fscan(r, &n, &x)

	mp := make(map[int64]int64)
	q := []int64{x}
	for len(q) > 0 {
		val := q[0]
		q = q[1:]

		s := strconv.FormatInt(val, 10)
		if int64(len(s)) >= n {
			fmt.Fprintln(w, mp[val])
			return
		}

		for _, v := range s {
			v := int64(v - '0')
			if v <= 1 {
				continue
			}
			_, ok := mp[val*v]
			if ok {
				continue
			}
			mp[val*v] = mp[val] + 1
			q = append(q, val*v)
		}
	}

	fmt.Fprintln(w, -1)
}
