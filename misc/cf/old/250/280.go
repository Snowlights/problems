package _50

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func CF280A(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var w, h, a float64
	fmt.Fscan(_r, &w, &h, &a)
	if a == 0 || a == 180 {
		fmt.Fprintln(_w, fmt.Sprintf("%.8f", w*h))
		return
	}

	if a > 90 {
		a = 180 - a
	}
	if w < h {
		w, h = h, w
	}
	a = a * math.Pi / 180
	k := 1/math.Sin(a) + 1/math.Tan(a)
	if h*k < w {
		fmt.Fprintln(_w, fmt.Sprintf("%.8f", h*h/math.Sin(a)))
		return
	}
	x := (k*w - h) / (k*k - 1)
	y := (k*h - w) / (k*k - 1)
	fmt.Fprintln(_w, fmt.Sprintf("%.8f", w*h-(x*x+y*y)/math.Tan(a)))

}

func CF280B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, v, ans int
	stack := make([]int, 0)
	for fmt.Fscan(r, &n); n > 0; n-- {
		fmt.Fscan(r, &v)
		for len(stack) > 0 && stack[len(stack)-1] < v {
			ans = max(ans, stack[len(stack)-1]^v)
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans = max(ans, stack[len(stack)-1]^v)
		}
		stack = append(stack, v)
	}
	fmt.Fprintln(w, ans)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
