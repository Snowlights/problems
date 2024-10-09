//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 给定长度为n的序列a，求出三个点i,j,k(0<=i<=j<=k<=n)，使得a[1]+...+a[i]-a[i+1]-...-a[j]+a[j+1]+...+a[k]-a[k+1]-...-a[n]最大

func CF846C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	// d0, d1, d2  (0 <= d0 <= d1 <= d2 <= n)
	// sum(0..d0) - sum(d0..d1) + sum(d1..d2) - sum(d2..n) 最大
	// 对于d1, 找到d0，使得 sum(0..d0) - sum(d0...d1) 最大
	// sum(d0...d1) = sum(0...d1) - sum(0...d0)
	// 2 * sum(0...d0) - sum(0...d1)

	// tmp := sum[d0] - (sum[d1] - sum[d0]) + (sum[d2] - sum[d1]) - (sum[n] - sum[d2])
	// 2 * sum[d0] - 2 * sum[d1] + 2 * sum[d2] - sum[n]
	// 所以对于d1来说，找到它前面的最大值，找到它后面的最大值即可

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i] + v
	}

	d0 := make([]int, n+1)
	d0[0] = 0
	for i := 1; i <= n; i++ {
		d0[i] = d0[i-1]
		if sum[i] > sum[d0[i-1]] {
			d0[i] = i
		}
	}
	d2 := make([]int, n+1)
	d2[n] = n
	for i := n - 1; i >= 0; i-- {
		d2[i] = d2[i+1]
		if sum[i] > sum[d2[i+1]] {
			d2[i] = i
		}
	}
	Println(sum, d0, d2)
	best := 2*sum[0] - 2*sum[0] + 2*sum[0] - sum[n]
	ans := []int{0, 0, 0}

	for i := 0; i <= n; i++ {
		cur := 2*(sum[d0[i]]-sum[i]+sum[d2[i]]) - sum[n]
		if cur > best {
			best = cur
			ans = []int{d0[i], i, d2[i]}
		}
	}
	for i := range ans {
		Fprint(out, ans[i], " ")
	}
}

func main() { CF846C(os.Stdin, os.Stdout) }
