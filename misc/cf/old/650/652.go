package _50

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF652A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var h1, h2, a, b int64
	fmt.Fscan(r, &h1, &h2, &a, &b)
	if h1+8*a >= h2 {
		fmt.Fprintln(w, 0)
	} else if b >= a {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, ((h2-h1-(8*a))+(12*(a-b))-1)/(12*(a-b)))
	}

}

func CF652B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	sort.Ints(arr)
	start, end := int64(0), n-1
	for start <= end {
		if start == end {
			fmt.Fprint(w, arr[start])
		} else {
			fmt.Fprint(w, arr[start], " ", arr[end], " ")
		}
		start++
		end--
	}

}

func CF652C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, ans int64
	p, rightP := [3e5 + 10]int64{}, [3e5 + 10]int64{}
	var left, right int64
	fmt.Fscan(r, &n, &m)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &left)
		p[left] = i
	}
	for i := int64(0); i < m; i++ {
		fmt.Fscan(r, &left, &right)
		left, right = p[left], p[right]
		if left > right {
			left, right = right, left
		}
		rightP[right] = max(rightP[right], left+1)
	}
	left = 1
	for i := int64(1); i <= n; i++ {
		left = max(left, rightP[i])
		ans += i - left + 1
	}
	fmt.Fprintln(w, ans)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
