package _050

import (
	"bufio"
	"fmt"
	"io"
)

func CF1060A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	bytes := make([]byte, n)
	fmt.Fscan(r, &bytes)
	h := make(map[byte]int)
	for _, v := range bytes {
		h[v]++
	}

	fmt.Fprintln(w, min(int64(n/11), int64(h['8'])))

}
func CF1060B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	ans := int64(0)
	for n > 9 {
		n -= 9
		ans += 9
		ans += n % 10
		n /= 10
	}

	fmt.Fprint(w, ans+n)
}

func CF1060C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, x, ans int64

	fmt.Fscan(r, &n, &m)
	f := func(n int64) []int64 {
		mins := make([]int64, n)
		for i := range mins {
			mins[i] = 1e9
		}
		a := make([]int64, n)
		for i := range a {
			fmt.Fscan(r, &a[i])
			for j, s := i, int64(0); j >= 0; j-- {
				s += a[j]
				mins[i-j] = min(mins[i-j], s)
			}
		}
		return mins
	}
	minA, minB := f(n), f(m)
	fmt.Fscan(r, &x)

	for i, v := range minA {
		for j, w := range minB {
			if int64((i+1)*(j+1)) > ans && int64(v)*int64(w) <= int64(x) {
				ans = int64((i + 1) * (j + 1))
			}
		}
	}
	fmt.Fprint(w, ans)
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
