package _050

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF1066A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	var L, v, ll, rr int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &L, &v, &ll, &rr)
		// 10 2 3 7
		fmt.Fprintln(w, L/v-rr/v+(ll-1)/v)
	}

}

func CF1066B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, rr, ans int64
	fmt.Fscan(r, &n, &rr)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	for i := int64(0); i < n; i++ {
		cnt := 0
		for j := min(i+rr-1, n-1); j >= 0 && j >= i-rr+1; j-- {
			if arr[j] == 1 {
				ans++
				cnt++
				i = j + rr - 1
				break
			}
		}
		if cnt == 0 {
			fmt.Fprintln(w, -1)
			return
		}
	}

	fmt.Fprintln(w, ans)
}

func CF1066C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, id int64
	fmt.Fscan(r, &n)
	var t string
	left, right := int64(0), int64(1)
	h := map[int64]int64{}
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &t, &id)
		switch t {
		case "L":
			h[id] = left
			left--
		case "R":
			h[id] = right
			right++
		default:
			fmt.Fprintln(w, min(h[id]-left-1, right-h[id]-1))
		}
	}
}

func CF1066D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k, ans int64
	fmt.Fscan(r, &n, &m, &k)
	arr := make([]int64, n)
	curM, leftK := m-1, k
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	for i := n - 1; i >= 0; i-- {
		if leftK >= arr[i] {
			leftK -= arr[i]
			ans++
		} else if curM > 0 {
			curM--
			leftK = k - arr[i]
			ans++
		} else {
			break
		}
	}
	fmt.Fprintln(w, ans)
}

func CF1066E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const mod = 998244353
	var n, m, ans int64
	var a, b string
	fmt.Fscan(r, &n, &m, &a, &b)
	bOne, pow := int64(strings.Count(b, "1")), int64(1)
	for i := n - 1; i >= 0 && m-n+i >= 0; i-- {
		if a[i]&1 == 1 {
			ans = (ans + (pow * bOne)) % mod
		}
		pow = pow * 2 % mod
		bOne -= int64(b[m-n+i] & 1)
	}

	fmt.Print(ans % mod)

}
