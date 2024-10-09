package _700

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func CF1706A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, m int
		fmt.Fscan(r, &n, &m)
		s := bytes.Repeat([]byte("B"), m)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}
		for _, v := range arr {
			left, right := v-1, m-v
			if left > right {
				left, right = right, left
			}

			if s[left] == 'A' {
				s[right] = 'A'
			} else {
				s[left] = 'A'
			}

		}
		fmt.Fprintln(w, string(s))
	}
}

func CF1706B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int
		fmt.Fscan(r, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}
		h := make(map[int]int)
		p := make(map[int]int)
		for i, v := range arr {
			pv, ok := p[v]
			if ok {
				// 奇偶性不同 可以楼层更高
				if (pv%2 == 1 && i%2 == 0) || (pv%2 == 0 && i%2 == 1) {
					h[v]++
				}
				p[v] = i
				continue
			}
			p[v] = i
			h[v]++
		}
		for i := 1; i <= n; i++ {
			fmt.Fprint(w, h[i], " ")
		}
		fmt.Fprintln(w)
	}
}

func CF1706C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	cal := func(arr []int64, i int64) int64 {
		return max(0, max(arr[i-1], arr[i+1])-arr[i]+1)
	}

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		arr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}

		if n&1 == 1 {
			ans := int64(0)
			for i := int64(1); i < n; i += 2 {
				ans += cal(arr, i)
			}
			fmt.Fprintln(w, ans)
			continue
		}

		ans := int64(0)
		for i := int64(2); i < n; i += 2 {
			ans += cal(arr, i)
		}
		res := ans
		for i := int64(1); i < n-1; i += 2 {
			ans += cal(arr, i) - cal(arr, i+1)
			res = min(res, ans)
		}
		fmt.Fprintln(w, res)
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
