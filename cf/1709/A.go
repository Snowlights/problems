package _709

import (
	"bufio"
	"fmt"
	"io"
)

func CF1709A(in io.Reader, out io.Writer) {

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		arr := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(in, &arr[j])
		}

		now := 0
		for k != 0 {
			k = arr[k-1]
			now++
		}
		if now == 3 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}

func CF1709B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int64
	fmt.Fscan(r, &n, &m)
	arr := make([]int64, n)
	prefix, sufFix := make([]int64, n), make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
		if i == 0 {
			continue
		}
		if arr[i] < arr[i-1] {
			prefix[i] = prefix[i-1] - arr[i] + arr[i-1]
			sufFix[i] = sufFix[i-1]
		} else {
			prefix[i] = prefix[i-1]
			sufFix[i] = sufFix[i-1] - arr[i-1] + arr[i]
		}
	}

	var start, end int64
	for i := int64(0); i < m; i++ {
		fmt.Fscan(r, &start, &end)
		if start < end {
			fmt.Fprintln(w, prefix[end-1]-prefix[start-1])
		} else {
			fmt.Fprintln(w, sufFix[start-1]-sufFix[end-1])
		}
	}
}
