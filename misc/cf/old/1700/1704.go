package _700

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
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

func CF1704C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, m int64
		fmt.Fscan(r, &n, &m)
		arr, ll := make([]int, m), make([]int, 0)
		for i := int64(0); i < m; i++ {
			fmt.Fscan(r, &arr[i])
		}
		sort.Ints(arr)
		for i := int64(0); i < m; i++ {
			if i == 0 {
				ll = append(ll, arr[i]+int(n)-arr[m-1]-1)
				continue
			}
			ll = append(ll, arr[i]-arr[i-1]-1)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(ll)))

		day, ans := 0, 0
		for _, v := range ll {
			if v > day*2 {
				if v-1-day*2 >= 2 {
					ans += v - 1 - day*2
					day += 2
				} else {
					ans += 1
					day += 1
				}
			}
		}
		fmt.Fprintln(w, int(n)-ans)
	}

}
