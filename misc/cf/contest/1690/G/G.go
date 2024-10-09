package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node struct {
	lr       [2]*node
	priority uint
	key      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
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

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
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
		o.lr[d] = t._delete(o.lr[d], key)
	}
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap) prev(key int) (prev *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

func (t *treap) next(key int) (next *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func run(_r io.Reader, _w io.Writer) {
	const eof = 0
	out := bufio.NewWriter(_w)
	defer out.Flush()

	_i, _n, buf := 0, 0, make([]byte, 1<<12) // 4KB
	// 读一个字符
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 { // EOF
				return eof
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	// 读一个整数，支持负数
	r := func() (x int) {
		b := rc()
		neg := false
		for ; '0' > b || b > '9'; b = rc() {
			// 某些多组数据的题目，不告诉有多少组数据，那么需要额外判断是否读到了 EOF
			if b == eof {
				return
			}
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	for T := r(); T > 0; T-- {
		n, m := r(), r()
		a := make([]int, n)
		for i := range a {
			a[i] = r()
		}
		ans := 0
		t := &treap{rd: 1}
		mi := a[0] + 1
		for i, v := range a {
			if v < mi {
				ans++
				t.put(i)
				mi = v
			}
		}
		for ; m > 0; m-- {
			k, d := r(), r()
			k--
			v := a[k] - d
			if o := t.prev(k); o != nil && a[k] >= a[o.key] && v < a[o.key] {
				ans++
				t.put(k)
			}
			for o := t.next(k); o != nil && a[o.key] >= v; o = t.next(k) {
				ans--
				t.delete(o.key)
			}
			a[k] = v
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}

}

func main() { run(os.Stdin, os.Stdout) }
