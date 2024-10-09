package _000

import (
	"bufio"
	"fmt"
	"io"
)

func CF1043A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int64, n)

	k := int64(-1)
	noted := int64(0)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
		k = max(k, int64(arr[i]))
		noted += int64(arr[i])
	}

	for {
		if k*n > 2*noted {
			break
		}
		k++
	}
	fmt.Fprintln(w, k)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func CF1043B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int64, n+1)
	barr := make([]int64, n)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &arr[i])
		barr[i-1] = arr[i] - arr[i-1]
	}

	res := make([]int64, 0)
	for i := int64(1); i <= n; i++ {
		found := true
		for j := int64(0); j < n; j++ {
			if arr[j+1] != barr[j%i]+arr[j] {
				found = false
				break
			}
		}
		if found {
			res = append(res, i)
		}
	}
	fmt.Fprintln(w, len(res))
	for _, v := range res {
		fmt.Fprint(w, v, " ")
	}

}

func CF1043C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)
	s += "b"
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			fmt.Fprint(w, "1 ")
		} else {
			fmt.Fprint(w, "0 ")
		}
	}
}

func CF1043D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int64
	fmt.Fscan(in, &n, &m)

	a := make([][]int64, m)
	pre := make([]int64, n+1)
	for i := int64(0); i < m; i++ {
		a[i] = make([]int64, n+1)
		for j := int64(1); j <= n; j++ {
			fmt.Fscan(r, &pre[j])
			a[i][pre[j]] = pre[j-1]
		}
	}

	ans, cnt := int64(0), int64(0)
	for i := int64(2); i <= n; i++ {
		j := int64(0)
		for j = int64(1); j < m; j++ {
			if a[0][pre[i]] != a[j][pre[i]] {
				break
			}
		}

		if j == m {
			cnt++
			ans += cnt
		} else {
			cnt = 0
		}
	}

	fmt.Fprintln(w, ans+n)
}
