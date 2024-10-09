package _700

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1708A(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	for ; n > 0; n-- {
		var m int64
		fmt.Fscan(r, &m)
		arr := make([]int64, m)
		for i := int64(0); i < m; i++ {
			fmt.Fscan(r, &arr[i])
		}
		can := true
		for j := int64(1); j < m; j++ {
			if arr[j]%arr[0] != 0 {
				can = false
				break
			}
		}
		if can {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}

}

func CF1708B(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	fmt.Fscan(_r, &T)
	for ; T > 0; T-- {
		var n, l, r int64
		fmt.Fscan(_r, &n, &l, &r)
		can := true
		for i := int64(1); i <= n; i++ {
			if r/i*i < l {
				can = false
				break
			}
		}

		if can {
			fmt.Fprintln(_w, "YES")
			for i := int64(1); i <= n; i++ {
				fmt.Fprint(_w, r/i*i, " ")
			}
			fmt.Fprintln(_w)
		} else {
			fmt.Fprintln(_w, "NO")
		}
	}

}

func CF1708C(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var T int64
	fmt.Fscan(_r, &T)
	for ; T > 0; T-- {
		var n, q int64
		fmt.Fscan(_r, &n, &q)
		arr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(_r, &arr[i])
		}
		s, cq := make([]byte, n), int64(0)
		for i := n - 1; i >= 0; i-- {
			if cq >= arr[i] {
				s[i] = '1'
			} else if cq < q {
				cq++
				s[i] = '1'
			} else {
				s[i] = '0'
			}
		}
		fmt.Fprintln(_w, string(s))
	}

}

func CF1708D(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()
	var T int64
	fmt.Fscan(_r, &T)
	for ; T > 0; T-- {
		var n, cnt int64
		fmt.Fscan(_r, &n)
		arr := make([]int, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(_r, &arr[i])
		}
		for i := int64(1); i < n; i++ {
			if len(arr) == 0 {
				break
			}
			sort.Ints(arr)
			b := make([]int, 0)
			if cnt > 0 {
				cnt--
				b = append(b, arr[0])
			}
			for j := 1; j < len(arr); j++ {
				if arr[j] == arr[j-1] {
					cnt++
				} else {
					b = append(b, arr[j]-arr[j-1])
				}
			}
			arr = b
		}
		if cnt > 0 {
			fmt.Fprintln(_w, 0)
		} else {
			fmt.Fprintln(_w, arr[0])
		}
	}

}
