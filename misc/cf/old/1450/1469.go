package _450

import (
	"bufio"
	"fmt"
	"io"
)

func CF1469A(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	for fmt.Fscan(_r, &T); T > 0; T-- {
		var s string
		fmt.Fscan(_r, &s)
		if s[0] == ')' {
			fmt.Fprintln(_w, "NO")
			continue
		}
		cnt := 0
		for _, v := range s {
			if v == '(' || v == '?' {
				cnt++
			} else if v == ')' {
				cnt--
			}
		}
		if cnt%2 == 0 {
			fmt.Fprintln(_w, "YES")
		} else {
			fmt.Fprintln(_w, "NO")
		}
	}

}

func CF1469B(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()
	var T int64

	for fmt.Fscan(_r, &T); T > 0; T-- {
		aM, aVal, bM, bVal := (0), (0), (0), (0)
		var a, b int
		fmt.Fscan(_r, &a)
		for i := (0); i < a; i++ {
			fmt.Fscan(_r, &b)
			aVal += b
			aM = max(aM, aVal)
		}
		fmt.Fscan(_r, &b)
		for i := 0; i < b; i++ {
			fmt.Fscan(_r, &a)
			bVal += a
			bM = max(bM, bVal)
		}
		fmt.Fprintln(_w, aM+bM)
	}

}

func CF1469C(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T, n, k, h, l int
	for fmt.Fscan(_r, &T); T > 0; T-- {
		fmt.Fscan(_r, &n, &k, &l)
		k--
		s := "YES"
		r := l
		for ; n > 2; n-- {
			fmt.Fscan(_r, &h)
			l = max(l-k, h)
			r = min(r, h) + k
			if l > r {
				s = "NO"
			}
		}
		fmt.Fscan(_r, &h)
		if l-k > h || r+k < h {
			s = "NO"
		}
		fmt.Fprintln(_w, s)
	}

}

func CF1469D(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	x, y := make([]int64, 2e5+10), make([]int64, 2e5+10)
	for fmt.Fscan(_r, &T); T > 0; T-- {
		var n int64
		fmt.Fscan(_r, &n)
		k, ans := n, 0
		for i := n - 1; i > 1; i-- {
			for (k-1)/i+1 >= i {
				x[ans] = n
				y[ans] = i
				ans++
				k = (k-1)/i + 1
			}
			x[ans] = i
			y[ans] = n
			ans++
		}
		fmt.Fprintln(_w, ans)
		for i := 0; i < ans; i++ {
			fmt.Fprintln(_w, x[i], " ", y[i])
		}
	}

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
