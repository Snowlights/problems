//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

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
	rn := func() (x int) {
		b := rc()
		neg := false
		for ; '0' > b; b = rc() {
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	var n, m, op, val, last, ans int
	t := &treap{}
	n, m = r(), r()
	for i := 0; i < n; i++ {
		t.put(r(), 1)
	}
	for ; m > 0; m-- {
		op, val = r(), rn()
		val ^= last
		switch op {
		case 1:
			t.put(val, 1)
			continue
		case 2:
			t.delete(val)
			continue
		case 3:
			cnt, _ := t.mRank(val)
			last = cnt + 1
		case 4:
			last = t.mSelect(val - 1).key
		case 5:
			last = t.prev(val).key
		case 6:
			last = t.next(val).key
		}
		ans ^= last
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }

type tpNode struct {
	lr       [2]*tpNode
	priority uint // max heap
	sz       int
	msz      int
	key      int
	val      int // dupCnt for multiset
}

// 设置如下返回值是为了方便使用 tpNode 中的 lr 数组
func (o *tpNode) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0 // 左儿子
	}
	return 1 // 右儿子
}

func (o *tpNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *tpNode) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

// 对于取名叫 maintain 还是 pushUp，由于操作的对象是当前节点，个人认为取名 maintain 更为准确
func (o *tpNode) maintain() {
	o.sz = o.val + o.lr[0].size() + o.lr[1].size()
	o.msz = int(o.val) + o.lr[0].mSize() + o.lr[1].mSize()
}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *tpNode) rotate(d int) *tpNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.sz = o.sz; x.msz = o.msz; o.maintain()
	o.maintain()
	x.maintain()
	return x
}

// 前驱（小于 key，且最大的数）
// 等价于 floor(key-1)
func (t *treap) prev(key int) (prev *tpNode) {
	// 另一种写法，适用于含有 lazy delete 的 BST，如替罪羊树等
	//rk, _ := t.mRank(key)
	//return t.mSelect(rk - 1)
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

// 后继（大于 key，且最小的数)
// 等价于 lowerBound(key+1)
func (t *treap) next(key int) (next *tpNode) {
	// 另一种写法，适用于含有 lazy delete 的 BST，如替罪羊树等
	//rk, o := t.mRank(key)
	//if o != nil {
	//	rk += o.val
	//}
	//return t.mSelect(rk)
	for o := t.root; o != nil; {
		if o.cmp(key) != 0 {
			o = o.lr[1]
		} else {
			next = o
			o = o.lr[0]
		}
	}
	return
}

type treap struct {
	rd   uint
	root *tpNode
}

// 也可以直接设 rd 为 1
func newTreap() *treap {
	return &treap{rd: uint(time.Now().UnixNano())/2 + 1}
}

// https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
// https://en.wikipedia.org/wiki/Xorshift
// 当然，也可以用 rand.Int() 或者 rd: rand.NewSource 后 rd.Int63()，后者速度略慢于 fastRand
func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

// 插入一键值对，返回插入后优先级最大的节点
// 先和二叉搜索树的插入一样，先把要插入的点插入到一个叶子上，并随机分配一个优先级，
// 然后跟维护堆一样，如果当前节点的优先级比根大就旋转，如果当前节点是根的左儿子就右旋如果当前节点是根的右儿子就左旋
func (t *treap) _put(o *tpNode, key, val int) *tpNode {
	if o == nil {
		return &tpNode{priority: uint(rand.Int()), sz: 1, msz: 1, key: key, val: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		// 优先级比根大就旋转
		if o.lr[d].priority > o.priority {
			// 是根的左儿子就右旋，反之左旋
			o = o.rotate(d ^ 1)
		}
	} else {
		//o.val = val
		o.val += val
	}
	o.maintain()
	return o
}

func (t *treap) put(key, val int) { t.root = t._put(t.root, key, val) }

// 删除一个键，返回删除后优先级最大的节点，若无节点返回 nil
// 因为 treap 满足堆性质，所以只需要把要删除的节点旋转到叶节点上，然后直接删除就可以了
// 具体的方法就是每次找到优先级最大的儿子，向与其相反的方向旋转，这样要删除的节点会不断下降直到叶节点，然后直接删除
func (t *treap) _delete(o *tpNode, key int) *tpNode {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.val > 1 {
			o.val--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			// o 有两颗子树，把优先级高的子树旋转到根，然后递归在另一棵子树中删除 o
			d = 0
			if o.lr[0].priority > o.lr[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.lr[d] = t._delete(o.lr[d], key)
		}
	}
	o.maintain()
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

// >= key 的元素个数
// 等价于 t.root.size() - t.mRank(key)
func (t *treap) lowerCount(key int) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			cnt += o.val + o.lr[1].mSize()
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			cnt += o.val + o.lr[1].mSize()
			return
		}
	}
	return
}

// < key 的元素个数
func (t *treap) mRank(key int) (cnt int, o *tpNode) {
	for o = t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			cnt += o.val + o.lr[0].mSize()
			o = o.lr[1]
		default:
			cnt += o.lr[0].mSize()
			// 额外加上 1 或 o.dupCnt 就是 <= key 的元素个数
			return
		}
	}
	return
}

// kth: 排名为 k 的节点 o（即有 k 个键小于 o.key）
// 维护子树和的写法见 https://codeforces.com/contest/1398/submission/119651187
func (t *treap) mSelect(k int) (o *tpNode) {
	//if k < 0 {
	//	return
	//}
	for o = t.root; o != nil; {
		if ls := o.lr[0].mSize(); k < ls {
			o = o.lr[0]
		} else {
			k -= ls + o.val
			if k < 0 {
				return
			}
			o = o.lr[1]
		}
	}
	return
}
