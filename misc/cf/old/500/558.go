package _00

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

//方法一：朴素思路
//
//提示 1：对于任意一个 a[i]，它能操作出多少个不同的数字？
//观察：每个数字不会超过 max(a)，继续增大是没有意义的。
//不妨以 a[0] 为准，设所有数字最终变成 x，
// 那么 x 必然是由 a[0] 的二进制表示的某个前缀+若干个尾 0 组成
// （即 x = a[0] 右移若干位再左移若干位）。
// 因此得到的不同 x 不会超过 O(17^2) 个。
//
//提示 2：枚举 x，然后计算每个 a[i] 变成 x 需要操作多少次，取次数之和的最小值，作为答案。

func CF558D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, x int
	fmt.Fscan(r, &n)
	cnt, vis := make([]int, 1e5+1), make([]int, 1e5+1)
	step := make([]int, 1e5+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &x)

		q := []struct{ x, step int }{{x, 0}}
		for len(q) > 0 {
			val := q[0]
			q = q[1:]
			if val.x > 1e5+1 {
				continue
			}
			if vis[val.x] == i {
				continue
			}
			vis[val.x] = i
			step[val.x] += val.step
			cnt[val.x]++
			q = append(q, struct{ x, step int }{val.x * 2, val.step + 1})
			q = append(q, struct{ x, step int }{val.x / 2, val.step + 1})
		}
	}
	ans := math.MaxInt32
	for i, v := range cnt {
		if v == n {
			ans = min(ans, step[i])
		}
	}
	fmt.Fprintln(w, ans)

}

// https://codeforces.com/contest/558/submission/173887212
//方法二：优化思路（上面链接的写法）
//
//方法一需要 O(17^2 * n) 的时间，可以进一步优化到 O(17n)。
//
//统计每个前缀的出现次数 cnt[v]，以及变为该前缀所需的操作次数 sum[v]。
//那么最大的满足 cnt[v] = n 的 v 就是 x 的最小值，继续往下减小是没有意义的。
//对应的 sum[v] 也是答案的最大值。
//根据中位数贪心的启发，尝试 2v 4v 8v ... 看能否得到更优解。
//考虑操作次数的变化量，从 v 到 2v，有 cnt[2v] 个前缀可以少操作一次，有 n-cnt[2v] 个前缀要多操作一次。如果 cnt[2v] > n-cnt[2v]，那么所有数变成 2v 会得到更小的操作次数。依此类推。
//如果 cnt[2v] > n-cnt[2v] 不成立，其实可以直接退出循环，这是因为数字越大 cnt 值就越小。

func CF558C2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	cnt := [2e5 + 1]int{}
	sum := [1e5 + 1]int{}
	var n, v int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v)
		for c := 0; v > 0; v /= 2 {
			cnt[v]++
			sum[v] += c
			c++
		}
	}

	v = 1e5
	for cnt[v] < n {
		v--
	}
	ans := sum[v]
	for v *= 2; cnt[v]*2 > n; v *= 2 {
		ans -= cnt[v]*2 - n
	}
	fmt.Fprint(out, ans)
}
