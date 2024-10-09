package algorithm

import (
	"math/rand"
	"time"
)

/* 树堆 treap=tree+heap
本质上属于笛卡尔树，见 cartesian_tree.go

https://oi-wiki.org/ds/treap/
https://en.wikipedia.org/wiki/Treap
复杂度证明 http://www.cs.cmu.edu/afs/cs/academic/class/15210-s12/www/lectures/lecture16.pdf
部分代码参考刘汝佳实现，见 https://github.com/klb3713/aoapc-book/blob/master/TrainingGuide/bookcodes/ch3/la5031.cpp
额外维护子树和的写法见 https://codeforces.com/contest/1398/submission/119651187

模板题 https://www.luogu.com.cn/problem/P3369 https://www.luogu.com.cn/problem/P6136
题目推荐 https://cp-algorithms.com/data_structures/treap.html#toc-tgt-8

EXTRA: FHQ Treap
https://baobaobear.github.io/post/20191215-fhq-treap/
FHQ-Treap 学习笔记 + 一堆题目 https://www.luogu.com.cn/blog/85514/fhq-treap-xue-xi-bi-ji
https://www.luogu.com.cn/blog/specialflag/solution-p3369
*/

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
	//	 rk += o.val
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

func (o *tpNode) pushDown() {
	// ...
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
		return &tpNode{
			priority: uint(rand.Int()), // t.fastRand()
			sz:       1, msz: 1, key: key, val: val}
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

// max <= key
// return nil if not found
func (t *treap) floor(key int) (floor *tpNode) {
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

// min >= key
// return nil if not found
func (t *treap) lowerBound(key int) (lb *tpNode) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
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

func (t *treap) size() int {
	return t.root.size()
}

// 中序遍历，返回所有键值
func (t *treap) keys() []int {
	keys := make([]int, 0, t.size())
	var f func(node *tpNode)
	f = func(o *tpNode) {
		if o == nil {
			return
		}
		o.pushDown()
		f(o.lr[0])
		keys = append(keys, o.key)
		// 如果是多重集则需要多次插入
		//for i := 0; i < o.value; i++ {
		//	keys = append(keys, o.key)
		//}
		f(o.lr[1])
	}
	f(t.root)
	return keys
}
