package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	num1 int64
	num2 int64
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].num1 < h[j].num1 }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var (
		res     int64
		sum     int64
		minHeap MinHeap
		items   []Item = make([]Item, 0, len(nums1))
	)

	for i := 0; i < len(nums1); i++ {
		items = append(items, Item{int64(nums1[i]), int64(nums2[i])})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].num2 > items[j].num2
	})

	for i := 0; i < len(items); i++ {
		if i < k {
			heap.Push(&minHeap, items[i])
			sum += items[i].num1
			res = sum * items[i].num2
			continue
		}

		item := heap.Pop(&minHeap).(Item)
		sum += items[i].num1 - item.num1
		heap.Push(&minHeap, items[i])
		res = max(res, sum*items[i].num2)
	}

	return res
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	fmt.Println(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
}
