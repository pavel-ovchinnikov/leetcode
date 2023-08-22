package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	pq := IntHeap(stones)
	heap.Init(&pq)

	for pq.Len() > 1 {
		a := heap.Pop(&pq).(int)
		b := heap.Pop(&pq).(int)
		if b != a {
			heap.Push(&pq, a-b)
		}
	}

	if pq.Len() == 0 {
		return 0
	}

	return pq[0]
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
