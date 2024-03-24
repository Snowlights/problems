//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF711D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	// 我们可以通过一次拓扑排序「剪掉」所有树枝，这样就可以将基环和树枝分开，从而简化后续处理流程：
	//- 如果要遍历基环，可以从拓扑排序后入度大于 0 的节点出发，在图上搜索；
	//- 如果要遍历树枝，可以以基环与树枝的连接处为起点，顺着反图来搜索树枝
	// （树枝上的节点在拓扑排序后的入度均为 0），从而将问题转化成一个树形问题。

	var n, to int
	Fscan(in, &n)
	g := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &to)
		to--
		g[i] = append(g[i], to)
		g[to] = append(g[to], i)
		degree[i]++
		degree[to]++
	}
	q := make([]int, 0)
	for i, v := range degree {
		if v == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, from := range tmp {
			for _, to := range g[from] {
				if degree[to]--; degree[to] == 1 {
					q = append(q, to)
				}
			}
		}
	}

	vis := make([]bool, n)
	var f func(int) int
	f = func(from int) int {
		size := 1
		vis[from] = true
		for _, to := range g[from] {
			if !vis[to] && degree[to] > 1 {
				size += f(to)
			}
		}
		return size
	}
	ans := int64(1)
	for i, b := range vis {
		if !b && degree[i] > 1 {
			sz := f(i)
			n -= sz
			ans = ans * (pow(2, int64(sz)) - 2) % mod
		}
	}
	Fprint(out, (ans*pow(2, int64(n))%mod+mod)%mod)
}

func CF711DV2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n int
	Fscan(in, &n)
	to := make([]int, n)
	for i := range to {
		Fscan(in, &to[i])
		to[i]--
	}

	// 具体来说，初始时间戳 clock=1 ，首次访问一个点 x 时，
	// 记录访问这个点的时间 time[x]=clock ，然后将 clock 加一。
	// 如果首次访问一个点，则记录当前时间 startTime=clock，并尝试从这个点出发，
	// 看能否找到环。如果找到了一个之前访问过的点 x，且访问 x 的时间不早于 startTime，
	// 则说明我们找到了一个新的环，此时的环长就是前后两次访问 x 的时间差，
	// 即 clock-time[x]。

	ans := int64(1)
	time := make([]int, n)
	clock := 1
	for x, t := range time {
		if t > 0 {
			continue
		}
		for t0 := clock; x >= 0; x = to[x] {
			if time[x] > 0 {
				if time[x] >= t0 {
					sz := clock - time[x]
					n -= sz
					ans = ans * (pow(2, int64(sz)) - 2) % mod
				}
				break
			}
			time[x] = clock
			clock++
		}
	}
	Fprint(out, (ans*pow(2, int64(n))%mod+mod)%mod)
}

func main() { CF711DV2(os.Stdin, os.Stdout) }
