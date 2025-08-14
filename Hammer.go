package main

import (
	"container/heap"
	"math"
)

type UIntHeap []uint

func (h UIntHeap) Len() int           { return len(h) }
func (h UIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h UIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *UIntHeap) Push(x any)        { *h = append(*h, x.(uint)) }
func (h *UIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Hammer(n int) uint {
	var limit uint = math.MaxUint
	h := UIntHeap{1}
	heap.Init(&h)

	for i := 0; i < n - 1; i++{
		m := heap.Pop(&h).(uint)
		for len(h) > 0 && m == (h)[0]{
			heap.Pop(&h)
		}

		for _,v := range []uint{2,3,5}{
			if m <= limit / v{
				heap.Push(&h, m*v)
			}
		}
	}

	return heap.Pop(&h).(uint)
}

func dotest(n int, exp uint) {
	Expect(Hammer(n), exp)
}

func TestFunc() {
	dotest(1, 1)
	dotest(2, 2)
	dotest(3, 3)
	dotest(4, 4)
	dotest(5, 5)
	dotest(6, 6)
	dotest(7, 8)
	dotest(8, 9)
	dotest(9, 10)
	dotest(10, 12)
	dotest(11, 15)
	dotest(12, 16)
	dotest(13, 18)
	dotest(14, 20)
	dotest(15, 24)
	dotest(16, 25)
	dotest(17, 27)
	dotest(18, 30)
	dotest(19, 32)
}
