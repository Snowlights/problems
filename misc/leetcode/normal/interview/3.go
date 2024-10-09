package interview

import (
	"container/heap"
	"sort"
)

// 面试题 03.04
type MyQueue struct {
	a []int
}

/** Initialize your data structure here. */
func Constructor0304() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.a[0]
	this.a = this.a[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.a[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.a) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// 面试题 03.05
type SortedStack struct {
	hp *hp2
}

func Constructor0305() SortedStack {
	return SortedStack{&hp2{}}
}

func (this *SortedStack) Push(val int) {
	heap.Push(this.hp, val)
}

func (this *SortedStack) Pop() {
	if this.hp.Len() == 0 {
		return
	}
	heap.Pop(this.hp)
}

func (this *SortedStack) Peek() int {
	if this.hp.Len() == 0 {
		return -1
	}
	return this.hp.IntSlice[0]
}

func (this *SortedStack) IsEmpty() bool {
	return this.hp.Len() == 0
}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */

// 动物收容所。有家动物收容所只收容狗与猫，
//且严格遵守“先进先出”的原则。在收养该收容所的动物时
//，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，
//或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，
//收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，
//实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。
//允许使用Java内置的LinkedList数据结构。
//
//enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
//
//dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

// 面试题 03.06
type AnimalShelf struct {
	dogs, cats []int
}

func Constructor0306() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	switch animal[1] {
	case 0:
		this.cats = append(this.cats, animal[0])
	case 1:
		this.dogs = append(this.dogs, animal[0])
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.cats) == 0 && len(this.dogs) == 0 {
		return []int{-1, -1}
	}

	if len(this.cats) != 0 && len(this.dogs) != 0 {
		if this.cats[0] < this.dogs[0] {
			res := this.cats[0]
			this.cats = this.cats[1:]
			return []int{res, 0}
		} else {
			res := this.dogs[0]
			this.dogs = this.dogs[1:]
			return []int{res, 1}
		}
	}

	if len(this.dogs) == 0 {
		res := this.cats[0]
		this.cats = this.cats[1:]
		return []int{res, 0}
	} else {
		res := this.dogs[0]
		this.dogs = this.dogs[1:]
		return []int{res, 1}
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dogs) == 0 {
		return []int{-1, -1}
	}

	res := this.dogs[0]
	this.dogs = this.dogs[1:]
	return []int{res, 1}
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cats) == 0 {
		return []int{-1, -1}
	}

	res := this.cats[0]
	this.cats = this.cats[1:]
	return []int{res, 0}
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
