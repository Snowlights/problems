package _100

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1197C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	tmp := make([]int, 0)
	for i := 1; i < n; i++ {
		tmp = append(tmp, arr[i-1]-arr[i])
	}
	sort.Ints(tmp)
	res := arr[n-1] - arr[0]
	for _, v := range tmp[:k-1] {
		res += v
	}
	fmt.Fprintln(w, res)

}

func CF1197D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	var k, v, s, ans int64
	fmt.Fscan(r, &n, &m, &k)
	minS := make([]int64, m)
	for i := 1; i < m; i++ {
		minS[i] = 1e18
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v)
		s += v*int64(m) - k
		for d := 0; d < m; d++ {
			ans = max(ans, s-minS[(i+d)%m]-k*int64(d))
		}
		minS[i%m] = min(minS[i%m], s)
	}
	fmt.Fprint(w, ans/int64(m))
}
