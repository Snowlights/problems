package _700

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1705A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, diff int64
		fmt.Fscan(r, &n, &diff)
		arr := make([]int, n*2)
		for i := int64(0); i < 2*n; i++ {
			fmt.Fscan(r, &arr[i])
		}
		sort.Ints(arr)
		can := true
		for i := int64(0); i < n; i++ {
			if int64(arr[i+n]-arr[i]) < diff {
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

func CF1705B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		arr := make([]int64, n)
		flag, ans := false, int64(0)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &arr[i])
			if !flag && arr[i] == 0 {
				continue
			}
			if i == n-1 {
				break
			}
			flag = true
			if arr[i] > 0 {
				ans += arr[i]
			} else {
				ans += 1
			}
		}
		fmt.Fprintln(w, ans)
	}
}

func CF1705C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n, c, q int64
		fmt.Fscan(r, &n, &c, &q)
		s := make([]byte, n)
		fmt.Fscan(r, &s)
		left, right := make([]int64, c+1), make([]int64, c+1)
		sum := make([]int64, c+1)
		sum[0] = int64(n)
		for i := int64(1); i <= c; i++ {
			fmt.Fscan(r, &left[i], &right[i])
			sum[i] = sum[i-1] + right[i] - left[i] + 1
		}
		for i := int64(0); i < q; i++ {
			var k int64
			fmt.Fscan(r, &k)
			now := int64(0)
			for sum[now] < k {
				now++
			}
			for now > 0 {
				k -= sum[now] - right[now]
				for now > 0 && sum[now-1] >= k {
					now--
				}
			}
			fmt.Fprintln(w, string(s[k-1]))
		}

	}
}

func CF1705D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	fmt.Fscan(r, &T)
	for ; T > 0; T-- {
		var n int64
		fmt.Fscan(r, &n)
		a, b := make([]byte, n), make([]byte, n)
		fmt.Fscan(r, &a, &b)

		if a[0] != b[0] || a[n-1] != b[n-1] {
			fmt.Fprintln(w, "-1")
			continue
		}
		num, ans := int64(0), int64(0)
		for i := int64(1); i < n; i++ {
			if a[i] != a[i-1] {
				num++
			}
			if b[i] != b[i-1] {
				num--
			}
			ans += abs(num)
		}
		if num != 0 {
			fmt.Fprintln(w, "-1")
		} else {
			fmt.Fprintln(w, ans)
		}

	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
