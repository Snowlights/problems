package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// https://atcoder.jp/contests/abl/tasks/abl_e
//
// 输入 n(≤2e5) 和 q(≤2e5)。
// 初始有一个长为 n 的字符串 s，所有字符都是 1，s 的下标从 1 开始。
// 然后输入 q 个替换操作，每个操作输入 L,R (1≤L≤R≤n) 和 d (1≤d≤9)。
// 你需要把 s 的 [L,R] 内的所有字符替换为 d。
// 对每个操作，把替换后的 s 看成一个十进制数，输出这个数模 998244353 的结果。

// 珂朵莉树 简单模版
const mod = 998244353

var ones, p10 []int
var n, ans int

type node struct {
	lr       [2]*node
	priority uint
	l, r, d  int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.l:
		return 0
	case b > o.l:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap struct {
	rd   uint
	root *node
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) _put(o *node, l, r, dig int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), l: l, r: r, d: dig}
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._put(o.lr[d], l, r, dig)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.d = dig
	}
	return o
}

func (t *treap) put(l, r, d int) { t.root = t._put(t.root, l, r, d) }

func (t *treap) _delete(o *node, l int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], l)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		d = 0
		if o.lr[0].priority > o.lr[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.lr[d] = t._delete(o.lr[d], l)
	}
	return o
}

func (t *treap) delete(l int) { t.root = t._delete(t.root, l) }

func (t *treap) floor(key int) (floor *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *treap) next(l int) (next *node) {
	for o := t.root; o != nil; {
		if o.cmp(l) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func (t *treap) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, d := o.r, o.d
		o.r = mid - 1
		t.put(mid, r, d)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap) updateAns(o *node, d int) {
	ans = (ans + (d-o.d)*ones[o.r-o.l+1]*p10[n-o.r]%mod + mod) % mod
}

func (t *treap) merge(l, r, d int) {
	t.prepare(l, r)
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		t.updateAns(o, d)
		t.delete(o.l)
	}
	o := t.floor(l)
	t.updateAns(o, d)
	o.r = r
	o.d = d
}

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, l, r, d int
	fmt.Fscan(in, &n, &q)
	ones = make([]int, n+1)
	p10 = make([]int, n+1)
	p10[0] = 1
	for i := 1; i <= n; i++ {
		ones[i] = (ones[i-1]*10 + 1) % mod
		p10[i] = p10[i-1] * 10 % mod
	}
	ans = ones[n]
	t := newTreap()
	t.put(1, n, 1)
	for ; q > 0; q-- {
		fmt.Fscan(in, &l, &r, &d)
		t.merge(l, r, d)
		fmt.Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
