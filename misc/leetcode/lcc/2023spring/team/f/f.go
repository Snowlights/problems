package f

import (
	"fmt"
	"math/bits"
)

func log10(n int32) int32 {
	if n < 10 {
		return 1
	}
	return log10(n/10) + 1
}

func gcd(a, b int32, x, y *int32) int32 {
	if !(b > 0) {
		*x, *y = 1, 0
		return a
	}
	res := gcd(b, a%b, y, x)
	// fmt.Println(res, *x, *y)
	*y -= *x * (a / b)
	return res
}

func inv(a, m int32) int32 {
	var x, y int32
	gcd(a, m, &x, &y)
	return (x%m + m) % m
}

func fac(n int32) int32 {
	res := int32(1)
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func C(n, m int32) int32 {
	res := int32(1)
	for i := int32(1); i <= m; i++ {
		res = res * (n - i + 1) / i
	}
	return res
}

const (
	S  = 16
	S1 = 32 - S
	M  = 1996090921
)

type Node struct {
	x, y, t int32
}

var h [(1 << S) + 105]*Node
var T int32 = 1

func H(x int32) int {
	m := uint32(x) * M
	d := m >> S1
	return int(d)
}

func insert(x int32) {
	index := H(x)
	fmt.Println(x, index)
	var i int
	var v *Node
	for i, v = range h[index:] {
		if v.x == x {
			v.y++
			return
		}
	}
	h[index+i].t = T
	h[index+i].x = x
	h[index+i].y = 1
}

func find(x int32) int32 {
	index := H(x)
	for _, v := range h[index:] {
		if v.x == x {
			return v.y
		}
	}
	return 0
}

type Node2 struct {
	v1, v2, l int32
}

const (
	N  = 9
	M0 = 205
)

type FTwo func(i, j int32) bool
type FOne func(j int32) bool

func pop(x int32) int32 {
	return int32(bits.OnesCount(uint(x)))
}

func higher(a, b int32) int32 {
	if a > b {
		return 1
	}
	return 0
}

func max(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

var pow10, pinv [M0]int32
var l [N]int32
var llen [1 << N]int32
var n, ans, p, r, B int32
var c [1 << N][]int32
var d [1 << (N + 1)][]*Node2

func cal(t0 int32, f1, f2 func(i, j int32) bool, f3 func(j int32) bool) {
	for i := (1 << n) + 1; i < (1 << (n + 1)); i++ {
		d[i] = []*Node2{}
	}
	for i := int32(1 << n); i < (1 << (n + 1)); i++ {
		if pop(int32(i)) <= t0 {
			for j, _f2 := int32((i-1)&i), false; (j >> n) > 0; j = (j - 1) & i {
				_f2 = f2(i, j)
				if _f2 || f1(i, j) {
					for _, t := range d[j] {
						l1 := llen[j-(1<<n)] - t.l + 2*higher(j, 1<<n)
						for _, v2 := range c[i-j] {
							d[i] = append(d[i], &Node2{
								v1: (t.v1 + pow10[l1]) % p,
								v2: int32((int64(t.v2)*int64(pow10[llen[i-j]+1]) + int64(v2)*10 + 9) % int64(p)),
								l:  t.l + llen[i-j] + 1,
							})
							if !_f2 {
								d[i] = append(d[i], &Node2{
									v1: int32((int64(t.v1) + int64(pow10[l1+llen[i-j]]) + int64(v2)*int64(pow10[l1])) % int64(p)),
									v2: int32(int64(t.v2)*10 + (9)%int64(p)),
									l:  t.l + 1,
								})
							}
						}

					}
				}
			}
			j := (1 << (n + 1)) - 1 - i
			T++
			fmt.Println("cal", j, T, f3(j))
			if f3(j) {
				continue
			}
			for _, v := range c[j] {
				insert(v)
			}
			for _, t := range d[i] {
				fmt.Println("auto t", int32(
					((int64(r)-int64(t.v2)-int64(t.v1)*int64(pow10[llen[j]+t.l]))%int64(p)+int64(p))*int64(pinv[t.l])%int64(p)))
				ans += find(int32(
					((int64(r)-int64(t.v2)-int64(t.v1)*int64(pow10[llen[j]+t.l]))%int64(p) + int64(p)) * int64(pinv[t.l]) % int64(p)))
			}
		}
	}
}

// todo undone
func treeOfInfiniteSouls(a []int, pp, rr int) int {
	for i := range h {
		h[i] = &Node{}
	}

	n = int32(len(a))
	B = (n + 2) / 3
	p, r = int32(pp), int32(rr)
	if p == 2 || p == 5 {
		if p == 2 && r == 1 || p == 5 && r == 4 {
			return int(C((n-1)*2, n-1) / n * fac(n))
		}
		return 0
	}

	pow10[0] = 1 % p
	for i := 1; i < M0; i++ {
		pow10[i] = int32(int64(pow10[i-1]) * 10 % int64(p))
	}
	for i := 0; i < M0; i++ {
		pinv[i] = inv(pow10[i], p)
	}
	for i := int32(0); i < n; i++ {
		l[i] = log10(int32(a[i]))
	}
	for i := 1; i < (1 << n); i++ {
		llen[i] = (pop(int32(i))*2 - 1) * 2
		for j := int32(0); j < n; j++ {
			if i&(1<<j) > 0 {
				llen[i] += l[j]
			}
		}
	}
	for i := int32(0); i < n; i++ {
		c[1<<i] = append(c[1<<i], int32((int64(a[i])*10+int64(pow10[l[i]+1])+9)%int64(p)))
	}

	for i := 1; i < (1 << n); i++ {
		if pop(int32(i)) <= B*2 {
			t := pow10[llen[i]-1] + 9
			for j := (i - 1) & i; j > 0; j = (j - 1) & i {
				if n == 9 || pop(int32(i)) < max((n+1)/2, 2) ||
					max(pop(int32(j)), pop(int32(i-j))) <= min(B, (n-1)/2) {
					for _, v1 := range c[j] {
						t1 := int64(v1)*int64(pow10[llen[i-j]+1]) + int64(t)
						for _, v2 := range c[i-j] {
							c[i] = append(c[i], int32(int64(v2)*10+t1)%p)
						}
					}
				}
			}
		}
	}

	d[1<<n] = append(d[1<<n], &Node2{0, 0, 0})
	yes := func(a, b int32) bool {
		return true
	}
	no := func(a, b int32) bool {
		return false
	}
	if n == 9 {
		cal(4, yes, no, func(j int32) bool { return pop(j) != 6 })
		cal(5, func(i, j int32) bool {
			return j != (1<<n) || pop(i-j) >= 2
		}, no, func(j int32) bool {
			return pop(j) != 5
		})
		cal(6, func(i, j int32) bool {
			return j != (1<<n) || pop(i-j) == 3
		}, func(i, j int32) bool {
			return j == (1<<n) && pop(i-j) == 4
		}, func(j int32) bool {
			return pop(j) != 4
		})
	} else {
		cal(n/2+1, yes, func(i, j int32) bool {
			return n%2 == 0 && pop(j) == 1 && pop(i-j) == n/2
		}, func(j int32) bool {
			return pop(j) < (n+1)/2 || pop(j) > B*2
		})
	}
	return int(ans)
}
