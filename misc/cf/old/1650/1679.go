package _650

import (
	"bufio"
	"fmt"
	"io"
)

func CF1670A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	const (
		four = 4
		six  = 6
	)

	findMax := func(n int64) int64 {
		left := n % four
		if left == 0 {
			return n / four
		}
		for i := n / four; i >= 0; i-- {
			sLeft := (n - i*4) % six
			if sLeft == 0 {
				return i + (n-i*4)/six
			}
		}
		return -1
	}

	findMin := func(n int64) int64 {
		left := n % six
		if left == 0 {
			return n / six
		}
		for i := n / six; i >= 0; i-- {
			sLeft := (n - i*6) % four
			if sLeft == 0 {
				return i + (n-i*6)/four
			}
		}
		return -1
	}

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		if n%2 == 1 {
			fmt.Fprintln(w, "-1")
			continue
		}

		f, s := findMax(n), findMin(n)
		if f == -1 || s == -1 {
			fmt.Fprintln(w, "-1")
			continue
		}
		fmt.Fprintln(w, f, s)
	}

}

func CF1670B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, q int64
	var total int64
	fmt.Fscan(r, &n, &q)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
		total += arr[i]
	}

	m := make(map[int64]int64)
	setAll, allVal := false, int64(0)
	for j := int64(0); j < q; j++ {
		var t, i, x int64
		fmt.Fscan(r, &t)

		switch t {
		case 1:
			fmt.Fscan(r, &i, &x)
			originalVal := arr[i-1]
			if setAll {
				originalVal = allVal
			}
			v, ok := m[i]
			if ok {
				originalVal = v
			}
			total += x - originalVal
			m[i] = x

		case 2:
			fmt.Fscan(r, &x)
			total = n * x
			m = make(map[int64]int64)
			setAll = true
			allVal = x
		}

		fmt.Fprintln(w, total)
	}

}

// 树状数组
func CF1670C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const N = 500000

	mx := [N + 1]int{}
	my := [N + 1]int{}
	xr := [N + 1]int{}
	yr := [N + 1]int{}
	var n, q int
	fmt.Fscan(r, &n, &q)

	addX := func(idx, v int) {
		for i := idx; i <= n; i = i | (i + 1) {
			mx[i] += v
		}
	}
	addY := func(idx, v int) {
		for i := idx; i <= n; i = i | (i + 1) {
			my[i] += v
		}
	}
	prefixX := func(idx int) int {
		var tot int
		for ; idx >= 0; idx = (idx & (idx + 1)) - 1 {
			tot += mx[idx]
		}
		return tot
	}
	prefixY := func(idx int) int {
		var tot int
		for ; idx >= 0; idx = (idx & (idx + 1)) - 1 {
			tot += my[idx]
		}
		return tot
	}

	for ; q > 0; q-- {
		var t int64
		fmt.Fscan(r, &t)
		switch t {
		case 1:
			var x, y int
			fmt.Fscan(r, &x, &y)
			if xr[x] == 0 {
				addX(x, 1)
			}
			if yr[y] == 0 {
				addY(y, 1)
			}

			xr[x]++
			yr[y]++
		case 2:
			var x, y int
			fmt.Fscan(r, &x, &y)
			xr[x]--
			yr[y]--
			if xr[x] == 0 {
				addX(x, -1)
			}
			if yr[y] == 0 {
				addY(y, -1)
			}
		case 3:
			var x1, y1, x2, y2 int
			fmt.Fscan(r, &x1, &y1, &x2, &y2)

			var f bool
			if prefixX(x2)-prefixX(x1-1) >= x2-x1+1 {
				f = true
			}

			if !f {
				if prefixY(y2)-prefixY(y1-1) >= y2-y1+1 {
					f = true
				}
			}
			if f {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}

	}

}
