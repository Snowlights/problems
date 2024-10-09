package _500

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1519A(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	for fmt.Fscan(_r, &T); T > 0; T-- {
		var r, b, d int64
		fmt.Fscan(_r, &r, &b, &d)
		m := min(r, b)
		switch m {
		case r:
			if r*(d+1) < b {
				fmt.Fprintln(_w, "NO")
			} else {
				fmt.Fprintln(_w, "YES")
			}
		case b:
			if b*(d+1) < r {
				fmt.Fprintln(_w, "NO")
			} else {
				fmt.Fprintln(_w, "YES")
			}
		}
	}

}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func CF1519B(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	for fmt.Fscan(_r, &T); T > 0; T-- {
		var n, m, k int64
		fmt.Fscan(_r, &n, &m, &k)
		if k == n*m-1 {
			fmt.Fprintln(_w, "YES")

		} else {
			fmt.Fprintln(_w, "NO")
		}
	}

}

func CF1519C(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	sq := make([]int64, 2e5+10)
	s, d := make([]int64, 2e5+10), make([]int64, 2e5+10)
	sToDList := make(map[int64][]int64)
	for fmt.Fscan(_r, &T); T > 0; T-- {
		var n int64
		fmt.Fscan(_r, &n)
		sq[n] = 0
		sToDList[n] = nil
		for i := int64(0); i < n; i++ {
			fmt.Fscan(_r, &s[i])
			sq[i] = 0
			sToDList[i] = nil
		}
		for i := int64(0); i < n; i++ {
			fmt.Fscan(_r, &d[i])
			sToDList[s[i]] = append(sToDList[s[i]], int64(d[i]))
		}

		for i := int64(1); i <= n; i++ {
			v := sToDList[i]
			sz := len(v)
			if sz == 0 {
				continue
			}
			sort.Slice(v, func(i, j int) bool { return v[i] > v[j] })
			prefix := make([]int64, len(v)+1)
			for j := 0; j < sz; j++ {
				prefix[j+1] = prefix[j] + int64(v[j])
			}
			for j := 1; j <= sz; j++ {
				sq[j] += prefix[len(v)/j*j]
			}
		}

		for k := int64(1); k <= n; k++ {
			fmt.Fprint(_w, sq[k], " ")
		}
		fmt.Fprintln(_w)
	}

}

func CF1519D(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n int64
	fmt.Fscan(_r, &n)
	a, b := make([]int64, n), make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(_r, &a[i])
	}
	sum := int64(0)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(_r, &b[i])
		sum += a[i] * b[i]
	}

	ans := sum
	for i := int64(0); i < n; i++ {
		cur := sum
		for l, r := i-1, i+1; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
		cur = sum
		for l, r := i-1, i; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
	}
	fmt.Fprintln(_w, ans)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
