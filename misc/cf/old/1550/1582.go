package _550

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func CF1582A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var a, b, c int64
		fmt.Fscan(r, &a, &b, &c)
		ans := 0
		if c%2 == 1 {
			ans += 3
		}
		if b%2 == 1 {
			ans += 2
		}
		if a%2 == 1 {
			ans += 1
		}
		fmt.Fprintln(w, ans%2)
	}
}

func CF1582B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, val int64
		fmt.Fscan(r, &n)
		a, b := int64(0), int64(0)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(r, &val)
			if val == 0 {
				a++
			} else if val == 1 {
				b++
			}
		}
		fmt.Fprintln(w, (1<<a)*b)
	}
}

func CF1582C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var s string
		var n int
		fmt.Fscan(r, &n, &s)
		f, res := false, math.MaxInt32
		for i := 0; i < 26; i++ {
			c := byte(i + 'a')
			left, right := 0, n-1
			found, ans := true, 0
			for left < right {
				if s[left] == s[right] {
					left++
					right--
					continue
				}
				if s[left] == c || s[right] == c {
					if s[left] == c {
						left++
						ans++
						continue
					}
					if s[right] == c {
						right--
						ans++
					}
					continue
				}
				found = false
				break
			}
			if found {
				f = true
				if res > ans {
					res = ans
				}
			}
		}
		if f {
			fmt.Fprintln(w, res)
		} else {
			fmt.Fprintln(w, -1)
		}
	}
}

func CF1582D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		arr := make([]int64, n+1)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}

		res := make([]int64, n)
		if n%2 == 0 {
			for i := 0; i < n; i += 2 {
				if (arr[i] > 0 && arr[i+1] > 0) || (arr[i] < 0 && arr[i+1] < 0) {
					res[i] = arr[i+1]
					res[i+1] = -arr[i]
				} else {
					res[i] = -arr[i+1]
					res[i+1] = arr[i]
				}
			}
		} else {
			for i := 0; i < n-3; i += 2 {
				if (arr[i] > 0 && arr[i+1] > 0) || (arr[i] < 0 && arr[i+1] < 0) {
					res[i] = arr[i+1]
					res[i+1] = -arr[i]
				} else {
					res[i] = -arr[i+1]
					res[i+1] = arr[i]
				}
			}

			a, b := arr[n-3], arr[n-2]+arr[n-1]
			aIdx, bIdx, cIdx := n-3, n-2, n-1
			if a == 0 || b == 0 {
				a, b = arr[n-2], arr[n-3]+arr[n-1]
				aIdx, bIdx, cIdx = n-2, n-3, n-1
				if a == 0 || b == 0 {
					a, b = arr[n-1], arr[n-3]+arr[n-2]
					aIdx, bIdx, cIdx = n-1, n-3, n-2
				}
			}
			if (a > 0 && b > 0) || (a < 0 && b < 0) {
				res[aIdx] = b
				res[bIdx] = -a
				res[cIdx] = -a
			} else {
				res[aIdx] = -b
				res[bIdx] = a
				res[cIdx] = a
			}
		}
		for _, v := range res {
			fmt.Fprint(w, v, " ")
		}
		fmt.Fprintln(w)
	}
}

func CF1582E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	for fmt.Fscan(in, &T); T > 0; T-- {
		fmt.Fscan(in, &n)
		a := make([]int64, n+1)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		suf := make([]int64, n+1)
		for i := n - 1; i >= 0; i-- {
			suf[i] = max(suf[i+1], a[i])
			a[i] += a[i+1]
		}
		suf2 := make([]int64, n)
		k := 2
		for l := n - 1; l >= k; l -= k - 1 {
			suf2[l-k+1] = 0
			for i := l - k; i >= 0; i-- {
				s := a[i] - a[i+k]
				if suf2[i+1] < s && s < suf[i+k] {
					suf2[i] = s
				} else {
					suf2[i] = suf2[i+1]
				}
			}
			if suf2[0] == 0 {
				break
			}
			suf, suf2 = suf2, suf
			k++
		}
		fmt.Fprintln(out, k-1)
	}
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
