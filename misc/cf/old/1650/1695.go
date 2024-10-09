package _650

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func CF1695A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, m int64
		fmt.Fscan(r, &n, &m)
		var i, j int64
		var iIdx, jIdx, val int64
		maxVal := int64(math.MinInt64)
		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				fmt.Fscan(r, &val)
				if val > maxVal {
					iIdx, jIdx, maxVal = i, j, val
				}
			}
		}

		ans := (iIdx + 1) * (jIdx + 1)
		ans = max(ans, (iIdx+1)*(m-jIdx))
		ans = max(ans, (n-iIdx)*(m-jIdx))
		ans = max(ans, (n-iIdx)*(jIdx+1))
		fmt.Fprintln(w, ans)
	}

}
func CF1695B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		arr := make([]int64, n)
		small := int64(0)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &arr[i])
			if arr[i] < arr[small] {
				small = i
			}
		}

		if n%2 == 1 || small%2 == 1 {
			fmt.Fprintln(w, "Mike")
		} else {
			fmt.Fprintln(w, "Joe")
		}

	}

}

func CF1695C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, m int64
		fmt.Fscan(r, &n, &m)
		arr := make([][]int64, n)
		mn, mx := make([][]int64, n), make([][]int64, n)
		for i := int64(0); i < n; i++ {
			arr[i] = make([]int64, m)
			mn[i], mx[i] = make([]int64, m), make([]int64, m)
			for j := int64(0); j < m; j++ {
				fmt.Fscan(r, &arr[i][j])
			}
		}
		mn[0][0], mx[0][0] = arr[0][0], arr[0][0]
		for i := int64(1); i < n; i++ {
			mx[i][0], mn[i][0] = mx[i-1][0]+arr[i][0], mx[i-1][0]+arr[i][0]
		}
		for j := int64(1); j < m; j++ {
			mx[0][j], mn[0][j] = mx[0][j-1]+arr[0][j], mx[0][j-1]+arr[0][j]
		}

		for i := int64(1); i < n; i++ {
			for j := int64(0); j < m; j++ {
				mx[i][j] = max(mx[i-1][j], mx[i][j-1]) + arr[i][j]
				mn[i][j] = min(mn[i-1][j], mn[i][j-1]) + arr[i][j]
			}
		}

		if mx[n-1][m-1]%2 == 1 || mn[n-1][m-1] > 0 || mx[n-1][m-1] < 0 {
			fmt.Fprintln(w, "NO")
		} else {
			fmt.Fprintln(w, "YES")
		}

	}

}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
